package main

import "strings"

// LengthSingle generates:
//
//     func Len1Ch(name string) bool {
//         switch len(name) {
//         case 2: // go, if
//             switch name[0] {
//             case 'g':
//                 return name == "go"
//         ...
type LengthSingle struct{}

func (LengthSingle) Generate(p Printer, keywords []string) {
	bylen := map[int][]string{}
	_, maxlen := keywordBounds(keywords)
	for _, keyword := range keywords {
		n := len(keyword)
		bylen[n] = append(bylen[n], keyword)
	}
	uniquech := map[int]int{}
	for namelen, names := range bylen {
		chi := findUniqueChar(names)
		if chi < 0 {
			return
		}
		uniquech[namelen] = chi
	}

	p.FuncName("Len1Ch")
	p.F("func Len1Ch(name string) bool {\n")
	{
		p.F("switch len(name) {")
		for namelen := 0; namelen < maxlen; namelen++ {
			names, ok := bylen[namelen]
			if !ok {
				continue
			}
			p.F("case %d: // %s\n", namelen, strings.Join(names, ", "))
			if len(names) == 1 {
				p.F("return name == %q\n", names[0])
				continue
			}

			chi := uniquech[namelen]
			p.F("switch name[%d] {\n", chi)
			for _, name := range names {
				p.F("case %q: return name == %q\n", name[chi], name)
			}
			p.F("}\n")
			p.F("return false\n")
		}
		p.F("}\n")
	}
	p.F("return false\n")
	p.F("}\n\n")
}
