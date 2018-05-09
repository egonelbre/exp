package combiner

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type TBBSleepyUintptr struct {
	head uintptr // *tbbNodeSleepyUintptr
	_    [7]uint64
	lock sync.Mutex
	cond sync.Cond
	// cacheline boundary
	batcher Batcher
	busy    int64
}

type tbbNodeSleepyUintptr struct {
	next     uintptr // *tbbNodeSleepyUintptr
	argument Argument
}

func NewTBBSleepyUintptr(batcher Batcher) *TBBSleepyUintptr {
	c := &TBBSleepyUintptr{
		batcher: batcher,
		head:    0,
	}
	c.cond.L = &c.lock
	return c
}

func (c *TBBSleepyUintptr) Do(arg Argument) {
	node := &tbbNodeSleepyUintptr{argument: arg}

	var cmp uintptr
	for {
		cmp = atomic.LoadUintptr(&c.head)
		node.next = cmp
		if atomic.CompareAndSwapUintptr(&c.head, cmp, uintptr(unsafe.Pointer(node))) {
			break
		}
	}

	if cmp != 0 {
		// 2. If we are not the combiner, wait for arg.next to become nil
		// (which means the operation is finished).
		c.lock.Lock()
		for {
			if atomic.LoadUintptr(&node.next) == 0 {
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
			cmp = atomic.LoadUintptr(&c.head)
			if atomic.CompareAndSwapUintptr(&c.head, cmp, 0) {
				break
			}
		}

		node = (*tbbNodeSleepyUintptr)(unsafe.Pointer(cmp))
		// Execute the list of operations.
		for node != nil {
			next := (*tbbNodeSleepyUintptr)(unsafe.Pointer(node.next))
			c.batcher.Include(node.argument)
			// Mark completion.
			atomic.StoreUintptr(&node.next, 0)
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
