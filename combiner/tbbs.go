package combiner

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type TBBSleepy struct {
	head unsafe.Pointer // *tbbSleepyNode
	_    [7]uint64
	lock sync.Mutex
	cond sync.Cond
	// cacheline boundary
	batcher Batcher
	busy    int64
}

type tbbSleepyNode struct {
	next     unsafe.Pointer // *tbbSleepyNode
	argument Argument
}

func NewTBBSleepy(batcher Batcher) *TBBSleepy {
	c := &TBBSleepy{
		batcher: batcher,
		head:    nil,
	}
	c.cond.L = &c.lock
	return c
}

func (c *TBBSleepy) DoAsync(op Argument) { c.do(op, true) }
func (c *TBBSleepy) Do(op Argument)      { c.do(op, false) }

func (c *TBBSleepy) do(arg Argument, async bool) {
	node := &tbbSleepyNode{argument: arg}

	var cmp unsafe.Pointer
	for {
		cmp = atomic.LoadPointer(&c.head)
		node.next = cmp
		if atomic.CompareAndSwapPointer(&c.head, cmp, unsafe.Pointer(node)) {
			break
		}
	}

	if cmp != nil {
		if async {
			return
		}

		// 2. If we are not the combiner, wait for arg.next to become nil
		// (which means the operation is finished).
		c.lock.Lock()
		for {
			if atomic.LoadPointer(&node.next) == nil {
				c.lock.Unlock()
				return
			}
			c.cond.Wait()
		}
		c.lock.Unlock()
	} else {
		// 3. We are the combiner.

		// wait for previous combiner to finish
		c.lock.Lock()
		for {
			if atomic.LoadInt64(&c.busy) != 1 {
				break
			}
			c.cond.Wait()
		}
		atomic.StoreInt64(&c.busy, 1)
		c.lock.Unlock()

		// First, execute own operation.
		c.batcher.Start()

		// Grab the batch of operations only once
		for {
			cmp = atomic.LoadPointer(&c.head)
			if atomic.CompareAndSwapPointer(&c.head, cmp, nil) {
				break
			}
		}

		node = (*tbbSleepyNode)(cmp)
		// Execute the list of operations.
		for node != nil {
			next := (*tbbSleepyNode)(node.next)
			c.batcher.Include(node.argument)
			// Mark completion.
			atomic.StorePointer(&node.next, nil)
			node = next
		}

		c.batcher.Finish()

		// allow next combiner to proceed
		c.lock.Lock()
		atomic.StoreInt64(&c.busy, 0)
		c.cond.Broadcast()
		c.lock.Unlock()
	}
}
