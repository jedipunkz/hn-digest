package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jedipunkz/hn-digest/internal/frontmatter"
	"github.com/jedipunkz/hn-digest/internal/gtranslate"
	"github.com/jedipunkz/hn-digest/internal/llm"
	"github.com/jedipunkz/hn-digest/internal/ogimage"
)

const (
	topN                = 30
	maxTranslationChars = 4000
	// summaryChars caps the extractive summary, which stands in whenever the
	// LLM summary is unavailable (no GEMINI_API_KEY, rate limit, outage).
	summaryChars = 900
	// llmFailureLimit stops calling the LLM for the rest of the run after this
	// many consecutive failures. A bad key or a retired model fails on every
	// article, and 30 doomed calls only slow the workflow down.
	llmFailureLimit = 3
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
	path      string
	hnID      int
	title     string
	hnURL     string
	sourceURL string
	imageURL  string
	// imageChecked distinguishes `image: ""` (the crawler looked and the page
	// declares none) from a missing key (crawled before image extraction existed).
	imageChecked bool
	score        int
	finalScore   int
	comments     int
	postedAt     string
	translation  string
	// summary is a previously generated LLM summary cached in the front
	// matter, so each article costs at most one API call ever.
	summary string
}

// summaryArticle is the JSON-serialisable output per article.
type summaryArticle struct {
	Rank       int    `json:"rank"`
	HnID       int    `json:"hn_id"`
	Title      string `json:"title"`
	TitleJA    string `json:"title_ja"`
	HnURL      string `json:"hn_url"`
	SourceURL  string `json:"source_url"`
	ImageURL   string `json:"image_url,omitempty"`
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
	llmMax := fs.Int("llm-max", topN, "max LLM summary calls per run (0 disables the LLM)")
	llmInterval := fs.Duration("llm-interval", 4*time.Second, "pause between LLM calls, to stay inside the free-tier requests-per-minute limit")
	if err := fs.Parse(args); err != nil {
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

	translator := &gtranslate.Translator{Client: &http.Client{Timeout: 60 * time.Second}}
	imageClient := &http.Client{Timeout: 20 * time.Second}
	results := make([]summaryArticle, 0, len(articles))

	summarizer := llm.FromEnv()
	if summarizer == nil {
		log.Println("GEMINI_API_KEY is not set; falling back to extractive summaries.")
	}
	llmCalls, llmFailures := 0, 0

	for i, art := range articles {
		rank := i + 1
		log.Printf("[%d/%d] final=%d hn=%d title=%s", rank, len(articles), art.finalScore, art.score, truncate(art.title, 60))
		// A failed title translation must not drop the article: fall back to the
		// original English title, exactly as summaries fall back to the lead.
		titleJA := art.title
		if translated, err := translator.Translate(ctx, art.title); err != nil {
			log.Printf("warning: translate title %q: %v", art.title, err)
		} else if cleaned := cleanTitle(translated); cleaned != "" {
			titleJA = cleaned
		}
		if !art.imageChecked && art.sourceURL != "" {
			art.imageURL = backfillImage(ctx, imageClient, art)
		}
		summary := art.summary
		if summary == "" && summarizer != nil && llmCalls < *llmMax && llmFailures < llmFailureLimit {
			// Pace the calls: the free tier is limited per minute, and a 429
			// burns a retry that a short wait avoids outright.
			if llmCalls > 0 && *llmInterval > 0 {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(*llmInterval):
				}
			}
			llmCalls++
			generated, err := summarizer.Summarize(ctx, art.title, translationBody(art.translation))
			if err != nil {
				// Never fatal: the extractive summary below still ships.
				llmFailures++
				log.Printf("warning: llm summary hn=%d: %v", art.hnID, err)
				if llmFailures >= llmFailureLimit {
					log.Printf("warning: %d consecutive LLM failures; extractive summaries for the rest of this run.", llmFailures)
				}
			} else {
				llmFailures = 0
				summary = generated
				if err := persistSummary(art.path, summary); err != nil {
					log.Printf("warning: cache summary for %s: %v", art.path, err)
				}
			}
		}
		if summary == "" {
			summary = leadSummary(art.translation, summaryChars)
		}
		results = append(results, summaryArticle{
			Rank:       rank,
			HnID:       art.hnID,
			Title:      art.title,
			TitleJA:    titleJA,
			HnURL:      art.hnURL,
			SourceURL:  art.sourceURL,
			ImageURL:   art.imageURL,
			Score:      art.score,
			FinalScore: art.finalScore,
			Comments:   art.comments,
			PostedAt:   art.postedAt,
			SummaryJA:  summary,
		})
	}

	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		return err
	}
	outPath := filepath.Join(*outDir, *date+".json")
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
		path := filepath.Join(dir, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("read %s: %v", entry.Name(), err)
			continue
		}
		text := string(data)
		hnID := frontmatter.Int(text, "hn_id")
		score := frontmatter.Int(text, "score")
		if hnID == 0 || score == 0 {
			continue
		}
		title := frontmatter.String(text, "title")
		translation := extractTranslation(text)
		multiplier := interestMultiplier(title, translation)
		articles = append(articles, parsedArticle{
			path:         path,
			hnID:         hnID,
			title:        title,
			hnURL:        frontmatter.String(text, "hn_url"),
			sourceURL:    frontmatter.String(text, "source"),
			imageURL:     frontmatter.String(text, "image"),
			imageChecked: frontmatter.Has(text, "image"),
			score:        score,
			finalScore:   int(float64(score) * multiplier),
			comments:     frontmatter.Int(text, "comments"),
			postedAt:     frontmatter.String(text, "posted_at"),
			translation:  translation,
			summary:      frontmatter.String(text, "summary_ja"),
		})
	}
	return articles, nil
}

// backfillImage fetches the preview image for an article that was crawled
// before image extraction existed, and records the result in the markdown front
// matter so the page is fetched once per article rather than once per run.
//
// Without this, the feed would only ever show images on stories crawled after
// the feature shipped — the crawler skips stories already in contents/, so it
// never revisits them.
func backfillImage(ctx context.Context, client *http.Client, art parsedArticle) string {
	imageURL, err := ogimage.Fetch(ctx, client, art.sourceURL)
	if err != nil {
		// Transient (timeout, 403, ...): leave the key absent so a later run
		// retries instead of recording "no image" forever.
		log.Printf("warning: fetch image for %s: %v", art.sourceURL, err)
		return ""
	}
	if err := persistImage(art.path, imageURL); err != nil {
		log.Printf("warning: record image for %s: %v", art.path, err)
	}
	return imageURL
}

func persistImage(path, imageURL string) error {
	return persistKey(path, "image", imageURL, withImage)
}

// persistSummary caches the LLM summary in the article's front matter so the
// hourly run re-uses it instead of paying for the same article again. This is
// what keeps the daily call count near the number of new stories rather than
// 30 x 24, which is what would actually blow through the free-tier quota.
func persistSummary(path, summary string) error {
	return persistKey(path, "summary_ja", summary, withSummary)
}

func persistKey(path, key, value string, rewrite func(string, string) (string, bool)) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	updated, ok := rewrite(string(data), value)
	if !ok {
		return fmt.Errorf("no front matter in %s", path)
	}
	return os.WriteFile(path, []byte(updated), 0o644)
}

// withImage sets the front matter image key, matching how cmd/hn-digest writes
// it (%q-quoted). Edits stay inside the front matter block: the article body can
// contain a line starting with "image:" too.
func withImage(text, imageURL string) (string, bool) {
	// Keep the crawler's field order for files written before `image` existed.
	return withFrontMatterKey(text, "image", imageURL, "article_title")
}

// withSummary caches the generated summary next to the other derived keys.
func withSummary(text, summary string) (string, bool) {
	return withFrontMatterKey(text, "summary_ja", summary, "image", "article_title")
}

func keyLineRe(key string) *regexp.Regexp {
	return regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(key) + `:.*$`)
}

// withFrontMatterKey sets key to a %q-quoted value inside the front matter
// block only, so a body line that happens to start with the same key is left
// alone. When the key is absent it is inserted after the first anchor present,
// or appended when none is.
func withFrontMatterKey(text, key, value string, anchors ...string) (string, bool) {
	const closing = "\n---\n"
	if !strings.HasPrefix(text, "---\n") {
		return text, false
	}
	end := strings.Index(text, closing)
	if end < 0 {
		return text, false
	}
	head, rest := text[:end], text[end:]
	line := fmt.Sprintf("%s: %q", key, value)
	if re := keyLineRe(key); re.MatchString(head) {
		return re.ReplaceAllLiteralString(head, line) + rest, true
	}
	for _, anchor := range anchors {
		if loc := keyLineRe(anchor).FindStringIndex(head); loc != nil {
			return head[:loc[1]] + "\n" + line + head[loc[1]:] + rest, true
		}
	}
	return head + "\n" + line + rest, true
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

// translationBody drops the "タイトル: / 記事タイトル: / 説明:" preamble that
// cmd/hn-digest puts above the translated article, so the summary leads with
// actual prose. Falls back to the whole text when there is no body marker.
func translationBody(translation string) string {
	// 記事本文 first: 説明 only stands in when the article body was not fetched.
	for _, marker := range []string{"記事本文:", "説明:"} {
		if idx := strings.Index(translation, marker); idx >= 0 {
			if body := strings.TrimSpace(translation[idx+len(marker):]); body != "" {
				return body
			}
		}
	}
	return strings.TrimPrefix(strings.TrimSpace(translation), "タイトル: ")
}

// leadSummary stands in for an abstractive summary: the lead of the Japanese
// translation, cut at the last sentence boundary so it does not end mid-word.
func leadSummary(translation string, max int) string {
	text := strings.TrimSpace(translationBody(translation))
	runes := []rune(text)
	if len(runes) <= max {
		return text
	}
	head := string(runes[:max])
	if idx := strings.LastIndexAny(head, "。！？"); idx > 0 {
		_, size := utf8.DecodeRuneInString(head[idx:])
		return head[:idx+size]
	}
	return strings.TrimSpace(head) + "…"
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
