package mispredict

import (
	"math/rand"
	"sort"
	"testing"
)

var input = func() []int {
	input := make([]int, 1e4)
	for i := range input {
		input[i] = rand.Intn(100)
	}
	return input
}()

var sorted = func() []int {
	sorted := make([]int, len(input))
	copy(sorted, input)
	sort.Ints(sorted)
	return sorted
}()

func BenchmarkRandom(b *testing.B) {
	total := 0
	for k := 0; k < b.N; k++ {
		total += sum(input)
	}
}

func BenchmarkSorted(b *testing.B) {
	total := 0
	for k := 0; k < b.N; k++ {
		total += sum(sorted)
	}
}

//go:noinline
func fn(v int) int {
	if v < 50 {
		return 0
	} else {
		return 1
	}
}

func sum(vs []int) (total int) {
	for _, v := range vs {
		total += fn(v)
	}
	return total
}
