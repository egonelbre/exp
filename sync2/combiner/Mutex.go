package combiner

import "sync"

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
