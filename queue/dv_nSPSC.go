package queue

import (
	"sync/atomic"
	"unsafe"

	"github.com/egonelbre/exp/sync2/spin"
)

var _ SPSC = (*SPSCns_dv)(nil)

// SPSCns_dv is a SPSC queue based on http://www.1024cores.net/home/lock-free-algorithms/queues/unbounded-spsc-queue
type SPSCns_dv struct {
	_    [8]uint64
	stub Node
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

func NewSPSCns_dv() *SPSCns_dv {
	q := &SPSCns_dv{}
	q.head = unsafe.Pointer(&q.stub)
	q.tail = unsafe.Pointer(&q.stub)
	q.first = unsafe.Pointer(&q.stub)
	q.tailCopy = unsafe.Pointer(&q.stub)
	return q
}

func (q *SPSCns_dv) Send(value Value) bool {
	n := q.alloc()
	n.Value = value
	n.next = nil
	atomic.StorePointer(&(*Node)(q.head).next, unsafe.Pointer(n))
	q.head = unsafe.Pointer(n)
	return true
}

func (q *SPSCns_dv) TrySend(value Value) bool { return q.Send(value) }

func (q *SPSCns_dv) Recv(value *Value) bool {
	var s spin.T256
	for s.Spin() {
		if q.TryRecv(value) {
			return true
		}
	}
	return false
}

func (q *SPSCns_dv) TryRecv(value *Value) bool {
	tail := (*Node)(q.tail)
	next := atomic.LoadPointer(&tail.next)
	if next == nil {
		return false
	}
	q.tail = next
	*value = (*Node)(next).Value
	return true
}

func (q *SPSCns_dv) alloc() *Node {
	// first tries to allocate node from internal node cache,
	// if attempt fails, allocates node via ::operator new()

	if q.first != q.tailCopy {
		n := (*Node)(q.first)
		q.first = n.next
		return n
	}

	q.tailCopy = atomic.LoadPointer(&q.tail)
	if q.first != q.tailCopy {
		n := (*Node)(q.first)
		q.first = n.next
		return n
	}

	return &Node{}
}
