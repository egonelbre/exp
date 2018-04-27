package combiner

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func TestFlatCombiner(t *testing.T) {
	t.Skip("broken")

	const N = 100
	var fc FlatCombiner
	var total int64
	var wg sync.WaitGroup

	procs := runtime.GOMAXPROCS(-1)
	wg.Add(procs)

	for proc := 0; proc < procs; proc++ {
		go func() {
			for i := int64(0); i < N; i++ {
				fc.Do(func() {
					atomic.AddInt64(&total, i)
				})
			}
			wg.Done()
		}()
	}

	wg.Wait()
	if total != N*int64(procs) {
		t.Fatalf("got %v expected %v", total, N*procs)
	}
}
