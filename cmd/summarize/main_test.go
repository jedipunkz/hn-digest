package main

import "testing"

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
