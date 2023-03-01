package maplookup_test

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"
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

type LookupMap map[string]struct{}

func (q LookupMap) Contains(x string) bool {
	_, ok := q[x]
	return ok
}

func lookupMap(entries []string) LookupMap {
	xs := make(map[string]struct{})
	for _, x := range entries {
		q := "x" + x + "x"
		xs[q[1:len(q)-1]] = struct{}{}
	}
	return LookupMap(xs)
}

var randomStrings []string
var sizes = []int{1, 4, 8, 10, 12, 14, 16, 18, 24}

const maxLookup = 10000

func init() {
	for i := 0; i < maxLookup; i++ {
		var s [128]byte
		rand.Read(s[:])
		randomStrings = append(randomStrings, string(s[:]))
	}
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

var z bool

func Benchmark(b *testing.B) {
	for _, size := range sizes {
		targets := randomStrings[:size]
		mismatches := randomStrings[size:]

		list := lookupList(targets)
		mapx := lookupMap(targets)

		for _, scenario := range scenarios(targets, mismatches) {
			b.Run(fmt.Sprintf("Lis/%d-%s", size, scenario.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, q := range scenario.query {
						z = list.Contains(q)
					}
				}
			})
			b.Run(fmt.Sprintf("Map/%d-%s", size, scenario.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, q := range scenario.query {
						z = mapx.Contains(q)
					}
				}
			})
		}
	}
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
	rand.Shuffle(len(all), func(i, k int) {
		all[i], all[k] = all[k], all[i]
	})
	return all
}
