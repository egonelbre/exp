package sortedset

import (
	"math/rand"
	"testing"
)

func newBenchSet() *Set {
	set := New()
	base := 0
	rand.Seed(0)
	for i := 0; i < 1e4; i += 1 {
		base += rand.Intn(15) + 1
		set.Add(base)
	}
	return set
}

func BenchmarkChanUnbuffered(b *testing.B) {
	set := newBenchSet()
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		for v := range set.IterChan(0) {
			total += v
		}
	}
	_ = total
}

func BenchmarkBuffered8(b *testing.B) {
	set := newBenchSet()
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		for v := range set.IterChan(8) {
			total += v
		}
	}
	_ = total
}

func BenchmarkBuffered16(b *testing.B) {
	set := newBenchSet()
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		for v := range set.IterChan(16) {
			total += v
		}
	}
	_ = total
}

func BenchmarkBuffered32(b *testing.B) {
	set := newBenchSet()
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		for v := range set.IterChan(32) {
			total += v
		}
	}
	_ = total
}

func BenchmarkBuffered64(b *testing.B) {
	set := newBenchSet()
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		for v := range set.IterChan(64) {
			total += v
		}
	}
	_ = total
}

func BenchmarkSlice(b *testing.B) {
	set := newBenchSet()
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		for _, v := range set.IterSlice() {
			total += v
		}
	}
	_ = total
}

func BenchmarkCallback(b *testing.B) {
	set := newBenchSet()
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		set.IterCallback(func(v int) {
			total += v
		})
	}
	_ = total
}

func BenchmarkBlockCallback8(b *testing.B) {
	set := newBenchSet()
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		set.IterBlockCallback(8, func(vs []int) {
			for v := range vs {
				total += v
			}
		})
	}
	_ = total
}

func BenchmarkBlockCallback16(b *testing.B) {
	set := newBenchSet()
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		set.IterBlockCallback(16, func(vs []int) {
			for v := range vs {
				total += v
			}
		})
	}
	_ = total
}

func BenchmarkBlockCallback32(b *testing.B) {
	set := newBenchSet()
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		set.IterBlockCallback(32, func(vs []int) {
			for v := range vs {
				total += v
			}
		})
	}
	_ = total
}

func BenchmarkBlockCallback64(b *testing.B) {
	set := newBenchSet()
	b.ResetTimer()
	total := 0
	for i := 0; i < b.N; i++ {
		set.IterBlockCallback(64, func(vs []int) {
			for v := range vs {
				total += v
			}
		})
	}
	_ = total
}
