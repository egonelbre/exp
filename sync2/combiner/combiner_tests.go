package combiner

import (
	"sync"
	"testing"
)

func RunTests(t *testing.T, procs int, create func(Batcher) Combiner) {
	t.Run("Sum", func(t *testing.T) {
		testSum(t, procs, create)
	})
	t.Run("SumSequence", func(t *testing.T) {
		testSumSequence(t, procs, create)
	})
}

func testSum(t *testing.T, procs int, create func(Batcher) Combiner) {
	const N = 100

	var worker Worker
	combiner := create(&worker)

	var wg sync.WaitGroup

	wg.Add(procs)
	for proc := 0; proc < procs; proc++ {
		go func() {
			for i := int64(0); i < N; i++ {
				combiner.Do(int64(1))
			}
			wg.Done()
		}()
	}

	wg.Wait()
	if worker.Total != N*int64(procs) {
		t.Fatalf("got %v expected %v", worker.Total, N*procs)
	}
}

func testSumSequence(t *testing.T, procs int, create func(Batcher) Combiner) {
	const N = 100

	var worker Worker
	combiner := create(&worker)

	var wg sync.WaitGroup

	wg.Add(procs)

	for proc := 0; proc < procs; proc++ {
		go func() {
			for i := int64(0); i < N; i++ {
				combiner.Do(i)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	if worker.Total != int64(procs)*N*(N-1)/2 {
		t.Fatalf("got %v expected %v", worker.Total, N*procs)
	}
}
