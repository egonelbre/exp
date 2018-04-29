package sync2

import (
	"sync"
	"sync/atomic"

	"github.com/egonelbre/exp/sync2/runtime2"
)

type Mutex struct {
	mu    sync.Mutex
	owner int64
}

func (mu *Mutex) Lock() {
	mu.mu.Lock()
	atomic.StoreInt64(&mu.owner, runtime2.GOID())
}

func (mu *Mutex) Unlock() {
	owner := mu.owner
	mu.mu.Unlock()
	atomic.CompareAndSwapInt64(&mu.owner, owner, 0)
}

func (mu *Mutex) Own() bool {
	return atomic.LoadInt64(&mu.owner) == runtime2.GOID()
}

func (mu *Mutex) MustOwn() {
	owner := atomic.LoadInt64(&mu.owner)
	if owner != runtime2.GOID() {
		if owner == 0 {
			panic("mutex was not locked")
		} else {
			panic("owner was different")
		}
	}
}
