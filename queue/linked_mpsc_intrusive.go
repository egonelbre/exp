package queue

import (
	"sync/atomic"
	"unsafe"

	"github.com/egonelbre/exp/sync2/spin"
)

var _ MPSC = (*LinkedMPSCi)(nil)

// LinkedMPSCi is a MPSC queue based on http://www.1024cores.net/home/lock-free-algorithms/queues/intrusive-mpsc-node-based-queue
type LinkedMPSCi struct {
	_    [8]uint64
	stub linkedMPSCi
	_    [7]uint64
	head unsafe.Pointer
	_    [7]uint64
	tail unsafe.Pointer
	_    [7]uint64
}

type linkedMPSCi struct {
	next  unsafe.Pointer // *linkedMPSCi
	value Value
}

func NewLinkedMPSCIntrusive() *LinkedMPSCi {
	q := &LinkedMPSCi{}
	q.head = unsafe.Pointer(&q.stub)
	q.tail = unsafe.Pointer(&q.stub)
	return q
}

func (q *LinkedMPSCi) MultipleProducers() {}

func (q *LinkedMPSCi) Send(value Value) bool {
	q.send(&linkedMPSCi{value: value})
	return true
}

func (q *LinkedMPSCi) send(n *linkedMPSCi) {
	n.next = nil
	prev := exchangePointer(&q.head, unsafe.Pointer(n))
	prevn := (*linkedMPSCi)(prev)
	atomic.StorePointer(&prevn.next, unsafe.Pointer(n))
}

func (q *LinkedMPSCi) Recv(value *Value) bool {
	var s spin.T256
	for s.Spin() {
		if q.TryRecv(value) {
			return true
		}
	}
	return false
}

func (q *LinkedMPSCi) TryRecv(value *Value) bool {
	tail := (*linkedMPSCi)(q.tail)
	next := atomic.LoadPointer(&tail.next)
	if tail == &q.stub {
		if next == nil {
			return false
		}
		q.tail = next
		tail = (*linkedMPSCi)(next)
		next = atomic.LoadPointer(&tail.next)
	}
	if next != nil {
		q.tail = next
		*value = tail.value
		return true
	}

	head := atomic.LoadPointer(&q.head)
	if q.tail != head {
		return false
	}

	q.send(&q.stub)
	next = atomic.LoadPointer(&tail.next)
	if next != nil {
		q.tail = next
		*value = tail.value
		return true
	}

	return false
}
