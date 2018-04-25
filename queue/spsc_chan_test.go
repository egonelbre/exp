package queue_test

import (
	"runtime"
	"strconv"
	"testing"
)

func BenchmarkChanSPSCProdCons(b *testing.B) {
	for _, chanSize := range []int{1, 10, 100} {
		b.Run(strconv.Itoa(chanSize), func(b *testing.B) {
			benchmarkChanSPSCProdCons(b, chanSize, 0)
		})
	}
	for _, chanSize := range []int{1, 10, 100} {
		b.Run("Work"+strconv.Itoa(chanSize), func(b *testing.B) {
			benchmarkChanSPSCProdCons(b, chanSize, 100)
		})
	}
}

func benchmarkChanSPSCProdCons(b *testing.B, chanSize, localWork int) {
	const CallsPerSched = 1000
	procs := runtime.GOMAXPROCS(-1)
	N := int64(b.N / CallsPerSched)

	q := make(chan int64, 2*procs)
	go func() {
		for i := int64(0); i < N; i++ {
			foo := 0
			for g := 0; g < CallsPerSched; g++ {
				for k := 0; k < localWork; k++ {
					foo *= 2
					foo /= 2
				}
				q <- i
			}
		}
	}()

	total := int64(0)
	for i := int64(0); i < N; i++ {
		foo := 0
		for g := 0; g < CallsPerSched; g++ {
			v := <-q
			if v != i {
				b.Fatalf("invalid value %v, expected %v", v, i)
			}
			for k := 0; k < localWork; k++ {
				foo *= 2
				foo /= 2
			}
			total += v
		}
	}

	expected := (N * (N - 1) / 2) * CallsPerSched
	if total != expected {
		b.Fatalf("incorrect total %v, expected %v", total, expected)
	}
}
