package fuzzytext

import (
	"unicode"
	"unicode/utf8"
)

func MatchSimple(p, s string) (match bool) {
	pch, pn := utf8.DecodeRuneInString(p)
	if pch == utf8.RuneError {
		return false
	}
	p = p[pn:]
	pch = unicode.ToLower(pch)

	for _, sch := range s {
		sch = unicode.ToLower(sch)
		if sch != pch {
			continue
		}

		pch, pn = utf8.DecodeRuneInString(p)
		if pch == utf8.RuneError {
			return pn == 0
		}
		p = p[pn:]
		pch = unicode.ToLower(pch)
	}

	return false
}
