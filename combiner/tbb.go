package combiner

import (
	"sync/atomic"
	"unsafe"
)

// based on https://software.intel.com/en-us/blogs/2013/02/22/combineraggregator-synchronization-primitive
type TBB struct {
	head    unsafe.Pointer // *tbbNode
	_       [7]uint64
	batcher Batcher
	busy    int64
}

type tbbNode struct {
	next     unsafe.Pointer // *tbbNode
	argument Argument
}

func NewTBB(batcher Batcher) *TBB {
	return &TBB{
		batcher: batcher,
		head:    nil,
	}
}

func (c *TBB) DoAsync(op Argument) { c.do(op, true) }
func (c *TBB) Do(op Argument)      { c.do(op, false) }

func (c *TBB) do(arg Argument, async bool) {
	node := &tbbNode{argument: arg}

	var cmp unsafe.Pointer
	for {
		cmp = atomic.LoadPointer(&c.head)
		node.next = cmp
		if atomic.CompareAndSwapPointer(&c.head, cmp, unsafe.Pointer(node)) {
			break
		}
	}

	if cmp != nil {
		if async {
			return
		}

		// 2. If we are not the combiner, wait for arg.next to become nil
		// (which means the operation is finished).
		for try := 0; atomic.LoadPointer(&node.next) != nil; spin(&try) {
		}
	} else {
		// 3. We are the combiner.

		// wait for previous combiner to finish
		for try := 0; atomic.LoadInt64(&c.busy) == 1; spin(&try) {
		}
		atomic.StoreInt64(&c.busy, 1)

		// First, execute own operation.
		c.batcher.Start()
		defer c.batcher.Finish()

		// Grab the batch of operations only once
		for {
			cmp = atomic.LoadPointer(&c.head)
			if atomic.CompareAndSwapPointer(&c.head, cmp, nil) {
				break
			}
		}

		node = (*tbbNode)(cmp)
		// Execute the list of operations.
		for node != nil {
			next := (*tbbNode)(node.next)
			c.batcher.Include(node.argument)
			// Mark completion.
			atomic.StorePointer(&node.next, nil)
			node = next
		}

		// allow next combiner to proceed
		atomic.StoreInt64(&c.busy, 0)
	}
}
