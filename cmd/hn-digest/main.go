package main

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	hnAPIBase       = "https://hacker-news.firebaseio.com/v0"
	hnItemURLFormat = "https://news.ycombinator.com/item?id=%d"
	userAgent       = "hn-digest/0.1 (+https://github.com/jedipunkz/hn-digest)"
)

const defaultTitleKeywords = "sre,site reliability,devops,dev ops,platform engineering,incident,on-call,observability,kubernetes,k8s,terraform,ci/cd,google cloud,gcp,cloud run,cloud sql,bigquery,gemini,ai,llm,llms,artificial intelligence,machine learning,ml,openai,anthropic,claude,deepmind,neural network"

func main() {
	if err := run(context.Background(), os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("hn-digest", flag.ExitOnError)
	source := fs.String("source", "firebase", "story ID source: firebase or algolia")
	section := fs.String("section", "topstories", "HN section: topstories, newstories, beststories, askstories, showstories, jobstories")
	limit := fs.Int("limit", 0, "maximum number of HN stories to process; 0 means no limit")
	outDir := fs.String("out", "contents", "base output directory")
	maxArticleChars := fs.Int("max-article-chars", 12000, "maximum extracted article characters to translate per story")
	timeout := fs.Duration("timeout", 25*time.Second, "HTTP request timeout")
	since := fs.Duration("since", 24*time.Hour, "only process stories posted within this duration; 0 disables time filtering")
	titleKeywords := fs.String("title-keywords", defaultTitleKeywords, "comma-separated title keywords to include; empty means no title filtering")
	date := fs.String("date", time.Now().Format(time.DateOnly), "output date directory")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *limit < 0 {
		return errors.New("--limit must be zero or positive")
	}
	if *source != "firebase" && *source != "algolia" {
		return errors.New("--source must be firebase or algolia")
	}

	now := time.Now
	client := &http.Client{Timeout: *timeout}
	translator := &googleTranslator{client: client}
	runner := &crawler{
		client:          client,
		translator:      translator,
		maxArticleChars: *maxArticleChars,
		baseOutputDir:   *outDir,
		outputDir:       filepath.Join(*outDir, *date),
		source:          *source,
		section:         *section,
		limit:           *limit,
		since:           *since,
		titleKeywords:   parseKeywords(*titleKeywords),
		now:             now,
	}
	return runner.run(ctx)
}

type crawler struct {
	client          *http.Client
	translator      translator
	maxArticleChars int
	baseOutputDir   string
	outputDir       string
	source          string
	section         string
	limit           int
	since           time.Duration
	titleKeywords   []string
	now             func() time.Time
}

type translator interface {
	translate(ctx context.Context, text string) (string, error)
}

type hnItem struct {
	ID          int    `json:"id"`
	Deleted     bool   `json:"deleted"`
	Type        string `json:"type"`
	By          string `json:"by"`
	Time        int64  `json:"time"`
	Text        string `json:"text"`
	Dead        bool   `json:"dead"`
	Parent      int    `json:"parent"`
	Poll        int    `json:"poll"`
	Kids        []int  `json:"kids"`
	URL         string `json:"url"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Descendants int    `json:"descendants"`
}

type article struct {
	Title       string
	Description string
	Text        string
}

func (c *crawler) run(ctx context.Context) error {
	ids, err := c.fetchIDs(ctx)
	if err != nil {
		return err
	}
	if c.limit > 0 && len(ids) > c.limit {
		ids = ids[:c.limit]
	}
	if err := os.MkdirAll(c.outputDir, 0o755); err != nil {
		return err
	}
	seen, err := loadSeen(c.baseOutputDir)
	if err != nil {
		return err
	}

	cutoff := time.Time{}
	if c.since > 0 {
		cutoff = c.now().Add(-c.since)
	}

	translatedCount := 0
	for i, id := range ids {
		item, err := c.fetchItem(ctx, id)
		if err != nil {
			log.Printf("fetch item %d: %v", id, err)
			continue
		}
		if item.Deleted || item.Dead || item.Type != "story" || strings.TrimSpace(item.Title) == "" {
			continue
		}
		if !cutoff.IsZero() && time.Unix(item.Time, 0).Before(cutoff) {
			continue
		}
		matched, keyword := titleMatchesKeywords(item.Title, c.titleKeywords)
		if !matched {
			continue
		}
		if seen.has(item) {
			log.Printf("[%d/%d] skipping duplicate %s", i+1, len(ids), item.Title)
			continue
		}
		if keyword == "" {
			log.Printf("translating %d: %s", translatedCount+1, item.Title)
		} else {
			log.Printf("translating %d [%s]: %s", translatedCount+1, keyword, item.Title)
		}
		if err := c.writeStory(ctx, item); err != nil {
			log.Printf("write story %d: %v", item.ID, err)
			continue
		}
		seen.add(item)
		translatedCount++
	}
	log.Printf("translated %d stories", translatedCount)
	return nil
}

func (c *crawler) fetchIDs(ctx context.Context) ([]int, error) {
	if c.source == "algolia" {
		return c.fetchAlgoliaIDs(ctx)
	}

	var ids []int
	if err := c.getJSON(ctx, fmt.Sprintf("%s/%s.json", hnAPIBase, c.section), &ids); err != nil {
		return nil, fmt.Errorf("fetch %s: %w", c.section, err)
	}
	return ids, nil
}

type algoliaResponse struct {
	Hits []struct {
		ObjectID string `json:"objectID"`
		Title    string `json:"title"`
	} `json:"hits"`
	Page    int `json:"page"`
	NBPages int `json:"nbPages"`
	NBHits  int `json:"nbHits"`
}

func (c *crawler) fetchAlgoliaIDs(ctx context.Context) ([]int, error) {
	if c.since <= 0 {
		return nil, errors.New("--source algolia requires --since")
	}

	cutoff := c.now().Add(-c.since).Unix()
	var ids []int
	for page := 0; ; page++ {
		params := url.Values{}
		params.Set("tags", "story")
		params.Set("numericFilters", fmt.Sprintf("created_at_i>=%d", cutoff))
		params.Set("hitsPerPage", "100")
		params.Set("page", strconv.Itoa(page))
		endpoint := "https://hn.algolia.com/api/v1/search_by_date?" + params.Encode()

		var result algoliaResponse
		if err := c.getJSON(ctx, endpoint, &result); err != nil {
			return nil, fmt.Errorf("fetch algolia page %d: %w", page, err)
		}
		for _, hit := range result.Hits {
			if matched, _ := titleMatchesKeywords(hit.Title, c.titleKeywords); !matched {
				continue
			}
			id, err := strconv.Atoi(hit.ObjectID)
			if err != nil {
				continue
			}
			ids = append(ids, id)
		}
		if page+1 >= result.NBPages || len(result.Hits) == 0 {
			log.Printf("loaded %d matching story IDs from Algolia (%d reported hits)", len(ids), result.NBHits)
			return ids, nil
		}
	}
}

func (c *crawler) fetchItem(ctx context.Context, id int) (hnItem, error) {
	var item hnItem
	if err := c.getJSON(ctx, fmt.Sprintf("%s/item/%d.json", hnAPIBase, id), &item); err != nil {
		return item, err
	}
	return item, nil
}

func (c *crawler) getJSON(ctx context.Context, rawURL string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status %s", resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c *crawler) writeStory(ctx context.Context, item hnItem) error {
	article := article{}
	if item.URL != "" {
		fetched, err := fetchArticle(ctx, c.client, item.URL, c.maxArticleChars)
		if err != nil {
			log.Printf("fetch article %d: %v", item.ID, err)
		} else {
			article = fetched
		}
	}

	sourceText := storyTranslationInput(item, article)
	translated, err := c.translator.translate(ctx, sourceText)
	if err != nil {
		return fmt.Errorf("translate: %w", err)
	}

	path := filepath.Join(c.outputDir, uniqueFilename(item.Title, item.ID))
	content := renderMarkdown(c.now(), item, article, translated)
	return os.WriteFile(path, []byte(content), 0o644)
}

type seenIndex struct {
	hnIDs   map[int]bool
	sources map[string]bool
}

func loadSeen(root string) (*seenIndex, error) {
	seen := &seenIndex{
		hnIDs:   map[int]bool{},
		sources: map[string]bool{},
	}
	if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
		return seen, nil
	} else if err != nil {
		return nil, err
	}

	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		text := string(data)
		if id := frontMatterInt(text, "hn_id"); id > 0 {
			seen.hnIDs[id] = true
		}
		if source := frontMatterString(text, "source"); source != "" {
			seen.sources[source] = true
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return seen, nil
}

func (s *seenIndex) has(item hnItem) bool {
	if s.hnIDs[item.ID] {
		return true
	}
	return item.URL != "" && s.sources[item.URL]
}

func (s *seenIndex) add(item hnItem) {
	s.hnIDs[item.ID] = true
	if item.URL != "" {
		s.sources[item.URL] = true
	}
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

func parseKeywords(input string) []string {
	var out []string
	for _, part := range strings.Split(input, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}

func titleMatchesKeywords(title string, keywords []string) (bool, string) {
	if len(keywords) == 0 {
		return true, ""
	}

	titleTokens := tokenizeForMatch(title)
	titleTokenSet := map[string]bool{}
	for _, token := range titleTokens {
		titleTokenSet[token] = true
	}
	normalizedTitle := " " + strings.Join(titleTokens, " ") + " "

	for _, keyword := range keywords {
		keywordTokens := tokenizeForMatch(keyword)
		if len(keywordTokens) == 0 {
			continue
		}
		if len(keywordTokens) == 1 {
			if titleTokenSet[keywordTokens[0]] {
				return true, keyword
			}
			continue
		}
		if strings.Contains(normalizedTitle, " "+strings.Join(keywordTokens, " ")+" ") {
			return true, keyword
		}
	}
	return false, ""
}

func tokenizeForMatch(input string) []string {
	var b strings.Builder
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			continue
		}
		b.WriteRune(' ')
	}
	return strings.Fields(b.String())
}

func storyTranslationInput(item hnItem, article article) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Title: %s\n", item.Title)
	if article.Title != "" && article.Title != item.Title {
		fmt.Fprintf(&b, "Article title: %s\n", article.Title)
	}
	if article.Description != "" {
		fmt.Fprintf(&b, "Description: %s\n", article.Description)
	}
	if item.Text != "" {
		fmt.Fprintf(&b, "HN text: %s\n", htmlToText(item.Text, 4000))
	}
	if article.Text != "" {
		fmt.Fprintf(&b, "\nArticle text:\n%s\n", article.Text)
	}
	return strings.TrimSpace(b.String())
}

func fetchArticle(ctx context.Context, client *http.Client, rawURL string, maxChars int) (article, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return article{}, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,text/plain;q=0.8,*/*;q=0.5")
	resp, err := client.Do(req)
	if err != nil {
		return article{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return article{}, fmt.Errorf("unexpected status %s", resp.Status)
	}
	contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	if !strings.Contains(contentType, "text/html") && !strings.Contains(contentType, "text/plain") && contentType != "" {
		return article{}, fmt.Errorf("unsupported content type %q", contentType)
	}

	limited := io.LimitReader(resp.Body, 2<<20)
	data, err := io.ReadAll(limited)
	if err != nil {
		return article{}, err
	}
	body := string(data)
	if strings.Contains(contentType, "text/plain") {
		return article{Text: trimRunes(cleanWhitespace(body), maxChars)}, nil
	}
	return article{
		Title:       firstMatch(body, `<title[^>]*>(.*?)</title>`),
		Description: firstMeta(body, "description"),
		Text:        extractReadableText(body, maxChars),
	}, nil
}

func extractReadableText(body string, maxChars int) string {
	body = stripTagBlock(body, "script")
	body = stripTagBlock(body, "style")
	body = stripTagBlock(body, "noscript")

	var chunks []string
	re := regexp.MustCompile(`(?is)<(p|h1|h2|h3|li|blockquote)[^>]*>(.*?)</\s*(p|h1|h2|h3|li|blockquote)\s*>`)
	for _, match := range re.FindAllStringSubmatch(body, -1) {
		text := htmlToText(match[2], 2000)
		if len([]rune(text)) >= 30 {
			chunks = append(chunks, text)
		}
	}
	if len(chunks) == 0 {
		chunks = append(chunks, htmlToText(body, maxChars))
	}
	return trimRunes(cleanWhitespace(strings.Join(chunks, "\n\n")), maxChars)
}

func stripTagBlock(input, tag string) string {
	re := regexp.MustCompile(`(?is)<` + tag + `[^>]*>.*?</` + tag + `\s*>`)
	return re.ReplaceAllString(input, " ")
}

func firstMatch(input, pattern string) string {
	re := regexp.MustCompile("(?is)" + pattern)
	match := re.FindStringSubmatch(input)
	if len(match) < 2 {
		return ""
	}
	return cleanWhitespace(htmlToText(match[1], 300))
}

func firstMeta(input, name string) string {
	patterns := []string{
		`<meta[^>]+name=["']` + regexp.QuoteMeta(name) + `["'][^>]+content=["']([^"']+)["'][^>]*>`,
		`<meta[^>]+content=["']([^"']+)["'][^>]+name=["']` + regexp.QuoteMeta(name) + `["'][^>]*>`,
		`<meta[^>]+property=["']og:` + regexp.QuoteMeta(name) + `["'][^>]+content=["']([^"']+)["'][^>]*>`,
		`<meta[^>]+content=["']([^"']+)["'][^>]+property=["']og:` + regexp.QuoteMeta(name) + `["'][^>]*>`,
	}
	for _, pattern := range patterns {
		if value := firstMatch(input, pattern); value != "" {
			return value
		}
	}
	return ""
}

func htmlToText(input string, maxChars int) string {
	replacer := strings.NewReplacer(
		"<br>", "\n", "<br/>", "\n", "<br />", "\n",
		"</p>", "\n\n", "</li>", "\n", "</h1>", "\n\n", "</h2>", "\n\n", "</h3>", "\n\n",
	)
	input = replacer.Replace(input)
	re := regexp.MustCompile(`(?is)<[^>]+>`)
	text := re.ReplaceAllString(input, " ")
	text = html.UnescapeString(text)
	return trimRunes(cleanWhitespace(text), maxChars)
}

func cleanWhitespace(input string) string {
	lines := strings.Split(input, "\n")
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.Join(strings.Fields(line), " ")
		if line != "" {
			out = append(out, line)
		}
	}
	return strings.TrimSpace(strings.Join(out, "\n"))
}

func trimRunes(input string, max int) string {
	if max <= 0 || len([]rune(input)) <= max {
		return input
	}
	runes := []rune(input)
	return strings.TrimSpace(string(runes[:max])) + "\n\n[truncated]"
}

type googleTranslator struct {
	client *http.Client
}

func (g *googleTranslator) translate(ctx context.Context, text string) (string, error) {
	chunks := splitForTranslate(text, 1800)
	out := make([]string, 0, len(chunks))
	for _, chunk := range chunks {
		translated, err := g.translateChunk(ctx, chunk)
		if err != nil {
			return "", err
		}
		out = append(out, translated)
	}
	return strings.Join(out, "\n\n"), nil
}

func (g *googleTranslator) translateChunk(ctx context.Context, text string) (string, error) {
	params := url.Values{}
	params.Set("client", "gtx")
	params.Set("sl", "auto")
	params.Set("tl", "ja")
	params.Set("dt", "t")
	params.Set("q", text)

	endpoint := "https://translate.googleapis.com/translate_a/single?" + params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := g.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("google translate status %s: %s", resp.Status, strings.TrimSpace(string(data)))
	}

	translated, err := parseGoogleTranslateResponse(data)
	if err != nil {
		return "", err
	}
	return html.UnescapeString(translated), nil
}

func parseGoogleTranslateResponse(data []byte) (string, error) {
	var parsed []any
	if err := json.Unmarshal(data, &parsed); err != nil {
		return "", err
	}
	if len(parsed) == 0 {
		return "", errors.New("google translate returned no translations")
	}

	sentences, ok := parsed[0].([]any)
	if !ok || len(sentences) == 0 {
		return "", errors.New("google translate returned no translations")
	}
	var b strings.Builder
	for _, rawSentence := range sentences {
		sentence, ok := rawSentence.([]any)
		if !ok || len(sentence) == 0 {
			continue
		}
		part, ok := sentence[0].(string)
		if ok {
			b.WriteString(part)
		}
	}
	translated := strings.TrimSpace(b.String())
	if translated == "" {
		return "", errors.New("google translate returned empty translation")
	}
	return translated, nil
}

func splitForTranslate(text string, maxBytes int) []string {
	if len(text) <= maxBytes {
		return []string{text}
	}
	paragraphs := strings.Split(text, "\n\n")
	var chunks []string
	var b bytes.Buffer
	flush := func() {
		if b.Len() == 0 {
			return
		}
		chunks = append(chunks, strings.TrimSpace(b.String()))
		b.Reset()
	}
	for _, p := range paragraphs {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if len(p) > maxBytes {
			flush()
			chunks = append(chunks, splitLongString(p, maxBytes)...)
			continue
		}
		if b.Len() > 0 && b.Len()+len(p)+2 > maxBytes {
			flush()
		}
		if b.Len() > 0 {
			b.WriteString("\n\n")
		}
		b.WriteString(p)
	}
	flush()
	return chunks
}

func splitLongString(input string, maxBytes int) []string {
	var chunks []string
	var b bytes.Buffer
	for _, r := range input {
		runeLen := utf8.RuneLen(r)
		if b.Len() > 0 && b.Len()+runeLen > maxBytes {
			chunks = append(chunks, b.String())
			b.Reset()
		}
		b.WriteRune(r)
	}
	if b.Len() > 0 {
		chunks = append(chunks, b.String())
	}
	return chunks
}

func renderMarkdown(now time.Time, item hnItem, article article, translated string) string {
	var b strings.Builder
	b.WriteString("---\n")
	writeYAMLString(&b, "source", item.URL)
	writeYAMLString(&b, "hn_url", fmt.Sprintf(hnItemURLFormat, item.ID))
	writeYAMLString(&b, "title", item.Title)
	writeYAMLString(&b, "article_title", article.Title)
	writeYAMLString(&b, "author", item.By)
	writeYAMLString(&b, "captured_at", now.Format(time.RFC3339))
	writeYAMLString(&b, "capture_tool", "hn-digest")
	fmt.Fprintf(&b, "hn_id: %d\n", item.ID)
	fmt.Fprintf(&b, "score: %d\n", item.Score)
	fmt.Fprintf(&b, "comments: %d\n", item.Descendants)
	fmt.Fprintf(&b, "posted_at: %q\n", time.Unix(item.Time, 0).UTC().Format(time.RFC3339))
	b.WriteString("tags:\n")
	b.WriteString("  - hacker-news\n")
	b.WriteString("  - translated\n")
	b.WriteString("---\n\n")

	fmt.Fprintf(&b, "# %s\n\n", item.Title)
	fmt.Fprintf(&b, "- HN: [%d](%s)\n", item.ID, fmt.Sprintf(hnItemURLFormat, item.ID))
	if item.URL != "" {
		fmt.Fprintf(&b, "- Source: [%s](%s)\n", displayHost(item.URL), item.URL)
	}
	fmt.Fprintf(&b, "- Score: %d\n", item.Score)
	fmt.Fprintf(&b, "- Comments: %d\n", item.Descendants)
	fmt.Fprintf(&b, "- Posted: %s\n", time.Unix(item.Time, 0).UTC().Format(time.RFC3339))

	b.WriteString("\n## Translation\n\n")
	b.WriteString(strings.TrimSpace(translated))
	b.WriteString("\n")

	if article.Description != "" || article.Text != "" || item.Text != "" {
		b.WriteString("\n## Original Extract\n\n")
		if article.Description != "" {
			fmt.Fprintf(&b, "%s\n\n", article.Description)
		}
		if item.Text != "" {
			fmt.Fprintf(&b, "%s\n\n", htmlToText(item.Text, 4000))
		}
		if article.Text != "" {
			b.WriteString(article.Text)
			b.WriteString("\n")
		}
	}
	return b.String()
}

func writeYAMLString(b *strings.Builder, key, value string) {
	if value == "" {
		fmt.Fprintf(b, "%s: \"\"\n", key)
		return
	}
	fmt.Fprintf(b, "%s: %q\n", key, value)
}

func displayHost(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil || parsed.Host == "" {
		return rawURL
	}
	return parsed.Host
}

func uniqueFilename(title string, id int) string {
	name := slug(title)
	if name == "" {
		name = "hn-" + strconv.Itoa(id)
	}
	hash := sha1.Sum([]byte(strconv.Itoa(id)))
	return fmt.Sprintf("%s-%s.md", name, hex.EncodeToString(hash[:])[:8])
}

func slug(input string) string {
	input = strings.TrimSpace(input)
	var b strings.Builder
	lastDash := false
	for _, r := range input {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
			lastDash = false
		case r == '.' || r == '_' || r == '-':
			b.WriteRune(r)
			lastDash = r == '-'
		default:
			if !lastDash {
				b.WriteRune('-')
				lastDash = true
			}
		}
	}
	out := strings.Trim(b.String(), ".-_ ")
	for strings.Contains(out, "--") {
		out = strings.ReplaceAll(out, "--", "-")
	}
	runes := []rune(out)
	if len(runes) > 90 {
		out = strings.Trim(string(runes[:90]), ".-_ ")
	}
	return out
}
