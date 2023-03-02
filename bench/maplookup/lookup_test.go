package maplookup_test

import (
	_ "embed"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"

	"github.com/egonelbre/exp/wordsearch/trie-compact"
)

type LookupList []string

func (q LookupList) Contains(x string) bool {
	for _, t := range q {
		if t == x {
			return true
		}
	}
	return false
}

type BinaryList []string

func (x BinaryList) Contains(target string) bool {
	n := len(x)
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if x[h] < target {
			i = h + 1
		} else {
			j = h
		}
	}

	return i < n && x[i] == target
}

func lookupList(entries []string) LookupList {
	sort.Strings(entries)
	full := strings.Join(entries, "")
	list := []string{}
	for _, s := range entries {
		list = append(list, full[:len(s)])
		full = full[len(s):]
	}
	return LookupList(list)
}

func binaryList(entries []string) BinaryList {
	return BinaryList(lookupList(entries))
}

type LookupMap map[string]struct{}

func (q LookupMap) Contains(x string) bool {
	_, ok := q[x]
	return ok
}

func lookupTrie(entries []string) *trie.Compact {
	root := trie.Uncompact{}
	for _, x := range entries {
		root.Insert(x)
	}
	return root.Compress()
}

func lookupMap(entries []string) LookupMap {
	xs := make(map[string]struct{})
	for _, x := range entries {
		q := "x" + x + "x"
		xs[q[1:len(q)-1]] = struct{}{}
	}
	return LookupMap(xs)
}

func lookupMaph(entries []string) LookupMap {
	xs := make(map[string]struct{}, len(entries))
	for _, x := range entries {
		q := "x" + x + "x"
		xs[q[1:len(q)-1]] = struct{}{}
	}
	return LookupMap(xs)
}

var randomStrings []string
var sizes = []int{1, 4, 6, 8, 10, 12, 14, 16, 18, 24, 28, 32}

const maxLookup = 10000

//go:embed enable1.txt
var dictionary []byte

func init() {
	randomStrings = strings.Split(string(dictionary), "\n")
	shuffleStrings(randomStrings)
}

const queryCount = 1000

type scenario struct {
	name  string
	query []string
}

func scenarios(targets, mismatches []string) []scenario {
	return []scenario{
		{
			name:  "100%",
			query: pickRandom(targets, queryCount),
		},
		{
			name: "90%",
			query: concat(
				pickRandom(targets, queryCount*90/100),
				pickRandom(mismatches, queryCount*10/100),
			),
		},
		{
			name: "50%",
			query: concat(
				pickRandom(targets, queryCount*50/100),
				pickRandom(mismatches, queryCount*50/100),
			),
		},
		{
			name: "10%",
			query: concat(
				pickRandom(targets, queryCount*10/100),
				pickRandom(mismatches, queryCount*90/100),
			),
		},
	}
}

var z int

func Benchmark(b *testing.B) {
	for _, size := range sizes {
		targets := randomStrings[:size]
		mismatches := randomStrings[size:]

		lisx := lookupList(targets)
		binx := binaryList(targets)
		trix := lookupTrie(targets)
		mapx := lookupMap(targets)
		maphx := lookupMaph(targets)

		for _, scenario := range scenarios(targets, mismatches) {
			b.Run(fmt.Sprintf("Lis/%d-%s", size, scenario.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, q := range scenario.query {
						if lisx.Contains(q) {
							z++
						}
					}
				}
			})

			b.Run(fmt.Sprintf("Bin/%d-%s", size, scenario.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, q := range scenario.query {
						if binx.Contains(q) {
							z++
						}
					}
				}
			})

			b.Run(fmt.Sprintf("Tri/%d-%s", size, scenario.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, q := range scenario.query {
						if trix.Contains(q) {
							z++
						}
					}
				}
			})

			b.Run(fmt.Sprintf("Map/%d-%s", size, scenario.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, q := range scenario.query {
						if mapx.Contains(q) {
							z++
						}
					}
				}
			})

			b.Run(fmt.Sprintf("Maph/%d-%s", size, scenario.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, q := range scenario.query {
						if maphx.Contains(q) {
							z++
						}
					}
				}
			})
		}
	}

	b.Log(z)
}

func concat(xs ...[]string) []string {
	all := []string{}
	for _, x := range xs {
		all = append(all, x...)
	}
	return all
}

func pickRandom(xs []string, n int) []string {
	all := append([]string{}, xs...)
	for len(all) < n {
		all = append(all, all...)
	}
	all = all[:n]
	shuffleStrings(all)
	return all
}

func shuffleStrings(xs []string) {
	rand.Shuffle(len(xs), func(i, k int) {
		xs[i], xs[k] = xs[k], xs[i]
	})
}
