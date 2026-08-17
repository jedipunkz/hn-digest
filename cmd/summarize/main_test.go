package main

import (
	"strings"
	"testing"
)

const mdWithoutImage = `---
source: "https://example.com/a"
hn_url: "https://news.ycombinator.com/item?id=1"
title: "A story"
article_title: "A story - Example"
hn_id: 1
score: 5
tags:
  - hacker-news
---

# A story

## Original Extract

image: this body line must not be touched
`

func TestWithImage(t *testing.T) {
	// Inserted after article_title, keeping the crawler's field order, and the
	// body line that also starts with "image:" is left alone.
	got, ok := withImage(mdWithoutImage, "https://cdn.example.com/a.png?w=1&s=2")
	if !ok {
		t.Fatal("withImage() reported no front matter")
	}
	if !strings.Contains(got, "article_title: \"A story - Example\"\nimage: \"https://cdn.example.com/a.png?w=1&s=2\"\n") {
		t.Fatalf("image not inserted after article_title:\n%s", got)
	}
	if !strings.Contains(got, "image: this body line must not be touched") {
		t.Fatal("body line was rewritten")
	}

	// Re-running replaces the existing key rather than adding a second one.
	again, ok := withImage(got, "")
	if !ok {
		t.Fatal("withImage() reported no front matter on rewrite")
	}
	frontMatter, _, _ := strings.Cut(strings.TrimPrefix(again, "---\n"), "\n---\n")
	if n := strings.Count(frontMatter, "\nimage: "); n != 1 {
		t.Fatalf("front matter has %d image keys, want 1:\n%s", n, frontMatter)
	}
	if !strings.Contains(again, "image: \"\"\n") {
		t.Fatalf("empty image not recorded:\n%s", again)
	}

	// A file without front matter is left untouched.
	if _, ok := withImage("# no front matter\n", "https://x/y.png"); ok {
		t.Fatal("withImage() accepted a file without front matter")
	}
}

func TestTranslationBody(t *testing.T) {
	full := "タイトル: T\n記事タイトル: A\n説明: D\n\n記事本文:\n本文です。"
	if got := translationBody(full); got != "本文です。" {
		t.Fatalf("translationBody(full) = %q", got)
	}
	if got := translationBody("タイトル: T\n説明: 説明文です。"); got != "説明文です。" {
		t.Fatalf("translationBody(no body) = %q", got)
	}
	if got := translationBody("タイトル: T のみ"); got != "T のみ" {
		t.Fatalf("translationBody(title only) = %q", got)
	}
}

func TestLeadSummary(t *testing.T) {
	short := "短い本文です。"
	if got := leadSummary(short, 900); got != short {
		t.Fatalf("leadSummary(short) = %q, want unchanged", got)
	}

	// Cuts at the last sentence boundary inside the limit, not mid-sentence.
	if got := leadSummary("一文目です。二文目はここで切れる長い文", 15); got != "一文目です。" {
		t.Fatalf("leadSummary() = %q, want %q", got, "一文目です。")
	}

	// No sentence boundary available: ellipsis, and still valid UTF-8.
	got := leadSummary("あいうえおかきくけこ", 5)
	if got != "あいうえお…" {
		t.Fatalf("leadSummary(no boundary) = %q", got)
	}
}
