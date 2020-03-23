package main

import (
	"fmt"
	"sort"

	"github.com/loov/hrtime"
)

var Valid = []string{
	"break", "case", "chan", "const", "continue",
	"default", "defer", "else", "fallthrough", "for", "func", "go",
	"goto", "if", "import", "interface", "map", "package", "range", "return",
	"select", "struct", "switch", "type", "var",
}

func main() {
	const repeat = 100000
	type result struct {
		name string
		ns   float64
	}
	results := make([]result, 0, len(AllBenches))

	for _, bench := range AllBenches {
		fn := bench.Fn
		fn(Valid, 1)

		start := hrtime.TSC()
		fn(Valid, repeat)
		finish := hrtime.TSC()

		totalns := float64((finish - start).ApproxDuration().Nanoseconds())
		avgns := totalns / float64(repeat*len(Valid))
		results = append(results, result{bench.Name, avgns})
	}

	sort.Slice(results, func(i, k int) bool {
		return results[i].ns > results[k].ns
	})

	for _, r := range results {
		fmt.Printf("%s\t%.2fns\n", r.name, r.ns)
	}
}
