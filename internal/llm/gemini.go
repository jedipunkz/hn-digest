// Package llm wraps the Gemini API free tier so cmd/summarize can produce an
// abstractive Japanese summary instead of the extractive lead.
//
// Everything here is best effort: the caller is expected to fall back to the
// extractive summary whenever a call fails, so a missing key, a rate limit or
// an outage degrades the digest instead of failing the workflow.
package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// defaultModel is the cheapest/fastest current Gemini tier, which is also the
// one with the most generous free-tier quota. Override with GEMINI_MODEL when
// Google renames or retires it — the name is data, not a code change.
const defaultModel = "gemini-3.5-flash-lite"

const defaultEndpoint = "https://generativelanguage.googleapis.com/v1beta"

// maxPromptChars caps what we send. The translated body is already capped at
// 4000 runes upstream; this is a second belt so a pathological file cannot
// burn the daily token quota on its own.
const maxPromptChars = 6000

// Client talks to the Gemini generateContent endpoint.
type Client struct {
	APIKey   string
	Model    string
	Endpoint string
	HTTP     *http.Client
	// Retries is the number of extra attempts on a retryable status (429, 5xx).
	Retries int
	// Backoff is the wait before the first retry; it doubles each attempt.
	Backoff time.Duration
}

// FromEnv returns a client configured from GEMINI_API_KEY / GEMINI_MODEL, or
// nil when no key is set. A nil client is the "no LLM configured" state and is
// safe to hold: Summarize on it returns an error rather than panicking.
func FromEnv() *Client {
	key := strings.TrimSpace(os.Getenv("GEMINI_API_KEY"))
	if key == "" {
		return nil
	}
	model := strings.TrimSpace(os.Getenv("GEMINI_MODEL"))
	if model == "" {
		model = defaultModel
	}
	return &Client{
		APIKey: key,
		Model:  model,
		// GEMINI_ENDPOINT exists so the API base can be redirected (a mirror,
		// a local stub) without a code change.
		Endpoint: strings.TrimSpace(os.Getenv("GEMINI_ENDPOINT")),
		HTTP:     &http.Client{Timeout: 60 * time.Second},
		Retries:  2,
		Backoff:  5 * time.Second,
	}
}

type genRequest struct {
	Contents         []content         `json:"contents"`
	GenerationConfig *generationConfig `json:"generationConfig,omitempty"`
}

type content struct {
	Parts []part `json:"parts"`
}

type part struct {
	Text string `json:"text"`
}

type generationConfig struct {
	Temperature     float64 `json:"temperature"`
	MaxOutputTokens int     `json:"maxOutputTokens"`
}

type genResponse struct {
	Candidates []struct {
		Content struct {
			Parts []part `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

const promptTemplate = `次の Hacker News 記事を日本語で要約してください。

制約:
- 3文以内、合計 300 文字以内。
- 記事に書かれている事実だけを使い、推測や感想を足さない。
- 「この記事は」などの前置きを書かず、本題から始める。
- 箇条書きや見出しは使わず、地の文で書く。

タイトル: %s

本文:
%s`

// Summarize returns a Japanese summary of the article, or an error the caller
// should treat as "use the extractive fallback".
func (c *Client) Summarize(ctx context.Context, title, body string) (string, error) {
	if c == nil {
		return "", fmt.Errorf("llm: no client configured")
	}
	body = strings.TrimSpace(body)
	if body == "" {
		return "", fmt.Errorf("llm: empty article body")
	}
	if runes := []rune(body); len(runes) > maxPromptChars {
		body = string(runes[:maxPromptChars])
	}

	payload, err := json.Marshal(genRequest{
		Contents: []content{{Parts: []part{{Text: fmt.Sprintf(promptTemplate, title, body)}}}},
		GenerationConfig: &generationConfig{
			Temperature:     0.2,
			MaxOutputTokens: 1024,
		},
	})
	if err != nil {
		return "", err
	}

	backoff := c.Backoff
	if backoff <= 0 {
		backoff = 5 * time.Second
	}
	var lastErr error
	for attempt := 0; attempt <= c.Retries; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return "", ctx.Err()
			case <-time.After(backoff):
			}
			backoff *= 2
		}
		text, retryable, err := c.call(ctx, payload)
		if err == nil {
			return text, nil
		}
		lastErr = err
		if !retryable {
			return "", err
		}
	}
	return "", lastErr
}

func (c *Client) call(ctx context.Context, payload []byte) (text string, retryable bool, err error) {
	endpoint := c.Endpoint
	if endpoint == "" {
		endpoint = defaultEndpoint
	}
	url := fmt.Sprintf("%s/models/%s:generateContent", strings.TrimRight(endpoint, "/"), c.Model)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return "", false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-goog-api-key", c.APIKey)

	client := c.HTTP
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		// Network-level failures are worth one more try.
		return "", true, err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return "", true, err
	}
	if resp.StatusCode != http.StatusOK {
		// 429 is the free-tier rate limit; 5xx is transient. Everything else
		// (401 bad key, 404 unknown model) will not fix itself by retrying.
		retryable = resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500
		return "", retryable, fmt.Errorf("llm: %s: %s", resp.Status, snippet(raw))
	}

	var parsed genResponse
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return "", false, fmt.Errorf("llm: decode response: %w", err)
	}
	if parsed.Error != nil {
		return "", false, fmt.Errorf("llm: %s", parsed.Error.Message)
	}
	var b strings.Builder
	for _, cand := range parsed.Candidates {
		for _, p := range cand.Content.Parts {
			b.WriteString(p.Text)
		}
		break // first candidate only
	}
	// An empty body with 200 happens when the response is filtered or the
	// output budget went entirely to reasoning tokens. Not retryable.
	out := Clean(b.String())
	if out == "" {
		return "", false, fmt.Errorf("llm: empty candidate")
	}
	return out, false, nil
}

// Clean flattens the model output into the single-line plain prose the front
// matter and the RSS description expect.
func Clean(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\r\n", "\n")
	lines := strings.Split(s, "\n")
	kept := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.TrimPrefix(line, "- ")
		line = strings.TrimPrefix(line, "* ")
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		kept = append(kept, line)
	}
	return strings.TrimSpace(strings.Join(kept, " "))
}

func snippet(b []byte) string {
	s := strings.TrimSpace(string(b))
	if len(s) > 200 {
		return s[:200]
	}
	return s
}
