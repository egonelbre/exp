package combiner

import (
	"sync"

	"github.com/egonelbre/exp/sync2"
)

type Mutex struct {
	mu      sync.Mutex
	batcher Batcher
}

func NewMutex(batcher Batcher) *Mutex {
	c := &Mutex{}
	c.batcher = batcher
	return c
}

func (c *Mutex) Do(v Argument) {
	c.mu.Lock()
	c.batcher.Start()
	c.batcher.Include(v)
	c.batcher.Finish()
	c.mu.Unlock()
}

type SpinMutex struct {
	mu      sync2.SpinMutex
	batcher Batcher
}

func NewSpinMutex(batcher Batcher) *SpinMutex {
	c := &SpinMutex{}
	c.batcher = batcher
	return c
}

func (c *SpinMutex) Do(v Argument) {
	c.mu.Lock()
	c.batcher.Start()
	c.batcher.Include(v)
	c.batcher.Finish()
	c.mu.Unlock()
}
