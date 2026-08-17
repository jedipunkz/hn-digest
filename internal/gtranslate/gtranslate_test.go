package gtranslate

import (
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
