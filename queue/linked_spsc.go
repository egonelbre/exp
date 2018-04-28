package queue

import (
	"sync/atomic"
	"unsafe"

	"github.com/egonelbre/exp/sync2/spin"
)

var _ SPSC = (*LinkedSPSC)(nil)

// LinkedSPSC is a SPSC queue based on http://www.1024cores.net/home/lock-free-algorithms/queues/unbounded-spsc-queue
type LinkedSPSC struct {
	_    [8]uint64
	stub linkedSPSC
	_    [7]uint64
	// producer
	head     unsafe.Pointer
	first    unsafe.Pointer
	tailCopy unsafe.Pointer
	_        [7]uint64
	// consumer
	tail unsafe.Pointer
	_    [7]uint64
}

type linkedSPSC struct {
	next  unsafe.Pointer // *linkedSPSC
	value Value
}

func NewLinkedSPSC() *LinkedSPSC {
	q := &LinkedSPSC{}
	q.head = unsafe.Pointer(&q.stub)
	q.tail = unsafe.Pointer(&q.stub)
	q.first = unsafe.Pointer(&q.stub)
	q.tailCopy = unsafe.Pointer(&q.stub)
	return q
}

func (q *LinkedSPSC) Send(value Value) bool {
	n := q.alloc()
	n.value = value
	n.next = nil
	atomic.StorePointer(&(*linkedSPSC)(q.head).next, unsafe.Pointer(n))
	q.head = unsafe.Pointer(n)
	return true
}

func (q *LinkedSPSC) Recv(value *Value) bool {
	var s spin.T256
	for s.Spin() {
		if q.TryRecv(value) {
			return true
		}
	}
	return false
}

func (q *LinkedSPSC) TryRecv(value *Value) bool {
	tail := (*linkedSPSC)(q.tail)
	next := atomic.LoadPointer(&tail.next)
	if next == nil {
		return false

	}
	q.tail = next
	*value = (*linkedSPSC)(next).value
	return true
}

func (q *LinkedSPSC) alloc() *linkedSPSC {
	// first tries to allocate node from internal node cache,
	// if attempt fails, allocates node via ::operator new()

	if q.first != q.tailCopy {
		n := (*linkedSPSC)(q.first)
		q.first = n.next
		return n
	}

	q.tailCopy = atomic.LoadPointer(&q.tail)
	if q.first != q.tailCopy {
		n := (*linkedSPSC)(q.first)
		q.first = n.next
		return n
	}

	return &linkedSPSC{}
}
