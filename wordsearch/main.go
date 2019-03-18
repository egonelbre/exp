package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/dgryski/go-cuckoof"
	"github.com/egonelbre/exp/wordsearch/trie-compact"
	"github.com/loov/hrtime"
)

func main() {
	root := trie.Uncompact{}

	r, err := os.Open("enable1.txt")
	if err != nil {
		panic(err)
	}

	var words []string
	wordssize := 0

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || utf8.RuneCountInString(line) == 1 {
			continue
		}
		words = append(words, line)
		wordssize += len(line)
		root.Insert(line)
	}

	compact := root.Compress()
	fmt.Printf("serialized %d bytes\n", compact.Size())
	fmt.Printf("%.1f bytes average word\n", float64(wordssize)/float64(len(words)))
	fmt.Printf("%.1f bytes per word\n", float64(compact.Size())/float64(len(words)))
	fmt.Printf("%.1f bytes per key\n", math.Log2(float64(compact.NodeCount()))/8)

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

	BenchmarkBinarySearch(words)
	BenchmarkCuckooFilter(words)
}

func BenchmarkBinarySearch(words []string) {
	start := hrtime.Now()
	for _, word := range words {
		_ = Search(words, word)
	}
	stop := hrtime.Now()
	fmt.Printf("average binary search lookup: %v\n", (stop-start)/time.Duration(len(words)))
}

func Search(words []string, word string) int {
	i, k := 0, len(words)
	for i < k {
		h := int(uint(i+k) >> 1)
		if !(words[h] >= word) {
			i = h + 1
		} else {
			k = h
		}
	}
	return i
}

func BenchmarkCuckooFilter(words []string) {
	filter := cuckoof.New(1 << 19)
	for _, word := range words {
		filter.Insert([]byte(word))
	}

	start := hrtime.Now()
	for _, word := range words {
		if !filter.Lookup([]byte(word)) {
			fmt.Println("did not find", word)
			break
		}
	}
	stop := hrtime.Now()
	fmt.Printf("cuckoo search lookup: %v\n", (stop-start)/time.Duration(len(words)))
}
