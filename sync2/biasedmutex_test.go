package sync2

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

func TestBiasedMutexProgress(t *testing.T) {
	for _, rt := range []int{1, 2, 4, 8} {
		for _, wt := range []int{1, 2, 4, 8} {
			name := fmt.Sprintf("r%vw%v", rt, wt)

			t.Run(name, func(t *testing.T) {
				m := NewBiasedMutex()
				m.SetReaderThreshold(rt)
				m.SetWriterThreshold(wt)

				testProgress(t, m)
			})
		}
	}
}

func TestStandardRWMutex(t *testing.T) {
	testProgress(t, &sync.RWMutex{})
}

type RWMutex interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

func testProgress(t *testing.T, m RWMutex) {
	const STEPS = 10000
	const R, W = 100, 100

	var totalProgress int64
	var readerProgress int64
	var writerProgress int64

	var sharedVariable int

	var activeWriters int64
	var activeReaders int64

	var wg sync.WaitGroup

	var mufatal sync.Mutex
	var fatals []string
	fatal := func(err string) {
		mufatal.Lock()
		fatals = append(fatals, err)
		mufatal.Unlock()
		atomic.AddInt64(&totalProgress, STEPS)
	}

	order := make([]byte, STEPS)

	wg.Add(R)
	for i := 0; i < R; i++ {
		go func() {
			defer wg.Done()

			localTotal := 0
			for {
				m.RLock()

				orderx := atomic.AddInt64(&totalProgress, 1)
				if orderx >= STEPS {
					m.RUnlock()
					break
				}
				order[orderx-1] = 'R'

				atomic.AddInt64(&activeReaders, 1)
				if atomic.LoadInt64(&activeWriters) > 0 {
					fatal("writer active during reading")
					atomic.AddInt64(&activeReaders, -1)
					m.RUnlock()
					break
				}
				atomic.AddInt64(&readerProgress, 1)
				localTotal += sharedVariable
				atomic.AddInt64(&activeReaders, -1)
				m.RUnlock()
			}
		}()
	}

	wg.Add(W)
	for i := 0; i < W; i++ {
		go func(i int) {
			defer wg.Done()

			step := 0
			for {
				m.Lock()

				orderx := atomic.AddInt64(&totalProgress, 1)
				if orderx >= STEPS {
					m.Unlock()
					break
				}
				order[orderx-1] = 'W'

				if atomic.AddInt64(&activeWriters, 1) > 1 {
					fatal("more than one writer")
					atomic.AddInt64(&activeWriters, -1)
					m.Unlock()
					break
				}
				if atomic.LoadInt64(&activeReaders) > 0 {
					fatal("reader active during writing")
					atomic.AddInt64(&activeWriters, -1)
					m.Unlock()
					break
				}
				writerProgress++
				sharedVariable += step * i
				step++
				atomic.AddInt64(&activeWriters, -1)
				m.Unlock()
			}
		}(i)
	}

	wg.Wait()

	if len(fatals) > 0 {
		t.Fatal("invalid semantics\n" + strings.Join(fatals, "\n"))
	}

	if readerProgress == 0 || writerProgress == 0 {
		t.Fatalf("no progress reader:%v writer:%v", readerProgress, writerProgress)
	}
	x := STEPS
	if x > 500 {
		x = 500
	}
	t.Logf("reader:%v writer:%v\n%v", readerProgress, writerProgress, string(order[:x]))
}
