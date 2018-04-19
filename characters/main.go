package main

import (
	"fmt"
	"regexp"

	"github.com/dpinela/charseg"
	"golang.org/x/text/unicode/norm"
)

func ByteCount(s string) int {
	return len(s)
}

func RuneCount(s string) int {
	return len(([]rune)(s))
}

var rxCount = regexp.MustCompile(`\PM\pM*|.`)

func RegexCount(s string) int {
	return len(rxCount.FindAllString(s, -1))
}

func NormalizedCount(form norm.Form, s string) int {
	var i norm.Iter
	i.InitString(form, s)
	var count int
	for !i.Done() {
		i.Next()
		count++
	}
	return count
}

func NFCCount(s string) int  { return NormalizedCount(norm.NFC, s) }
func NFDCount(s string) int  { return NormalizedCount(norm.NFD, s) }
func NFKCCount(s string) int { return NormalizedCount(norm.NFKC, s) }
func NFKDCount(s string) int { return NormalizedCount(norm.NFKD, s) }

func GraphemeCount(s string) int {
	var count int
	for s != "" {
		grapheme := charseg.FirstGraphemeCluster(s)
		count++
		s = s[len(grapheme):]
	}
	return count
}

var examples = []string{
	"hello",
	"你好",
	"hĕllŏ",
	"l̲i̲n̲e̲",
	"ﬁ",
	"ﬃ",
	"㈎",
	"ẛ̣",
}

func main() {
	fmt.Printf("bytes\trunes\tNFC\tNFD\tNFKC\tNFKD\tRegex\tGraph..\tText\n")
	for _, example := range examples {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%q\n",
			ByteCount(example),
			RuneCount(example),
			NFCCount(example),
			NFDCount(example),
			NFKCCount(example),
			NFKDCount(example),
			RegexCount(example),
			GraphemeCount(example),
			example,
		)
	}
}
