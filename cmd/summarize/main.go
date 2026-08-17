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
)

const (
	topN                = 30
	maxTranslationChars = 4000
	// summaryChars caps the extractive summary. GitHub Models was retired
	// (410 github_models_retirement_brownout), so there is no free LLM to
	// abstract the article with; the lead of the Japanese translation stands in.
	summaryChars = 900
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
	results := make([]summaryArticle, 0, len(articles))

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
		summary := leadSummary(art.translation, summaryChars)
		results = append(results, summaryArticle{
			Rank:       rank,
			HnID:       art.hnID,
			Title:      art.title,
			TitleJA:    titleJA,
			HnURL:      art.hnURL,
			SourceURL:  art.sourceURL,
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
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
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
			hnID:        hnID,
			title:       title,
			hnURL:       frontmatter.String(text, "hn_url"),
			sourceURL:   frontmatter.String(text, "source"),
			score:       score,
			finalScore:  int(float64(score) * multiplier),
			comments:    frontmatter.Int(text, "comments"),
			postedAt:    frontmatter.String(text, "posted_at"),
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
