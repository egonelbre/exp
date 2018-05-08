package combiner

import (
	"sync/atomic"
	"unsafe"
)

var _ Combiner = (*BoundedUintptr)(nil)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type BoundedUintptr struct {
	head    uintptr // *boundedUintptrNode
	_       pad7
	batcher Batcher
	limit   int
}

type boundedUintptrNode struct {
	next     uintptr // *boundedUintptrNode
	argument Argument
}

func NewBoundedUintptr(batcher Batcher, limit int) *BoundedUintptr {
	return &BoundedUintptr{
		batcher: batcher,
		limit:   limit,
		head:    0,
	}
}

const (
	boundedUintptrLocked     = uintptr(1)
	boundedUintptrHandoffTag = uintptr(2)
)

func (c *BoundedUintptr) Do(arg Argument) {
	node := &boundedUintptrNode{argument: arg}

	var cmp uintptr
	for {
		cmp = atomic.LoadUintptr(&c.head)
		xchg := boundedUintptrLocked
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

			if next&boundedUintptrHandoffTag != 0 {
				node.next &^= boundedUintptrHandoffTag
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
			if cmp != boundedUintptrLocked {
				xchg = boundedUintptrLocked
			}

			if atomic.CompareAndSwapUintptr(&c.head, cmp, xchg) {
				break
			}
		}

		// No more operations to combine, return.
		if cmp == boundedUintptrLocked {
			break
		}

	combiner:
		// Execute the list of operations.
		for cmp != boundedUintptrLocked {
			node = (*boundedUintptrNode)(unsafe.Pointer(cmp))
			if count == c.limit {
				atomic.StoreUintptr(&node.next, node.next|boundedUintptrHandoffTag)
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
