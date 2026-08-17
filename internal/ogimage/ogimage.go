// Package ogimage extracts an article's social preview image (og:image, then
// twitter:image) from a page's HTML.
package ogimage

import (
	"context"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Browser-ish UA: several publishers serve a bare redirect page to unknown
// clients, and that page carries no og:image.
const userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"

// Preference order. og:image is near-universal on news sites; twitter:image is
// the common fallback.
var metaPatterns = []string{
	`<meta[^>]+property=["']og:image:secure_url["'][^>]+content=["']([^"']+)["']`,
	`<meta[^>]+property=["']og:image["'][^>]+content=["']([^"']+)["']`,
	`<meta[^>]+content=["']([^"']+)["'][^>]+property=["']og:image["']`,
	`<meta[^>]+name=["']twitter:image["'][^>]+content=["']([^"']+)["']`,
	`<meta[^>]+content=["']([^"']+)["'][^>]+name=["']twitter:image["']`,
}

// FromHTML returns the preview image as an absolute URL, or "" when the page
// declares none usable. Relative paths resolve against base, which must be the
// URL that actually served the HTML.
func FromHTML(body string, base *url.URL) string {
	for _, pattern := range metaPatterns {
		re := regexp.MustCompile("(?is)" + pattern)
		match := re.FindStringSubmatch(body)
		if len(match) < 2 {
			continue
		}
		raw := html.UnescapeString(strings.TrimSpace(match[1]))
		if resolved := resolve(raw, base); resolved != "" {
			return resolved
		}
	}
	return ""
}

func resolve(raw string, base *url.URL) string {
	ref, err := url.Parse(raw)
	if err != nil {
		return ""
	}
	if base != nil {
		ref = base.ResolveReference(ref)
	}
	// Feed readers only render http(s) images; skip data: and friends.
	if ref.Scheme != "http" && ref.Scheme != "https" || ref.Host == "" {
		return ""
	}
	return ref.String()
}

// Fetch GETs pageURL and returns its preview image URL. An empty string with a
// nil error means the page was read fine but declares no usable image, which is
// a final answer; a non-nil error is worth retrying later.
func Fetch(ctx context.Context, client *http.Client, pageURL string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pageURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("unexpected status %s", resp.Status)
	}
	if contentType := strings.ToLower(resp.Header.Get("Content-Type")); contentType != "" &&
		!strings.Contains(contentType, "text/html") && !strings.Contains(contentType, "xhtml") {
		// Not an HTML page, so there are no meta tags to read. Final answer.
		return "", nil
	}

	// The <head> is all that matters; 512 KiB covers it without pulling whole pages.
	data, err := io.ReadAll(io.LimitReader(resp.Body, 512<<10))
	if err != nil {
		return "", err
	}
	return FromHTML(string(data), resp.Request.URL), nil
}
