package ogimage

import (
	"net/url"
	"testing"
)

func TestFromHTML(t *testing.T) {
	base, err := url.Parse("https://example.com/posts/one")
	if err != nil {
		t.Fatal(err)
	}

	// og:image wins over twitter:image, and &amp; is unescaped so the query
	// string stays usable.
	body := `<meta name="twitter:image" content="https://cdn.example.com/t.png">
	<meta property="og:image" content="https://cdn.example.com/a.png?w=1&amp;sig=xyz">`
	want := "https://cdn.example.com/a.png?w=1&sig=xyz"
	if got := FromHTML(body, base); got != want {
		t.Fatalf("FromHTML() = %q, want %q", got, want)
	}

	// Relative paths resolve against the page URL.
	if got := FromHTML(`<meta property="og:image" content="/img/a.png">`, base); got != "https://example.com/img/a.png" {
		t.Fatalf("FromHTML(relative) = %q", got)
	}

	// Reversed attribute order, and twitter:image as the only candidate.
	if got := FromHTML(`<meta content="https://cdn.example.com/t.png" name="twitter:image"/>`, base); got != "https://cdn.example.com/t.png" {
		t.Fatalf("FromHTML(twitter) = %q", got)
	}

	// Non-http schemes and pages without an image yield "".
	if got := FromHTML(`<meta property="og:image" content="data:image/png;base64,AAAA">`, base); got != "" {
		t.Fatalf("FromHTML(data URI) = %q, want empty", got)
	}
	if got := FromHTML(`<meta property="og:image:width" content="1200">`, base); got != "" {
		t.Fatalf("FromHTML(no image) = %q, want empty", got)
	}
}
