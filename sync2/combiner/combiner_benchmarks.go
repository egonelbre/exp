package combiner

import (
	"sync"
	"testing"
)

func RunBenchmarks(b *testing.B, procs int, create func(Batcher) Combiner) {
	b.Helper()
	b.Run("Sum/s0i0f0", func(b *testing.B) {
		var worker Worker
		benchSum(b, &worker, procs, create)
	})

	b.Run("Sum/s0i0f100", func(b *testing.B) {
		var worker Worker
		worker.WorkFinish = 100
		benchSum(b, &worker, procs, create)
	})
}

func benchSum(b *testing.B, worker *Worker, procs int, create func(Batcher) Combiner) {
	const N = 100
	combiner := create(worker)
	defer StartClose(combiner)()

	var wg sync.WaitGroup
	wg.Add(procs)

	left := b.N
	for i := 0; i < procs; i++ {
		chunk := left / (procs - i)
		go func(n int) {
			for i := 0; i < n; i++ {
				combiner.Do(int64(1))
			}
			wg.Done()
		}(chunk)
		left -= chunk
	}

	wg.Wait()
}
