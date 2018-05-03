package combiner

import (
	"sync/atomic"
	"unsafe"
)

var _ Combiner = (*BasicU)(nil)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type BasicU struct {
	head    uintptr // *basicUNode
	_       pad7
	batcher Batcher
}

type basicUNode struct {
	next     uintptr // *basicUNode
	argument Argument
}

func NewBasicU(batcher Batcher) *BasicU {
	return &BasicU{
		batcher: batcher,
		head:    0,
	}
}

const basicULocked = uintptr(1)

func (c *BasicU) Do(op Argument) {
	node := &basicUNode{argument: op}

	// c.head can be in 3 states:
	// c.head == 0: no operations in-flight, initial state.
	// c.head == LOCKED: single operation in-flight, no combining opportunities.
	// c.head == pointer to some basicUNode that is subject to combining.
	//            The args are chainded into a lock-free list via 'next' fields.
	// node.next also can be in 3 states:
	// node.next == pointer to other node.
	// node.next == LOCKED: denotes the last node in the list.
	// node.next == 0: the operation is finished.

	// The function proceeds in 3 steps:
	// 1. If c.head == nil, exchange it to LOCKED and become the combiner.
	// Otherwise, enqueue own node into the c->head lock-free list.

	var cmp uintptr
	for {
		cmp = atomic.LoadUintptr(&c.head)
		xchg := basicULocked
		if cmp != 0 {
			// There is already a combiner, enqueue itself.
			xchg = uintptr(unsafe.Pointer(node))
			node.next = cmp
		}

		if atomic.CompareAndSwapUintptr(&c.head, cmp, xchg) {
			break
		}
	}

	if cmp != 0 {
		// 2. If we are not the combiner, wait for node.next to become nil
		// (which means the operation is finished).
		for try := 0; atomic.LoadUintptr(&node.next) != 0; spin(&try) {
		}
	} else {
		// 3. We are the combiner.

		// First, execute own operation.
		c.batcher.Start()
		defer c.batcher.Finish()

		c.batcher.Include(node.argument)

		// Then, look for combining opportunities.
		for {
			for {
				cmp = atomic.LoadUintptr(&c.head)
				// If there are some operations in the list,
				// grab the list and replace with LOCKED.
				// Otherwise, exchange to nil.
				var xchg uintptr = 0
				if cmp != basicULocked {
					xchg = basicULocked
				}

				if atomic.CompareAndSwapUintptr(&c.head, cmp, xchg) {
					break
				}
			}

			// No more operations to combine, return.
			if cmp == basicULocked {
				break
			}

			// Execute the list of operations.
			for cmp != basicULocked {
				node = (*basicUNode)(unsafe.Pointer(cmp))
				cmp = node.next

				c.batcher.Include(node.argument)
				// Mark completion.
				atomic.StoreUintptr(&node.next, 0)
			}
		}
	}
}
