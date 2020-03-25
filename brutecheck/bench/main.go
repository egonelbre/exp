package main

import (
	"fmt"
	"runtime"

	"github.com/loov/hrtime"
	"gonum.org/v1/gonum/stat"
)

var Valid = []string{
	"break", "case", "chan", "const", "continue",
	"default", "defer", "else", "fallthrough", "for", "func", "go",
	"goto", "if", "import", "interface", "map", "package", "range", "return",
	"select", "struct", "switch", "type", "var",
}

func main() {
	runtime.LockOSThread()

	const repeat = 20000
	const rounds = 20

	type result struct {
		name string
		ns   [rounds]float64
	}

	results := make([]result, len(AllBenches))
	for i, bench := range AllBenches {
		results[i].name = bench.Name
	}

	for round := 0; round < rounds; round++ {
		for i, bench := range AllBenches {
			fn := bench.Fn
			fn(Valid, 1)

			start := hrtime.TSC()
			fn(Valid, repeat)
			finish := hrtime.TSC()

			totalns := float64((finish - start).ApproxDuration().Nanoseconds())
			avgns := totalns / float64(repeat*len(Valid))

			results[i].ns[round] = avgns
		}
	}

	fmt.Printf("func,mean,stddev")
	for round := 0; round < rounds; round++ {
		fmt.Printf(",#%d ns", round)
	}
	fmt.Println()

	for _, r := range results {
		fmt.Printf("%s,", r.name)
		mean, stddev := stat.MeanStdDev(r.ns[:], nil)
		fmt.Printf("%.2f, %.2f", mean, stddev)
		for _, ns := range r.ns {
			fmt.Printf(",%.2f", ns)
		}
		fmt.Println()
	}
}
