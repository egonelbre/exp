package combiner

import (
	"runtime"
	"strconv"
	"sync"
	"testing"
)

const benchWork = true

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
	runtime.Gosched()
	c.mu.Unlock()
}

func TestBasicCombiner(t *testing.T) {
	const N = 100

	var total Total
	combiner := NewBasic(&total)

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

func BenchmarkBasicCombiner(b *testing.B) {
	for _, procs := range []int{4, 8, 16, 32, 64, 128, 256} {
		b.Run("p"+strconv.Itoa(procs), func(b *testing.B) {
			var wg sync.WaitGroup
			wg.Add(procs)

			var total Total
			combiner := NewBasic(&total)

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

func BenchmarkLockCombiner(b *testing.B) {
	for _, procs := range []int{4, 8, 16, 32, 64, 128, 256} {
		b.Run("p"+strconv.Itoa(procs), func(b *testing.B) {
			var wg sync.WaitGroup
			wg.Add(procs)

			var total Total

			left := b.N
			for i := 0; i < procs; i++ {
				chunk := left / (procs - i)
				go func(n int) {
					runtime.LockOSThread()
					for i := 0; i < n; i++ {
						total.Start()
						total.Include(1)
						total.Finish()
					}
					wg.Done()
				}(chunk)
				left -= chunk
			}

			wg.Wait()
		})
	}
}
