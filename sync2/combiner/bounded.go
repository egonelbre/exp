package combiner

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type Bounded struct {
	exe   Executor
	limit int
	head  unsafe.Pointer // *boundedArg
}

type boundedArg struct {
	value int64
	next  unsafe.Pointer // *boundedArg
	busy  int64
}

func NewBounded(exe Executor, limit int) *Bounded {
	return &Bounded{
		exe:   exe,
		limit: limit,
		head:  nil,
	}
}

var boundedLockedElem = boundedArg{}
var boundedLockedArg = &boundedLockedElem
var boundedLocked = (unsafe.Pointer)(boundedLockedArg)

const boundedHandoff = 2

func (c *Bounded) Do(v int64) {
	arg := &boundedArg{value: v}

	// c.head can be in 3 states:
	// c.head == 0: no operations in-flight, initial state.
	// c.head == LOCKED: single operation in-flight, no combining opportunities.
	// c.head == pointer to some boundedArg that is subject to combining.
	//            The args are chainded into a lock-free list via 'next' fields.
	// arg.next also can be in 3 states:
	// arg.next == pointer to other arg.
	// arg.next == LOCKED: denotes the last arg in the list.
	// arg.next == 0: the operation is finished.

	// The function proceeds in 3 steps:
	// 1. If c.head == nil, exchange it to LOCKED and become the combiner.
	// Otherwise, enqueue own arg into the c->head lock-free list.

	var cmp unsafe.Pointer
	for {
		cmp = atomic.LoadPointer(&c.head)
		xchg := boundedLocked
		if cmp != nil {
			// There is already a combiner, enqueue itself.
			xchg = (unsafe.Pointer)(arg)
			arg.next = cmp
		}

		if atomic.CompareAndSwapPointer(&c.head, cmp, xchg) {
			break
		}
	}

	if cmp != nil {
		// 2. If we are not the combiner, wait for arg.next to become nil
		// (which means the operation is finished).

		const maxspin = 256
		spin := 0
		for atomic.LoadPointer(&arg.next) != nil {
			spin++
			if spin >= maxspin {
				runtime.Gosched()
				spin = 0
			}
		}
	} else {
		// 3. We are the combiner.

		// First, execute own operation.
		c.exe.Start()
		defer c.exe.Finish()

		c.exe.Include(arg.value)

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

			arg = (*boundedArg)(cmp)
			// Execute the list of operations.
			for arg != boundedLockedArg {
				next := (*boundedArg)(arg.next)
				c.exe.Include(arg.value)
				// Mark completion.
				atomic.StorePointer(&arg.next, nil)
				arg = next
			}
		}
	}
}
