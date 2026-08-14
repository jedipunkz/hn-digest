package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	// GitHub Models (https://models.inference.ai.azure.com) was fully retired on
	// 2026-07-30, so the summariser now talks to any OpenAI-compatible chat
	// completions API. The default targets the Gemini API's OpenAI-compatible
	// endpoint, which has a free tier; override with SUMMARIZE_API_BASE /
	// SUMMARIZE_MODEL to use OpenAI, Groq, OpenRouter, etc.
	defaultAPIBase      = "https://generativelanguage.googleapis.com/v1beta/openai"
	defaultModel        = "gemini-3.5-flash-lite"
	topN                = 30
	maxTranslationChars = 4000

	// defaultMinInterval spaces out model calls so a burst stays inside the
	// per-minute request limits that free tiers impose.
	defaultMinInterval = 4 * time.Second
)

// Interest categories: each matched category contributes its bonus once.
// final_score = int(hn_score * (1.0 + sum of matched bonuses))
type interest struct {
	name     string
	keywords []string
	bonus    float64
}

var interests = []interest{
	{
		name: "ai",
		keywords: []string{
			"ai", "llm", "llms", "machine learning", "ml", "openai", "anthropic",
			"claude", "gpt", "gemini", "neural", "artificial intelligence", "deepmind",
			"embedding", "rag", "inference", "fine-tun", "foundation model",
		},
		bonus: 0.3,
	},
	{
		name: "sre",
		keywords: []string{
			"sre", "site reliability", "incident", "on-call", "on call", "observability",
			"monitoring", "alerting", "postmortem", "runbook", "pagerduty", "chaos engineering",
			"toil", "error budget",
		},
		bonus: 0.5,
	},
	{
		name: "platform",
		keywords: []string{
			"platform engineering", "platform", "kubernetes", "k8s", "terraform",
			"infrastructure", "devops", "dev ops", "ci/cd", "cloud run", "gcp",
			"google cloud", "gitops", "helm", "argo", "pulumi", "internal developer",
		},
		bonus: 0.8,
	},
}

// parsedArticle holds data extracted from a markdown file.
type parsedArticle struct {
	hnID        int
	title       string
	hnURL       string
	sourceURL   string
	score       int
	finalScore  int
	comments    int
	postedAt    string
	translation string
}

// summaryArticle is the JSON-serialisable output per article.
type summaryArticle struct {
	Rank       int    `json:"rank"`
	HnID       int    `json:"hn_id"`
	Title      string `json:"title"`
	TitleJA    string `json:"title_ja"`
	HnURL      string `json:"hn_url"`
	SourceURL  string `json:"source_url"`
	Score      int    `json:"score"`
	FinalScore int    `json:"final_score"`
	Comments   int    `json:"comments"`
	PostedAt   string `json:"posted_at"`
	SummaryJA  string `json:"summary_ja"`
}

type daySummary struct {
	Date        string           `json:"date"`
	GeneratedAt string           `json:"generated_at"`
	Articles    []summaryArticle `json:"articles"`
}

type chatRequest struct {
	Model       string    `json:"model"`
	Messages    []message `json:"messages"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func main() {
	if err := run(context.Background(), os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("summarize", flag.ExitOnError)
	date := fs.String("date", time.Now().UTC().Format(time.DateOnly), "date to summarize (YYYY-MM-DD)")
	outDir := fs.String("out", "summaries", "output directory for summary JSON files")
	contentsDir := fs.String("contents", "contents", "base contents directory")
	n := fs.Int("n", topN, "number of top articles to include")
	refresh := fs.Bool("refresh", false, "regenerate summaries even if they already exist")
	minInterval := fs.Duration("min-interval", defaultMinInterval, "minimum delay between model calls")
	if err := fs.Parse(args); err != nil {
		return err
	}

	cfg, err := loadAPIConfig()
	if err != nil {
		return err
	}

	dayDir := filepath.Join(*contentsDir, *date)
	if _, err := os.Stat(dayDir); os.IsNotExist(err) {
		log.Printf("No contents for %s, nothing to do.", *date)
		return nil
	}

	articles, err := loadArticles(dayDir)
	if err != nil {
		return fmt.Errorf("load articles: %w", err)
	}
	if len(articles) == 0 {
		log.Println("No valid articles found.")
		return nil
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].finalScore > articles[j].finalScore
	})
	if len(articles) > *n {
		articles = articles[:*n]
	}

	outPath := filepath.Join(*outDir, *date+".json")

	// This job re-runs hourly over the same day, so previously generated
	// summaries are reused instead of being paid for again. Without this the
	// job burns len(articles)*2 model calls every hour for mostly identical work.
	cached := map[int]summaryArticle{}
	if !*refresh {
		cached = loadExistingSummaries(outPath)
	}

	client := &http.Client{Timeout: 60 * time.Second}
	results := make([]summaryArticle, 0, len(articles))
	throttle := newThrottle(*minInterval)
	var generated, reused, failed int

	for i, art := range articles {
		rank := i + 1
		entry := summaryArticle{
			Rank:       rank,
			HnID:       art.hnID,
			Title:      art.title,
			HnURL:      art.hnURL,
			SourceURL:  art.sourceURL,
			Score:      art.score,
			FinalScore: art.finalScore,
			Comments:   art.comments,
			PostedAt:   art.postedAt,
		}

		// Scores and ranks move between runs, so only the generated text is
		// carried over from the cache; the metadata above is always refreshed.
		if prev, ok := cached[art.hnID]; ok && prev.SummaryJA != "" {
			entry.TitleJA = prev.TitleJA
			entry.SummaryJA = prev.SummaryJA
			if entry.TitleJA == "" {
				entry.TitleJA = art.title
			}
			results = append(results, entry)
			reused++
			continue
		}

		log.Printf("[%d/%d] final=%d hn=%d title=%s", rank, len(articles), art.finalScore, art.score, truncate(art.title, 60))

		throttle.wait(ctx)
		titleJA, err := callModel(ctx, client, cfg, titleMessages(art.title))
		if err != nil {
			log.Printf("skip %q: translate title: %v", truncate(art.title, 60), err)
			failed++
			continue
		}
		titleJA = cleanTitle(titleJA)
		if titleJA == "" {
			titleJA = art.title
		}

		throttle.wait(ctx)
		summary, err := callModel(ctx, client, cfg, summaryMessages(art.title, art.translation))
		if err != nil {
			log.Printf("skip %q: summarize: %v", truncate(art.title, 60), err)
			failed++
			continue
		}

		entry.TitleJA = titleJA
		entry.SummaryJA = summary
		results = append(results, entry)
		generated++
	}

	log.Printf("generated=%d reused=%d failed=%d total=%d", generated, reused, failed, len(results))

	// Never replace a good file with an empty one: if every article failed and
	// nothing was reused, the provider is down and the existing file is better.
	if len(results) == 0 {
		if failed > 0 {
			return fmt.Errorf("all %d articles failed; leaving %s unchanged", failed, outPath)
		}
		log.Println("No summaries produced.")
		return nil
	}

	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(daySummary{
		Date:        *date,
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Articles:    results,
	}, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(outPath, append(data, '\n'), 0o644); err != nil {
		return err
	}
	log.Printf("Saved %d summaries to %s", len(results), outPath)
	return nil
}

// apiConfig describes the OpenAI-compatible endpoint used for inference.
type apiConfig struct {
	endpoint string
	key      string
	model    string
}

func loadAPIConfig() (apiConfig, error) {
	key := os.Getenv("SUMMARIZE_API_KEY")
	if key == "" {
		return apiConfig{}, errors.New("SUMMARIZE_API_KEY is not set (GitHub Models was retired on 2026-07-30; " +
			"set an API key for an OpenAI-compatible provider)")
	}
	base := strings.TrimSuffix(firstNonEmpty(os.Getenv("SUMMARIZE_API_BASE"), defaultAPIBase), "/")
	return apiConfig{
		endpoint: base + "/chat/completions",
		key:      key,
		model:    firstNonEmpty(os.Getenv("SUMMARIZE_MODEL"), defaultModel),
	}, nil
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

// loadExistingSummaries reads already-generated summaries so they can be reused.
// A missing or unreadable file simply means nothing is cached.
func loadExistingSummaries(path string) map[int]summaryArticle {
	out := map[int]summaryArticle{}
	data, err := os.ReadFile(path)
	if err != nil {
		return out
	}
	var day daySummary
	if err := json.Unmarshal(data, &day); err != nil {
		log.Printf("ignoring unreadable %s: %v", path, err)
		return out
	}
	for _, a := range day.Articles {
		out[a.HnID] = a
	}
	return out
}

// throttle enforces a minimum delay between successive model calls.
type throttle struct {
	interval time.Duration
	last     time.Time
}

func newThrottle(interval time.Duration) *throttle {
	return &throttle{interval: interval}
}

func (t *throttle) wait(ctx context.Context) {
	if t.interval <= 0 {
		return
	}
	if !t.last.IsZero() {
		if d := t.interval - time.Since(t.last); d > 0 {
			select {
			case <-time.After(d):
			case <-ctx.Done():
			}
		}
	}
	t.last = time.Now()
}

func loadArticles(dir string) ([]parsedArticle, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var articles []parsedArticle
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			log.Printf("read %s: %v", entry.Name(), err)
			continue
		}
		text := string(data)
		hnID := frontMatterInt(text, "hn_id")
		score := frontMatterInt(text, "score")
		if hnID == 0 || score == 0 {
			continue
		}
		title := frontMatterString(text, "title")
		translation := extractTranslation(text)
		multiplier := interestMultiplier(title, translation)
		articles = append(articles, parsedArticle{
			hnID:        hnID,
			title:       title,
			hnURL:       frontMatterString(text, "hn_url"),
			sourceURL:   frontMatterString(text, "source"),
			score:       score,
			finalScore:  int(float64(score) * multiplier),
			comments:    frontMatterInt(text, "comments"),
			postedAt:    frontMatterString(text, "posted_at"),
			translation: translation,
		})
	}
	return articles, nil
}

var translationRe = regexp.MustCompile(`(?s)## Translation\n\n(.*?)(?:\n## |\z)`)

func extractTranslation(text string) string {
	match := translationRe.FindStringSubmatch(text)
	if len(match) < 2 {
		return ""
	}
	t := strings.TrimSpace(match[1])
	runes := []rune(t)
	if len(runes) > maxTranslationChars {
		return string(runes[:maxTranslationChars])
	}
	return t
}

func interestMultiplier(title, translation string) float64 {
	combined := strings.ToLower(title + " " + translation)
	multiplier := 1.0
	for _, in := range interests {
		for _, kw := range in.keywords {
			if strings.Contains(combined, kw) {
				multiplier += in.bonus
				break // each category contributes at most once
			}
		}
	}
	return multiplier
}

// summaryMessages builds the chat messages for summarising an article.
func summaryMessages(title, translation string) []message {
	return []message{
		{
			Role:    "system",
			Content: "あなたは技術ニュースの要約者です。Hacker Newsの記事を日本語で詳しく要約してください。",
		},
		{
			Role: "user",
			Content: fmt.Sprintf(
				"以下のHacker News記事を日本語で詳しく要約してください。\n"+
					"以下の点を必ず含めて、600〜900文字程度（6〜10文）で記述してください。\n"+
					"- 記事の主題と背景\n"+
					"- 技術的な要点や仕組み、利用されている技術スタック\n"+
					"- 著者の主張や結論\n"+
					"- HNコミュニティで注目されている理由や論点\n\n"+
					"タイトル: %s\n\n本文:\n%s",
				title, translation,
			),
		},
	}
}

// titleMessages builds the chat messages for translating an article title.
func titleMessages(title string) []message {
	return []message{
		{
			Role:    "system",
			Content: "あなたは技術ニュースの翻訳者です。Hacker Newsの記事タイトルを自然な日本語に翻訳してください。",
		},
		{
			Role: "user",
			Content: fmt.Sprintf(
				"以下のHacker News記事のタイトルを自然な日本語に翻訳してください。\n"+
					"訳文のみを1行で出力し、引用符や説明、前置きは付けないでください。\n"+
					"製品名・固有名詞・技術用語はそのまま残して構いません。\n\n"+
					"タイトル: %s",
				title,
			),
		},
	}
}

func callModel(ctx context.Context, client *http.Client, cfg apiConfig, messages []message) (string, error) {
	const maxAttempts = 4
	for attempt := 0; attempt < maxAttempts; attempt++ {
		content, retryAfter, err := callModelOnce(ctx, client, cfg, messages)
		if err == nil {
			return content, nil
		}
		if retryAfter == 0 || attempt == maxAttempts-1 {
			return "", err
		}
		log.Printf("rate limited, waiting %s before retry (%d/%d)", retryAfter, attempt+1, maxAttempts-1)
		select {
		case <-time.After(retryAfter):
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
	panic("unreachable")
}

func callModelOnce(ctx context.Context, client *http.Client, cfg apiConfig, messages []message) (string, time.Duration, error) {
	reqBody := chatRequest{
		Model:       cfg.model,
		Messages:    messages,
		MaxTokens:   1200,
		Temperature: 0.3,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", 0, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, cfg.endpoint, bytes.NewReader(body))
	if err != nil {
		return "", 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cfg.key)

	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return "", 0, err
	}

	// The status code is inspected before the body is parsed: a retired or
	// misconfigured endpoint answers with an empty body, and parsing first turns
	// that into a misleading "unexpected end of JSON input" instead of the
	// actual HTTP error.
	var chatResp chatResponse
	parseErr := json.Unmarshal(data, &chatResp)

	if resp.StatusCode == http.StatusTooManyRequests {
		retryAfter := 65 * time.Second
		if ra := resp.Header.Get("Retry-After"); ra != "" {
			if secs, err2 := strconv.Atoi(ra); err2 == nil {
				retryAfter = time.Duration(secs+5) * time.Second
			}
		}
		return "", retryAfter, fmt.Errorf("API %s: %s", resp.Status, apiErrorMessage(chatResp, data, resp.Status))
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", 0, fmt.Errorf("API %s: %s", resp.Status, apiErrorMessage(chatResp, data, resp.Status))
	}
	if parseErr != nil {
		return "", 0, fmt.Errorf("parse response: %w", parseErr)
	}
	if len(chatResp.Choices) == 0 {
		return "", 0, errors.New("no choices in response")
	}
	return strings.TrimSpace(chatResp.Choices[0].Message.Content), 0, nil
}

// apiErrorMessage prefers the provider's structured error, falling back to the
// raw body so empty or non-JSON responses still produce a usable message.
func apiErrorMessage(resp chatResponse, body []byte, status string) string {
	if resp.Error != nil && resp.Error.Message != "" {
		return resp.Error.Message
	}
	if snippet := strings.TrimSpace(string(body)); snippet != "" {
		return truncate(snippet, 200)
	}
	return status + " (empty response body)"
}

func frontMatterInt(text, key string) int {
	value := frontMatterString(text, key)
	if value == "" {
		return 0
	}
	n, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return n
}

func frontMatterString(text, key string) string {
	re := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(key) + `:\s*"?([^"\n]+)"?\s*$`)
	match := re.FindStringSubmatch(text)
	if len(match) < 2 {
		return ""
	}
	return strings.TrimSpace(match[1])
}

// cleanTitle normalises a translated title: collapse to a single line and strip
// surrounding quotes the model may add.
func cleanTitle(s string) string {
	if idx := strings.IndexByte(s, '\n'); idx >= 0 {
		s = s[:idx]
	}
	s = strings.TrimSpace(s)
	s = strings.Trim(s, `"'「」`)
	return strings.TrimSpace(s)
}

func truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	return string(runes[:n])
}
