package main

type MapCheck struct{}

func (MapCheck) Generate(p Printer, keywords []string) {
	minlen, _ := keywordBounds(keywords)
	if minlen < 2 {
		return
	}

	p.F("var __mapcheck = map[string]struct{}{\n")
	for _, keyword := range keywords {
		p.F("%q: struct{}{},\n", keyword)
	}
	p.F("}\n")

	p.FuncName("MapCheck")
	p.F("func MapCheck(name string) bool {\n")
	p.F("_, ok := __mapcheck[name]\n")
	p.F("return ok\n")

	p.F("}\n\n")
}
