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
	value   int64
	next    unsafe.Pointer // *boundedArg
	handoff int64
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

func (c *Bounded) Do(v int64) {
	arg := &boundedArg{value: v}

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

	handoff := false
	if cmp != nil {
		// 2. If we are not the combiner, wait for arg.next to become nil
		// (which means the operation is finished).
		const maxspin = 256
		spin := 0
		for {
			next := atomic.LoadPointer(&arg.next)
			if next == nil {
				return
			}

			if atomic.LoadInt64(&arg.handoff) == 1 {
				// start combining from the current position
				handoff = true
				break
			}

			spin++
			if spin >= maxspin {
				runtime.Gosched()
				spin = 0
			}
		}
	}

	// 3. We are the combiner.

	// First, execute own operation.
	c.exe.Start()
	defer c.exe.Finish()

	var count int
	if !handoff {
		c.exe.Include(arg.value)
		count++
	} else {
		// Execute the list of operations.
		for arg != boundedLockedArg {
			if count == c.limit {
				atomic.StoreInt64(&arg.handoff, 1)
				return
			}
			next := (*boundedArg)(arg.next)
			c.exe.Include(arg.value)
			count++
			// Mark completion.
			atomic.StorePointer(&arg.next, nil)
			arg = next
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

		arg = (*boundedArg)(cmp)

		// Execute the list of operations.
		for arg != boundedLockedArg {
			if count == c.limit {
				atomic.StoreInt64(&arg.handoff, 1)
				return
			}
			next := (*boundedArg)(arg.next)
			c.exe.Include(arg.value)
			count++
			// Mark completion.
			atomic.StorePointer(&arg.next, nil)
			arg = next
		}
	}
}
