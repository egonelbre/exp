package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/egonelbre/exp/wordsearch/trie-compact"
)

func main() {
	root := trie.Uncompact{}

	r, err := os.Open("wordlist.txt")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || utf8.RuneCountInString(line) == 1 {
			continue
		}
		if strings.ToLower(line) != line {
			continue
		}
		root.Insert(line)
	}

	compact := root.Compress()
	fmt.Println(compact)
}
