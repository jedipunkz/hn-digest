package main

import (
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
