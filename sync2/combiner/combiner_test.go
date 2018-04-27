package combiner

import (
	"runtime"
	"strconv"
	"sync"
	"testing"

	"github.com/loov/hrtime"
)

const benchWork = false

type Total struct {
	mu    sync.Mutex
	total int64
}

func (c *Total) Start() {
	c.mu.Lock()
}

func (c *Total) Include(v int64) {
	c.total += v

	if benchWork {
		foo := 1
		for i := 0; i < 100; i++ {
			foo *= 2
			foo /= 2
		}
	}
}

func (c *Total) Finish() {
	c.mu.Unlock()
}

func test(t *testing.T, create func(Executor) Combiner) {
	t.Run("Summing", func(t *testing.T) {
		testSumming(t, create)
	})
	t.Run("Latencies", func(t *testing.T) {
		testLatencies(t, create)
	})
}

func bench(b *testing.B, create func(Executor) Combiner) {
	b.Run("Summing", func(b *testing.B) {
		benchSumming(b, create)
	})
}

func testSumming(t *testing.T, create func(Executor) Combiner) {
	const N = 100

	var total Total
	combiner := create(&total)

	var wg sync.WaitGroup

	procs := runtime.GOMAXPROCS(-1) * 8
	wg.Add(procs)

	for proc := 0; proc < procs; proc++ {
		go func() {
			for i := int64(0); i < N; i++ {
				combiner.Do(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	if total.total != N*int64(procs) {
		t.Fatalf("got %v expected %v", total.total, N*procs)
	}
}

func testLatencies(t *testing.T, create func(Executor) Combiner) {
	const N = 300
	const K = 30
	for _, procs := range []int{4, 8, 16, 32, 64, 128, 256} {
		t.Run("p"+strconv.Itoa(procs), func(t *testing.T) {
			hrs := make([]*hrtime.Benchmark, procs)
			for i := range hrs {
				hrs[i] = hrtime.NewBenchmark(N)
			}

			var wg sync.WaitGroup
			wg.Add(procs)

			var total Total
			combiner := create(&total)

			for i := 0; i < procs; i++ {
				go func(b *hrtime.Benchmark) {
					runtime.LockOSThread()
					for b.Next() {
						for k := 0; k < K; k++ {
							combiner.Do(1)
						}
					}
					wg.Done()
				}(hrs[i])
			}
			wg.Wait()

			hist := hrtime.CombinedHistogram(10, hrs...)
			hist.Divide(K)
			t.Log("\n" + hist.String())
		})
	}
}

func benchSumming(b *testing.B, create func(Executor) Combiner) {
	for _, procs := range []int{4, 8, 16, 32, 64, 128, 256} {
		b.Run("p"+strconv.Itoa(procs), func(b *testing.B) {
			var wg sync.WaitGroup
			wg.Add(procs)

			var total Total
			combiner := create(&total)

			left := b.N
			for i := 0; i < procs; i++ {
				chunk := left / (procs - i)
				go func(n int) {
					runtime.LockOSThread()
					for i := 0; i < n; i++ {
						combiner.Do(1)
					}
					wg.Done()
				}(chunk)
				left -= chunk
			}

			wg.Wait()
		})
	}
}
