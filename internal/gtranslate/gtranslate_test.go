package gtranslate

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSplitForTranslate(t *testing.T) {
	got := SplitForTranslate(strings.Repeat("a", 12), 5)
	if len(got) != 3 {
		t.Fatalf("len(SplitForTranslate()) = %d, want 3: %#v", len(got), got)
	}
	for _, chunk := range got {
		if len(chunk) > 5 {
			t.Fatalf("chunk too large: %q", chunk)
		}
	}
}

func TestParseResponse(t *testing.T) {
	data := []byte(`[[["こんにちは","Hello",null,null,10],["世界"," world",null,null,10]],null,"en"]`)
	got, err := ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if got != "こんにちは世界" {
		t.Fatalf("ParseResponse() = %q", got)
	}
}

// A rate-limited chunk must not discard the chunks that did translate: the
// endpoint throttles per source IP and an average article is several chunks.
func TestTranslatePartialFailureKeepsOriginalChunk(t *testing.T) {
	const ok = `[[["こんにちは","hello",null,null,10]],null,"en"]`
	var calls int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls == 1 {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		fmt.Fprint(w, ok)
	}))
	defer srv.Close()

	tr := &Translator{Client: srv.Client(), Endpoint: srv.URL}
	first, second := strings.Repeat("a", 1000), strings.Repeat("b", 1000)
	got, err := tr.Translate(context.Background(), first+"\n\n"+second)
	if err != nil {
		t.Fatalf("Translate() error = %v, want nil", err)
	}
	if want := first + "\n\nこんにちは"; got != want {
		t.Fatalf("Translate() = %q, want %q", got, want)
	}
}

func TestTranslateAllChunksFailReturnsError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
	}))
	defer srv.Close()

	tr := &Translator{Client: srv.Client(), Endpoint: srv.URL}
	if _, err := tr.Translate(context.Background(), "hello"); err == nil {
		t.Fatal("Translate() error = nil, want an error when every chunk fails")
	}
}
