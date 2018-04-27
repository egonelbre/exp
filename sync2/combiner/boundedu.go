package combiner

import (
	"sync/atomic"
	"unsafe"
)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type BoundedU struct {
	exe   Executor
	limit int
	head  uintptr // *boundedUArg
}

type boundedUArg struct {
	value int64
	next  uintptr // *boundedUArg
}

func NewBoundedU(exe Executor, limit int) *BoundedU {
	return &BoundedU{
		exe:   exe,
		limit: limit,
		head:  0,
	}
}

const (
	boundeduLocked     = uintptr(1)
	boundeduHandoffTag = uintptr(2)
)

func (c *BoundedU) Do(v int64) {
	arg := &boundedUArg{value: v}

	var cmp uintptr
	for {
		cmp = atomic.LoadUintptr(&c.head)
		xchg := boundeduLocked
		if cmp != 0 {
			// There is already a combiner, enqueue itself.
			xchg = uintptr(unsafe.Pointer(arg))
			arg.next = cmp
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
		var s spinner
		for s.spin() {
			next := atomic.LoadUintptr(&arg.next)
			if next == 0 {
				return
			}

			if next&boundeduHandoffTag != 0 {
				arg.next &^= boundeduHandoffTag
				// DO COMBINING
				handoff = true
				break
			}
		}
		if !handoff {
			return
		}
	}

	// 3. We are the combiner.

	// First, execute own operation.
	c.exe.Start()
	defer c.exe.Finish()

	if !handoff {
		c.exe.Include(arg.value)
		count++
	}

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
			arg = (*boundedUArg)(unsafe.Pointer(cmp))
			if count == c.limit {
				atomic.StoreUintptr(&arg.next, arg.next|boundeduHandoffTag)
				return
			}
			cmp = arg.next

			c.exe.Include(arg.value)
			count++
			// Mark completion.
			atomic.StoreUintptr(&arg.next, 0)
		}
	}
}
