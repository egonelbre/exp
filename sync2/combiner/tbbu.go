package combiner

import (
	"sync/atomic"
	"unsafe"
)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type TBBU struct {
	head    uintptr // *tbbNodeU
	_       pad7
	batcher Batcher
	busy    int64
}

type tbbNodeU struct {
	next     uintptr // *tbbNodeU
	argument Argument
}

func NewTBBU(batcher Batcher) *TBBU {
	return &TBBU{
		batcher: batcher,
		head:    0,
	}
}

func (c *TBBU) Do(arg Argument) {
	node := &tbbNodeU{argument: arg}

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
		for try := 0; atomic.LoadUintptr(&node.next) != 0; spin(&try) {
		}
	} else {
		// 3. We are the combiner.

		// wait for previous combiner to finish
		for try := 0; atomic.LoadInt64(&c.busy) == 1; spin(&try) {
		}
		atomic.StoreInt64(&c.busy, 1)

		// First, execute own operation.
		c.batcher.Start()
		defer c.batcher.Finish()

		// Grab the batch of operations only once
		for {
			cmp = atomic.LoadUintptr(&c.head)
			if atomic.CompareAndSwapUintptr(&c.head, cmp, 0) {
				break
			}
		}

		node = (*tbbNodeU)(unsafe.Pointer(cmp))
		// Execute the list of operations.
		for node != nil {
			next := (*tbbNodeU)(unsafe.Pointer(node.next))
			c.batcher.Include(node.argument)
			// Mark completion.
			atomic.StoreUintptr(&node.next, 0)
			node = next
		}

		// allow next combiner to proceed
		atomic.StoreInt64(&c.busy, 0)
	}
}
