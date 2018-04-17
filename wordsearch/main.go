package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/egonelbre/exp/wordsearch/trie-compact"
	"github.com/loov/hrtime"
)

func main() {
	root := trie.Uncompact{}

	words := []string{}
	r, err := os.Open("enable1.txt")
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

	start := hrtime.Now()
	for _, word := range words {
		if !compact.Contains(word) {
			fmt.Println("did not find", word)
			break
		}
	}
	stop := hrtime.Now()
	fmt.Printf("average lookup: %v\n", (stop-start)/time.Duration(len(words)))

	fmt.Println(compact.Contains("something"))
	fmt.Println(compact.Contains("NOTHING"))
}
