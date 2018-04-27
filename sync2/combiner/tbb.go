package combiner

import (
	"sync/atomic"
	"unsafe"
)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type TBB struct {
	exe  Executor
	busy int64
	head unsafe.Pointer // *tbbArg
}

type tbbArg struct {
	value int64
	next  unsafe.Pointer // *tbbArg
}

func NewTBB(exe Executor) *TBB {
	return &TBB{
		exe:  exe,
		head: nil,
	}
}

var tbbLockedElem = tbbArg{}
var tbbLockedArg = &tbbLockedElem
var tbbLocked = (unsafe.Pointer)(tbbLockedArg)

func (c *TBB) Do(v int64) {
	arg := &tbbArg{value: v}

	var cmp unsafe.Pointer
	for {
		cmp = atomic.LoadPointer(&c.head)
		arg.next = cmp
		if atomic.CompareAndSwapPointer(&c.head, cmp, unsafe.Pointer(arg)) {
			break
		}
	}

	if cmp != nil {
		// 2. If we are not the combiner, wait for arg.next to become nil
		// (which means the operation is finished).
		var s spinner
		for atomic.LoadPointer(&arg.next) != nil && s.spin() {
		}
	} else {
		// 3. We are the combiner.

		// wait for previous combiner to finish
		var s spinner
		for atomic.LoadInt64(&c.busy) == 1 && s.spin() {
		}
		atomic.StoreInt64(&c.busy, 1)

		// First, execute own operation.
		c.exe.Start()
		defer c.exe.Finish()

		// Grab the batch of operations only once
		for {
			cmp = atomic.LoadPointer(&c.head)
			if atomic.CompareAndSwapPointer(&c.head, cmp, nil) {
				break
			}
		}

		arg = (*tbbArg)(cmp)
		// Execute the list of operations.
		for arg != nil {
			next := (*tbbArg)(arg.next)
			c.exe.Include(arg.value)
			// Mark completion.
			atomic.StorePointer(&arg.next, nil)
			arg = next
		}

		// allow next combiner to proceed
		atomic.StoreInt64(&c.busy, 0)
	}
}
