package lcs

import (
	"math/rand"
	"strconv"
	"testing"
)

const (
	N = 160000
	A = 16
)

var (
	dataA = make([]int64, N)
	dataB = make([]int64, N)
	dataC = make([]int64, N)
)

func init() {
	r := rand.New(rand.NewSource(0))
	for i := range dataA {
		dataA[i] = r.Int63n(A)
		dataB[i] = r.Int63n(A)
		dataC[i] = dataA[i]
	}
}
func benchmark(b *testing.B, lcs func(a, b []int64) int) {
	for _, size := range []int{10, 100, 1000, 10000} {
		b.Run("Random"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lcs(dataA[:size], dataB[:size])
			}
		})

		b.Run("Full"+strconv.Itoa(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = lcs(dataA[:size], dataC[:size])
			}
		})
	}
}

func BenchmarkBasic(b *testing.B) { benchmark(b, Basic) }
func BenchmarkLift(b *testing.B)  { benchmark(b, Lift) }
func BenchmarkWave(b *testing.B)  { benchmark(b, Wave) }
