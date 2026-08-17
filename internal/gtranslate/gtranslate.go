// Package gtranslate wraps the free (unauthenticated) Google Translate
// endpoint used by both cmd/hn-digest and cmd/summarize.
package gtranslate

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"strings"
	"unicode/utf8"
)

const userAgent = "hn-digest/0.1 (+https://github.com/jedipunkz/hn-digest)"

// Translator translates text into Japanese via translate.googleapis.com.
type Translator struct {
	Client *http.Client
}

// Translate translates text into Japanese, splitting long input into chunks the
// endpoint accepts.
func (t *Translator) Translate(ctx context.Context, text string) (string, error) {
	chunks := SplitForTranslate(text, 1800)
	out := make([]string, 0, len(chunks))
	for _, chunk := range chunks {
		translated, err := t.translateChunk(ctx, chunk)
		if err != nil {
			return "", err
		}
		out = append(out, translated)
	}
	return strings.Join(out, "\n\n"), nil
}

func (t *Translator) translateChunk(ctx context.Context, text string) (string, error) {
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

	resp, err := t.Client.Do(req)
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

	translated, err := ParseResponse(data)
	if err != nil {
		return "", err
	}
	return html.UnescapeString(translated), nil
}

// ParseResponse extracts the translated text from a translate_a/single body.
func ParseResponse(data []byte) (string, error) {
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

// SplitForTranslate splits text into chunks of at most maxBytes, preferring
// paragraph boundaries.
func SplitForTranslate(text string, maxBytes int) []string {
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
