package queue_test

import (
	"runtime"
	"strconv"
	"testing"

	"github.com/egonelbre/exp/queue"
)

func TestSPSCBasic(t *testing.T) {
	const N = 100
	q := queue.NewSPSC_NonBlocking(N*2, 7)
	for i := int64(0); i < N; i++ {
		if !q.Send(i) {
			t.Fatalf("failed to send: %v", i)
		}
	}
	q.SendFlush()

	for i := int64(0); i < N; i++ {
		var w int64
		if !q.Recv(&w) {
			t.Fatalf("failed to recv: %v (got %v)", i, w)
		}
		if w != i {
			t.Fatalf("got %v, expected %v", i, w)
		}
	}
}
func BenchmarkSPSCProdCons(b *testing.B) {
	for _, chanSize := range []int{1, 10, 100} {
		for _, batchSize := range []int{1, 4, 8, 16, 32} {
			name := strconv.Itoa(chanSize) + ":b" + strconv.Itoa(batchSize)
			b.Run(name, func(b *testing.B) {
				benchmarkSPSCProdCons(b, batchSize, chanSize, 0)
			})
		}
	}
	for _, chanSize := range []int{1, 10, 100} {
		for _, batchSize := range []int{1, 4, 8} {
			name := strconv.Itoa(chanSize) + ":b" + strconv.Itoa(batchSize)
			b.Run("Work"+name, func(b *testing.B) {
				benchmarkSPSCProdCons(b, batchSize, chanSize, 0)
			})
		}
	}
}

func benchmarkSPSCProdCons(b *testing.B, batchSize, chanSize, localWork int) {
	const CallsPerSched = 1000
	procs := runtime.GOMAXPROCS(-1)
	N := int64(b.N / CallsPerSched)

	q := queue.NewSPSC(2*procs, batchSize)
	go func() {
		for i := int64(0); i < N; i++ {
			foo := 0
			for g := 0; g < CallsPerSched; g++ {
				for k := 0; k < localWork; k++ {
					foo *= 2
					foo /= 2
				}
				if !q.Send(i) {
					b.Fatalf("failed to send")
				}
			}
		}
		q.SendFlush()
	}()

	total := int64(0)
	for i := int64(0); i < N; i++ {
		foo := 0
		for g := 0; g < CallsPerSched; g++ {
			var v int64
			if !q.Recv(&v) {
				b.Fatalf("failed to extract")
			}
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