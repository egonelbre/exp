package main

import (
	"fmt"
	"sort"
)

type TwoHashShiftTableAlt struct{}

func (TwoHashShiftTableAlt) Generate(p Printer, keywords []string) {
	minlen, _ := keywordBounds(keywords)
	if minlen < 2 {
		return
	}

	for _, comb0 := range combiners {
		for _, comb1 := range combiners {
			for shiftn := byte(0); shiftn < 8; shiftn++ {
				for shift0 := byte(0); shift0 < 8; shift0++ {
				next:
					for shift1 := byte(0); shift1 < 8; shift1++ {
						var seen [256]bool
						hashes := make([]byte, len(keywords))
						for i, keyword := range keywords {
							h := twoHashCalcAlt(keyword, comb0, comb1, shiftn, shift0, shift1)
							if seen[h] {
								continue next
							}
							hashes[i] = h
							seen[h] = true
						}

						sort.Sort(&sortByHash{keywords, hashes})

						fnname := fmt.Sprintf("TwoHashTableAlt_%s%s_Shift%d%d%d", comb0.Name, comb1.Name, shiftn, shift0, shift1)
						p.FuncName(fnname)
						p.F("func %s(name string) bool {\n", fnname)
						p.F("if len(name) < 2 { return false }\n")

						maxhash := hashes[len(hashes)-1]
						p.F("h := %s\n", twoHashFormatAlt(comb0, comb1, shiftn, shift0, shift1))
						p.F("if h > %d { return false }\n", maxhash)

						p.F("return [...]string{\n")
						for i, keyword := range keywords {
							p.F("%d: %q,\n", hashes[i], keyword)
						}
						p.F("}[h] == name\n")
						//p.F("return names[h] == name\n")

						p.F("}\n\n")
					}
				}
			}
		}
	}
}
