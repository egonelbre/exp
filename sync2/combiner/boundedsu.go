package combiner

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

var _ Combiner = (*BoundedSleepyUintptr)(nil)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type BoundedSleepyUintptr struct {
	head    uintptr // *boundedSleepyUintptrNode
	_       pad7
	lock    sync.Mutex
	cond    sync.Cond
	_       pad7
	batcher Batcher
	limit   int
}

type boundedSleepyUintptrNode struct {
	next     uintptr // *boundedSleepyUintptrNode
	argument Argument
}

func NewBoundedSleepyUintptr(batcher Batcher, limit int) *BoundedSleepyUintptr {
	c := &BoundedSleepyUintptr{
		batcher: batcher,
		limit:   limit,
		head:    0,
	}
	c.cond.L = &c.lock
	return c
}

const (
	boundedSleepyUintptrLocked     = uintptr(1)
	boundedSleepyUintptrHandoffTag = uintptr(2)
)

func (c *BoundedSleepyUintptr) Do(arg Argument) {
	node := &boundedSleepyUintptrNode{argument: arg}

	var cmp uintptr
	for {
		cmp = atomic.LoadUintptr(&c.head)
		xchg := boundedSleepyUintptrLocked
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
		c.lock.Lock()
		for {
			next := atomic.LoadUintptr(&node.next)
			if next == 0 {
				c.lock.Unlock()
				return
			}

			if next&boundedSleepyUintptrHandoffTag != 0 {
				node.next &^= boundedSleepyUintptrHandoffTag
				// DO COMBINING
				handoff = true
				break
			}
			c.cond.Wait()
		}
		c.lock.Unlock()
	}

	// 3. We are the combiner.

	// First, execute own operation.
	c.batcher.Start()
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
			if cmp != boundedSleepyUintptrLocked {
				xchg = boundedSleepyUintptrLocked
			}

			if atomic.CompareAndSwapUintptr(&c.head, cmp, xchg) {
				break
			}
		}

		// No more operations to combine, return.
		if cmp == boundedSleepyUintptrLocked {
			break
		}

	combiner:
		// Execute the list of operations.
		for cmp != boundedSleepyUintptrLocked {
			node = (*boundedSleepyUintptrNode)(unsafe.Pointer(cmp))
			if count == c.limit {
				atomic.StoreUintptr(&node.next, node.next|boundedSleepyUintptrHandoffTag)
				c.batcher.Finish()

				c.lock.Lock()
				c.cond.Broadcast()
				c.lock.Unlock()
				return
			}
			cmp = node.next

			c.batcher.Include(node.argument)
			count++
			// Mark completion.
			atomic.StoreUintptr(&node.next, 0)
		}
	}

	c.batcher.Finish()

	c.lock.Lock()
	c.cond.Broadcast()
	c.lock.Unlock()
}
