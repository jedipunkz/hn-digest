package frontmatter

import "testing"

const doc = `---
source: "https://example.com/a"
title: "Kubemend – an agent that never trusts its own \"fixed\""
hn_id: 49328606
score: 2
---
`

func TestString(t *testing.T) {
	// %q-quoted values may embed escaped quotes; they must survive intact.
	want := `Kubemend – an agent that never trusts its own "fixed"`
	if got := String(doc, "title"); got != want {
		t.Fatalf("String(title) = %q, want %q", got, want)
	}
	if got := String(doc, "source"); got != "https://example.com/a" {
		t.Fatalf("String(source) = %q", got)
	}
	if got := String(doc, "missing"); got != "" {
		t.Fatalf("String(missing) = %q, want empty", got)
	}
}

func TestInt(t *testing.T) {
	if got := Int(doc, "hn_id"); got != 49328606 {
		t.Fatalf("Int(hn_id) = %d", got)
	}
	if got := Int(doc, "title"); got != 0 {
		t.Fatalf("Int(title) = %d, want 0", got)
	}
}
