package combiner

import (
	"sync/atomic"
	"unsafe"
)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type Basic struct {
	exe  Executor
	head uintptr // *basicArg
}

type basicArg struct {
	value int64
	next  uintptr // *basicArg
}

func NewBasic(exe Executor) *Basic {
	return &Basic{
		exe:  exe,
		head: 0,
	}
}

const basicLocked = uintptr(1)

func (c *Basic) Do(v int64) {
	arg := &basicArg{value: v}

	// c.head can be in 3 states:
	// c.head == 0: no operations in-flight, initial state.
	// c.head == LOCKED: single operation in-flight, no combining opportunities.
	// c.head == pointer to some basicArg that is subject to combining.
	//            The args are chainded into a lock-free list via 'next' fields.
	// arg.next also can be in 3 states:
	// arg.next == pointer to other arg.
	// arg.next == LOCKED: denotes the last arg in the list.
	// arg.next == 0: the operation is finished.

	// The function proceeds in 3 steps:
	// 1. If c.head == nil, exchange it to LOCKED and become the combiner.
	// Otherwise, enqueue own arg into the c->head lock-free list.

	var cmp uintptr
	for {
		cmp = atomic.LoadUintptr(&c.head)
		xchg := basicLocked
		if cmp != 0 {
			// There is already a combiner, enqueue itself.
			xchg = uintptr(unsafe.Pointer(arg))
			arg.next = cmp
		}

		if atomic.CompareAndSwapUintptr(&c.head, cmp, xchg) {
			break
		}
	}

	if cmp != 0 {
		// 2. If we are not the combiner, wait for arg.next to become nil
		// (which means the operation is finished).
		var s spinner
		for atomic.LoadUintptr(&arg.next) != 0 && s.spin() {
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
				cmp = atomic.LoadUintptr(&c.head)
				// If there are some operations in the list,
				// grab the list and replace with LOCKED.
				// Otherwise, exchange to nil.
				var xchg uintptr = 0
				if cmp != basicLocked {
					xchg = basicLocked
				}

				if atomic.CompareAndSwapUintptr(&c.head, cmp, xchg) {
					break
				}
			}

			// No more operations to combine, return.
			if cmp == basicLocked {
				break
			}

			// Execute the list of operations.
			for cmp != basicLocked {
				arg = (*basicArg)(unsafe.Pointer(cmp))
				cmp = arg.next

				c.exe.Include(arg.value)
				// Mark completion.
				atomic.StoreUintptr(&arg.next, 0)
			}
		}
	}
}
