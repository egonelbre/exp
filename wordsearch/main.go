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

	words := []string{}
	r, err := os.Open("wordlist")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || utf8.RuneCountInString(line) == 1 {
			continue
		}
		words = append(words, line)
		root.Insert(line)
	}

	compact := root.Compress()
	fmt.Println(compact.Size(), "bytes")

	for _, word := range words {
		if !compact.Contains(word) {
			fmt.Println("did not find", word)
			break
		}
	}

	fmt.Println(compact.Contains("something"))
	fmt.Println(compact.Contains("NOTHING"))
}
