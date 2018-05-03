package combiner

import (
	"sync/atomic"
	"unsafe"
)

var _ Combiner = (*Bounded)(nil)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type Bounded struct {
	head    unsafe.Pointer // *boundedNode
	_       pad7
	batcher Batcher
	limit   int
}

type boundedNode struct {
	next     unsafe.Pointer // *boundedNode
	handoff  int64
	argument Argument
}

func NewBounded(batcher Batcher, limit int) *Bounded {
	return &Bounded{
		batcher: batcher,
		limit:   limit,
		head:    nil,
	}
}

var boundedLockedElem = boundedNode{}
var boundedLockedNode = &boundedLockedElem
var boundedLocked = (unsafe.Pointer)(boundedLockedNode)

func (c *Bounded) Do(arg Argument) {
	node := &boundedNode{argument: arg}

	var cmp unsafe.Pointer
	for {
		cmp = atomic.LoadPointer(&c.head)
		xchg := boundedLocked
		if cmp != nil {
			// There is already a combiner, enqueue itself.
			xchg = (unsafe.Pointer)(node)
			node.next = cmp
		}

		if atomic.CompareAndSwapPointer(&c.head, cmp, xchg) {
			break
		}
	}

	handoff := false
	if cmp != nil {
		// 2. If we are not the combiner, wait for arg.next to become nil
		// (which means the operation is finished).
		for try := 0; ; spin(&try) {
			next := atomic.LoadPointer(&node.next)
			if next == nil {
				return
			}

			if atomic.LoadInt64(&node.handoff) == 1 {
				// start combining from the current position
				handoff = true
				break
			}
		}
	}

	// 3. We are the combiner.

	// First, execute own operation.
	c.batcher.Start()
	defer c.batcher.Finish()

	var count int
	if !handoff {
		c.batcher.Include(node.argument)
		count++
	} else {
		// Execute the list of operations.
		for node != boundedLockedNode {
			if count == c.limit {
				atomic.StoreInt64(&node.handoff, 1)
				return
			}
			next := (*boundedNode)(node.next)
			c.batcher.Include(node.argument)
			count++
			// Mark completion.
			atomic.StorePointer(&node.next, nil)
			node = next
		}
	}

	// Then, look for combining opportunities.
	for {
		for {
			cmp = atomic.LoadPointer(&c.head)
			// If there are some operations in the list,
			// grab the list and replace with LOCKED.
			// Otherwise, exchange to nil.
			var xchg unsafe.Pointer = nil
			if cmp != boundedLocked {
				xchg = boundedLocked
			}
			if atomic.CompareAndSwapPointer(&c.head, cmp, xchg) {
				break
			}
		}

		// No more operations to combine, return.
		if cmp == boundedLocked {
			break
		}

		node = (*boundedNode)(cmp)

		// Execute the list of operations.
		for node != boundedLockedNode {
			if count == c.limit {
				atomic.StoreInt64(&node.handoff, 1)
				return
			}
			next := (*boundedNode)(node.next)
			c.batcher.Include(node.argument)
			count++
			// Mark completion.
			atomic.StorePointer(&node.next, nil)
			node = next
		}
	}
}
