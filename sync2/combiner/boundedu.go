package combiner

import (
	"sync/atomic"
	"unsafe"
)

var _ Combiner = (*BoundedU)(nil)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type BoundedU struct {
	head    uintptr // *boundedUNode
	_       pad7
	batcher Batcher
	limit   int
}

type boundedUNode struct {
	next     uintptr // *boundedUNode
	argument Argument
}

func NewBoundedU(batcher Batcher, limit int) *BoundedU {
	return &BoundedU{
		batcher: batcher,
		limit:   limit,
		head:    0,
	}
}

const (
	boundeduLocked     = uintptr(1)
	boundeduHandoffTag = uintptr(2)
)

func (c *BoundedU) Do(arg Argument) {
	node := &boundedUNode{argument: arg}

	var cmp uintptr
	for {
		cmp = atomic.LoadUintptr(&c.head)
		xchg := boundeduLocked
		if cmp != 0 {
			// There is already a combiner, enqueue itself.
			xchg = uintptr(unsafe.Pointer(node))
			node.next = cmp
		}

		if atomic.CompareAndSwapUintptr(&c.head, cmp, xchg) {
			break
		}
	}

	count := 0
	handoff := false
	if cmp != 0 {
		// 2. If we are not the combiner, wait for arg.next to become nil
		// (which means the operation is finished).
		for try := 0; ; spin(&try) {
			next := atomic.LoadUintptr(&node.next)
			if next == 0 {
				return
			}

			if next&boundeduHandoffTag != 0 {
				node.next &^= boundeduHandoffTag
				// DO COMBINING
				handoff = true
				break
			}
		}
	}

	// 3. We are the combiner.

	// First, execute own operation.
	c.batcher.Start()
	defer c.batcher.Finish()
	c.batcher.Include(node.argument)
	count++

	// Then, look for combining opportunities.
	for {
		if handoff { // using goto, to keep it similar to D. Vyukov-s implementation
			handoff = false
			goto combiner
		}

		for {
			cmp = atomic.LoadUintptr(&c.head)
			// If there are some operations in the list,
			// grab the list and replace with LOCKED.
			// Otherwise, exchange to nil.
			var xchg uintptr = 0
			if cmp != boundeduLocked {
				xchg = boundeduLocked
			}

			if atomic.CompareAndSwapUintptr(&c.head, cmp, xchg) {
				break
			}
		}

		// No more operations to combine, return.
		if cmp == boundeduLocked {
			break
		}

	combiner:
		// Execute the list of operations.
		for cmp != boundeduLocked {
			node = (*boundedUNode)(unsafe.Pointer(cmp))
			if count == c.limit {
				atomic.StoreUintptr(&node.next, node.next|boundeduHandoffTag)
				return
			}
			cmp = node.next

			c.batcher.Include(node.argument)
			count++
			// Mark completion.
			atomic.StoreUintptr(&node.next, 0)
		}
	}
}
