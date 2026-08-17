package main

import (
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

func TestFirstImageURL(t *testing.T) {
	base, err := url.Parse("https://example.com/posts/one")
	if err != nil {
		t.Fatal(err)
	}

	// og:image wins over twitter:image, and &amp; is unescaped so the query
	// string stays usable.
	body := `<meta name="twitter:image" content="https://cdn.example.com/t.png">
	<meta property="og:image" content="https://cdn.example.com/a.png?w=1&amp;sig=xyz">`
	want := "https://cdn.example.com/a.png?w=1&sig=xyz"
	if got := firstImageURL(body, base); got != want {
		t.Fatalf("firstImageURL() = %q, want %q", got, want)
	}

	// Relative paths resolve against the page URL.
	if got := firstImageURL(`<meta property="og:image" content="/img/a.png">`, base); got != "https://example.com/img/a.png" {
		t.Fatalf("firstImageURL(relative) = %q", got)
	}

	// Reversed attribute order, and twitter:image as the only candidate.
	if got := firstImageURL(`<meta content="https://cdn.example.com/t.png" name="twitter:image"/>`, base); got != "https://cdn.example.com/t.png" {
		t.Fatalf("firstImageURL(twitter) = %q", got)
	}

	// Non-http schemes and pages without an image yield "".
	if got := firstImageURL(`<meta property="og:image" content="data:image/png;base64,AAAA">`, base); got != "" {
		t.Fatalf("firstImageURL(data URI) = %q, want empty", got)
	}
	if got := firstImageURL(`<meta property="og:image:width" content="1200">`, base); got != "" {
		t.Fatalf("firstImageURL(no image) = %q, want empty", got)
	}
}

func TestSlug(t *testing.T) {
	got := slug(`Show HN: Go + GitHub Actions / "digest"?`)
	want := "Show-HN-Go-GitHub-Actions-digest"
	if got != want {
		t.Fatalf("slug() = %q, want %q", got, want)
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

func TestTitleMatchesKeywords(t *testing.T) {
	keywords := parseKeywords("sre,devops,google cloud,gcp,ai,llm")
	tests := []struct {
		title string
		want  bool
	}{
		{"Running PostgreSQL on Google Cloud SQL", true},
		{"Show HN: AI code review for pull requests", true},
		{"A guide to SRE incident reviews", true},
		{"Avoiding said-bookisms in technical writing", false},
		{"A small compiler written in C", false},
	}
	for _, tt := range tests {
		got, _ := titleMatchesKeywords(tt.title, keywords)
		if got != tt.want {
			t.Fatalf("titleMatchesKeywords(%q) = %v, want %v", tt.title, got, tt.want)
		}
	}
}
