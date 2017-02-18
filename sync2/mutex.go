package sync2

import (
	"sync"
	"sync/atomic"
)

type Mutex struct {
	mu    sync.Mutex
	owner int64
}

func (mu *Mutex) Lock() {
	mu.mu.Lock()
	atomic.StoreInt64(&mu.owner, goid())
}

func (mu *Mutex) Unlock() {
	owner := mu.owner
	mu.mu.Unlock()
	atomic.CompareAndSwapInt64(&mu.owner, owner, 0)
}

func (mu *Mutex) Own() bool {
	return atomic.LoadInt64(&mu.owner) == goid()
}

func (mu *Mutex) MustOwn() {
	owner := atomic.LoadInt64(&mu.owner)
	if owner != goid() {
		if owner == 0 {
			panic("mutex was not locked")
		} else {
			panic("owner was different")
		}
	}
}
