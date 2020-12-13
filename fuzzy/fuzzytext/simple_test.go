package fuzzytext_test

import (
	"testing"

	"github.com/egonelbre/exp/fuzzy/fuzzytext"
)

func TestMatchSimple(t *testing.T) {
	var tests = []struct {
		pattern, text string
		match         bool
	}{
		{"", "fuzzy", false},
		{"fts", "", false},
		{"fts", "ft", false},
		{"fts", "fts", true},
		{"f", "fuzzy text search", true},
		{"ft", "fuzzy text search", true},
		{"fts", "fuzzy text search", true},
		{"ftsh", "fuzzy text search", true},
		{"fxz", "fuzzy text search", false},
		{"FTs", "fuzzy Text Search", true},
	}

	for _, test := range tests {
		matched := fuzzytext.MatchSimple(test.pattern, test.text)
		if matched != test.match {
			t.Errorf("MatchSimple(%q, %q) != %v", test.pattern, test.text, test.match)
		}
	}
}
