package queue_test

import (
	"runtime"
	"strconv"
	"testing"
)

func BenchmarkChanMPSCProdCons(b *testing.B) {
	for _, chanSize := range []int{1, 10, 100} {
		b.Run(strconv.Itoa(chanSize), func(b *testing.B) {
			benchmarkChanMPSCProdCons(b, chanSize, 0)
		})
	}
	for _, chanSize := range []int{1, 10, 100} {
		b.Run("Work"+strconv.Itoa(chanSize), func(b *testing.B) {
			benchmarkChanMPSCProdCons(b, chanSize, 100)
		})
	}
}

func benchmarkChanMPSCProdCons(b *testing.B, chanSize, localWork int) {
	const CallsPerSched = 1000
	procs := runtime.GOMAXPROCS(-1)
	N := b.N / CallsPerSched

	q := make(chan int64, 2*procs)
	for i := 0; i < procs; i++ {
		go func() {
			for i := int64(0); i < int64(N); i++ {
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
	}

	total := int64(0)
	for i := int64(0); i < int64(procs*N); i++ {
		foo := 0
		for g := 0; g < CallsPerSched; g++ {
			v := <-q
			for k := 0; k < localWork/procs; k++ {
				foo *= 2
				foo /= 2
			}
			total += v
		}
	}
	expected := int64(procs * (N * (N - 1) / 2) * CallsPerSched)
	if total != expected {
		b.Fatalf("incorrect total %v, expected %v", total, expected)
	}
}
