package combiner

import (
	"runtime"
	"sync/atomic"
	"unsafe"

	"github.com/egonelbre/exp/sync2/runtime2"
)

const (
	max_cores = 64
	max_items = 256
	max_spin  = 256
)

type Flat struct {
	proc [maxProcessors]flatProc
	tail unsafe.Pointer
}

type flatProc struct {
	root flatNode
	_    [7]int64
}

type flatNode struct {
	request  func()
	wait     int64
	complete int64
	next     unsafe.Pointer
}

func (q *Flat) Do(request func()) {
	pid := runtime2.ProcessorHint()
	nextNode := &q.proc[pid%maxProcessors].root
	nextNode.complete = 0
	atomic.StoreInt64(&nextNode.wait, 1)

	var curNode *flatNode
	for {
		curNode = (*flatNode)(atomic.LoadPointer(&q.tail))
		if atomic.CompareAndSwapPointer(&q.tail, (unsafe.Pointer)(curNode), (unsafe.Pointer)(nextNode)) {
			break
		}
	}

	curNode.request = request
	atomic.StorePointer((*unsafe.Pointer)(&curNode.next), (unsafe.Pointer)(nextNode))

	spin := 0
	for atomic.LoadInt64(&curNode.wait) == 1 {
		if spin++; spin >= max_spin {
			runtime.Gosched()
			spin = 0
		}
	}

	if curNode.complete == 1 {
		return
	}

	n := curNode
	for combined := 0; combined < max_items; combined++ {
		n.request()
		nn := (*flatNode)(atomic.LoadPointer(&n.next))
		atomic.StorePointer(&n.next, nil)
		n.request = nil
		n.complete = 1
		atomic.StoreInt64(&n.wait, 0)

		n = nn
		if n == nil {
			break
		}
	}
}
