package main

import (
	"encoding/binary"
)

type Uint64Switch struct{}

func (Uint64Switch) Generate(p Printer, keywords []string) {
	minlen, _ := keywordBounds(keywords)
	if minlen < 2 {
		return
	}

	p.F("import \"encoding/binary\"\n")

	p.FuncName("Uint64Switch")
	p.F("func Uint64Switch(name string) bool {\n")
	{
		p.F("var key [8]byte\n")
		p.F("copy(key[:], name)\n")
		p.F("h := binary.LittleEndian.Uint64(key[:])\n")

		p.F("switch h {\n")
		for _, keyword := range keywords {
			var key [8]byte
			copy(key[:], keyword)
			h := binary.LittleEndian.Uint64(key[:])

			p.F("case %d:\n", h)
			if len(keyword) <= 8 {
				p.F("return len(name) == %d\n", len(keyword))
			} else {
				p.F("return name[8:] == %q\n", keyword[8:])
			}
		}
		p.F("}\n")
	}
	p.F("return false\n")
	p.F("}\n\n")
}
