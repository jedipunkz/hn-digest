package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSlug(t *testing.T) {
	got := slug(`Show HN: Go + GitHub Actions / "digest"?`)
	want := "Show-HN-Go-GitHub-Actions-digest"
	if got != want {
		t.Fatalf("slug() = %q, want %q", got, want)
	}
}

func TestSplitForTranslate(t *testing.T) {
	got := splitForTranslate(strings.Repeat("a", 12), 5)
	if len(got) != 3 {
		t.Fatalf("len(splitForTranslate()) = %d, want 3: %#v", len(got), got)
	}
	for _, chunk := range got {
		if len(chunk) > 5 {
			t.Fatalf("chunk too large: %q", chunk)
		}
	}
}

func TestLoadSeen(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "2026-05-30"), 0o755); err != nil {
		t.Fatal(err)
	}
	content := `---
source: "https://example.com/story"
hn_id: 123
---
`
	if err := os.WriteFile(filepath.Join(root, "2026-05-30", "story.md"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	seen, err := loadSeen(root)
	if err != nil {
		t.Fatal(err)
	}
	if !seen.has(hnItem{ID: 123}) {
		t.Fatal("expected HN ID to be seen")
	}
	if !seen.has(hnItem{ID: 456, URL: "https://example.com/story"}) {
		t.Fatal("expected source URL to be seen")
	}
	if seen.has(hnItem{ID: 789, URL: "https://example.com/other"}) {
		t.Fatal("unexpected duplicate")
	}
}

func TestParseGoogleTranslateResponse(t *testing.T) {
	data := []byte(`[[["こんにちは","Hello",null,null,10],["世界"," world",null,null,10]],null,"en"]`)
	got, err := parseGoogleTranslateResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if got != "こんにちは世界" {
		t.Fatalf("parseGoogleTranslateResponse() = %q", got)
	}
}

func TestMaxLabel(t *testing.T) {
	if got := maxLabel(0); got != "unlimited" {
		t.Fatalf("maxLabel(0) = %q", got)
	}
	if got := maxLabel(50); got != "50" {
		t.Fatalf("maxLabel(50) = %q", got)
	}
}
