package llm

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func testClient(t *testing.T, h http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	return &Client{
		APIKey:   "test-key",
		Model:    "gemini-test",
		Endpoint: srv.URL,
		HTTP:     srv.Client(),
		Retries:  1,
		Backoff:  time.Millisecond,
	}
}

func TestSummarize(t *testing.T) {
	var gotPath, gotKey, gotBody string
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotKey = r.Header.Get("x-goog-api-key")
		b, _ := io.ReadAll(r.Body)
		gotBody = string(b)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"candidates":[{"content":{"parts":[{"text":"- 要約の一文目。\n\n二文目。"}]}}]}`)
	})

	got, err := c.Summarize(context.Background(), "A story", "body text")
	if err != nil {
		t.Fatalf("Summarize() error: %v", err)
	}
	// Bullets and blank lines are flattened: the front matter and the RSS
	// description both want one line of prose.
	if want := "要約の一文目。 二文目。"; got != want {
		t.Fatalf("Summarize() = %q, want %q", got, want)
	}
	if want := "/models/gemini-test:generateContent"; gotPath != want {
		t.Fatalf("path = %q, want %q", gotPath, want)
	}
	if gotKey != "test-key" {
		t.Fatalf("api key header = %q", gotKey)
	}
	if !strings.Contains(gotBody, "A story") || !strings.Contains(gotBody, "body text") {
		t.Fatalf("request body missing title or body: %s", gotBody)
	}
}

func TestSummarizeRetriesRateLimit(t *testing.T) {
	calls := 0
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls == 1 {
			w.WriteHeader(http.StatusTooManyRequests)
			io.WriteString(w, `{"error":{"code":429,"message":"quota"}}`)
			return
		}
		io.WriteString(w, `{"candidates":[{"content":{"parts":[{"text":"ok"}]}}]}`)
	})

	got, err := c.Summarize(context.Background(), "t", "b")
	if err != nil {
		t.Fatalf("Summarize() error: %v", err)
	}
	if got != "ok" || calls != 2 {
		t.Fatalf("got %q after %d calls, want %q after 2", got, calls, "ok")
	}
}

func TestSummarizeDoesNotRetryClientError(t *testing.T) {
	calls := 0
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"error":{"code":401,"message":"bad key"}}`)
	})

	if _, err := c.Summarize(context.Background(), "t", "b"); err == nil {
		t.Fatal("Summarize() succeeded, want error")
	}
	// A bad key never fixes itself; retrying it 30 times only stalls the run.
	if calls != 1 {
		t.Fatalf("calls = %d, want 1", calls)
	}
}

func TestSummarizeEmptyCandidate(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"candidates":[{"content":{"parts":[{"text":"  "}]}}]}`)
	})
	if _, err := c.Summarize(context.Background(), "t", "b"); err == nil {
		t.Fatal("Summarize() succeeded on an empty candidate, want error")
	}
}

func TestSummarizeNilClientAndEmptyBody(t *testing.T) {
	var c *Client
	if _, err := c.Summarize(context.Background(), "t", "b"); err == nil {
		t.Fatal("Summarize() on nil client succeeded, want error")
	}
	if _, err := testClient(t, func(http.ResponseWriter, *http.Request) {}).Summarize(context.Background(), "t", "   "); err == nil {
		t.Fatal("Summarize() on empty body succeeded, want error")
	}
}

func TestFromEnv(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "")
	if FromEnv() != nil {
		t.Fatal("FromEnv() returned a client without an API key")
	}
	t.Setenv("GEMINI_API_KEY", "k")
	t.Setenv("GEMINI_MODEL", "")
	c := FromEnv()
	if c == nil || c.Model != defaultModel {
		t.Fatalf("FromEnv() = %+v, want default model %q", c, defaultModel)
	}
	t.Setenv("GEMINI_MODEL", "custom-model")
	if c := FromEnv(); c.Model != "custom-model" {
		t.Fatalf("model = %q, want custom-model", c.Model)
	}
}
