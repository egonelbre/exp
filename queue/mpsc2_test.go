package queue_test

import (
	"runtime"
	"strconv"
	"sync"
	"testing"

	"github.com/egonelbre/exp/queue"
)

func BenchmarkMPSC2ProdCons(b *testing.B) {
	for _, chanSize := range []int{1, 10, 100} {
		for _, batchSize := range []int{1, 8, 32, 64} {
			if batchSize > chanSize {
				continue
			}
			name := strconv.Itoa(chanSize) + ":b" + strconv.Itoa(batchSize)
			b.Run(name, func(b *testing.B) {
				benchmarkMPSC2ProdCons(b, batchSize, chanSize, 0)
			})
		}
	}
	for _, chanSize := range []int{1, 10, 100} {
		for _, batchSize := range []int{1, 8, 32, 64} {
			if batchSize > chanSize {
				continue
			}
			name := strconv.Itoa(chanSize) + ":b" + strconv.Itoa(batchSize)
			b.Run("Work"+name, func(b *testing.B) {
				benchmarkMPSC2ProdCons(b, batchSize, chanSize, 100)
			})
		}
	}
}

func benchmarkMPSC2ProdCons(b *testing.B, batchSize, chanSize, localWork int) {
	const CallsPerSched = 1000
	procs := runtime.GOMAXPROCS(-1)
	N := b.N / CallsPerSched

	q := queue.NewMPSC2(batchSize, chanSize)
	for i := 0; i < procs; i++ {
		go func() {
			for i := int64(0); i < int64(N); i++ {
				foo := 0
				for g := 0; g < CallsPerSched; g++ {
					q.Send(i)
					for k := 0; k < localWork; k++ {
						foo *= 2
						foo /= 2
					}
				}
			}
		}()
	}

	total := int64(0)
	for i := int64(0); i < int64(procs*N); i++ {
		foo := 0
		for g := 0; g < CallsPerSched; g++ {
			var v int64
			q.Recv(&v)
			for k := 0; k < localWork/procs; k++ {
				foo *= 2
				foo /= 2
			}
			total += v
		}
	}
	//expected := int64(procs * (N * (N - 1) / 2) * CallsPerSched)
	//if total != expected {
	//	b.Fatalf("incorrect total %v, expected %v", total, expected)
	//}
}
func BenchmarkMPSC2Basic(b *testing.B) {
	const procs = 4
	var q queue.MPSC2
	q.Init(64, 8192)

	var wg sync.WaitGroup
	wg.Add(procs + 1)

	total := b.N
	for i := 0; i < procs; i++ {
		chunk := total / (procs - i)
		go func(n int) {
			runtime.LockOSThread()
			for i := 0; i < n; i++ {
				q.Send(int64(i))
			}
			wg.Done()
		}(chunk)
		total -= chunk
	}

	go func(n int) {
		runtime.LockOSThread()
		for i := 0; i < n; i++ {
			var v int64
			q.Recv(&v)
		}
		wg.Done()
	}(b.N)

	wg.Wait()
}
