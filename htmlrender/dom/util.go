package dom

import (
	"bytes"
	"fmt"
	"html"
	"strings"
)

func EscapeCharData(s string) string {
	s = html.EscapeString(s)
	s = strings.Replace(s, "&#39;", "'", -1)
	s = strings.Replace(s, "&#34;", "\"", -1)
	s = strings.Replace(s, " ", "&nbsp;", -1)
	return s
}

func EscapeAttribute(s string) string {
	s = html.EscapeString(s)
	s = strings.Replace(s, "&#39;", "'", -1)
	// s = strings.Replace(s, "&#34;", "\"", -1)
	s = strings.Replace(s, " ", "&nbsp;", -1)
	return s
}

func EscapeString(s string) string {
	s = html.EscapeString(s)
	// s = strings.Replace(s, "&#39;", "'", -1)
	// s = strings.Replace(s, "&#34;", "\"", -1)
	s = strings.Replace(s, " ", "&nbsp;", -1)
	return s
}

// NormalizeURL normalizes url to be safely included as an href
// based on golang.org/pkg/html/template
func NormalizeURL(s string) string {
	if i := strings.IndexRune(s, ':'); i >= 0 && strings.IndexRune(s[:i], '/') < 0 {
		protocol := strings.ToLower(s[:i])
		if protocol != "http" && protocol != "https" && protocol != "mailto" {
			return "#ZurlZ"
		}
	}

	var b bytes.Buffer
	written := 0
	// The byte loop below assumes that all URLs use UTF-8 as the
	// content-encoding. This is similar to the URI to IRI encoding scheme
	// defined in section 3.1 of  RFC 3987, and behaves the same as the
	// EcmaScript builtin encodeURIComponent.
	// It should not cause any misencoding of URLs in pages with
	// Content-type: text/html;charset=UTF-8.
	for i, n := 0, len(s); i < n; i++ {
		c := s[i]
		switch c {
		// Single quote and parens are sub-delims in RFC 3986, but we
		// escape them so the output can be embedded in single
		// quoted attributes and unquoted CSS url(...) constructs.
		// Single quotes are reserved in URLs, but are only used in
		// the obsolete "mark" rule in an appendix in RFC 3986
		// so can be safely encoded.
		case '!', '#', '$', '&', '*', '+', ',', '/', ':', ';', '=', '?', '@', '[', ']':
			continue
		// Unreserved according to RFC 3986 sec 2.3
		// "For consistency, percent-encoded octets in the ranges of
		// ALPHA (%41-%5A and %61-%7A), DIGIT (%30-%39), hyphen (%2D),
		// period (%2E), underscore (%5F), or tilde (%7E) should not be
		// created by URI producers
		case '-', '.', '_', '~':
			continue
		case '%':
			// When normalizing do not re-encode valid escapes.
			if i+2 < len(s) && isHex(s[i+1]) && isHex(s[i+2]) {
				continue
			}
		default:
			// Unreserved according to RFC 3986 sec 2.3
			if 'a' <= c && c <= 'z' {
				continue
			}
			if 'A' <= c && c <= 'Z' {
				continue
			}
			if '0' <= c && c <= '9' {
				continue
			}
		}
		b.WriteString(s[written:i])
		fmt.Fprintf(&b, "%%%02x", c)
		written = i + 1
	}
	if written == 0 {
		return s
	}
	b.WriteString(s[written:])
	return b.String()
}

func isHex(c byte) bool {
	return '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F'
}
