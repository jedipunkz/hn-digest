// Package frontmatter reads values out of the YAML front matter that
// cmd/hn-digest writes into each contents markdown file.
package frontmatter

import (
	"regexp"
	"strconv"
	"strings"
)

// String returns the front matter value for key, or "" when absent.
//
// Values are written with %q, so they are Go-quoted and may contain escaped
// quotes (`title: "a \"b\" c"`). Unquote is the exact inverse; a value that
// does not round-trip falls back to trimming the surrounding quotes.
func String(text, key string) string {
	re := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(key) + `:[ \t]*(.*)$`)
	match := re.FindStringSubmatch(text)
	if len(match) < 2 {
		return ""
	}
	value := strings.TrimSpace(match[1])
	if unquoted, err := strconv.Unquote(value); err == nil {
		return strings.TrimSpace(unquoted)
	}
	return strings.TrimSpace(strings.Trim(value, `"`))
}

// Int returns the front matter value for key as an int, or 0 when absent or
// unparsable.
func Int(text, key string) int {
	n, err := strconv.Atoi(String(text, key))
	if err != nil {
		return 0
	}
	return n
}
