package queue

import (
	"sync/atomic"
	"unsafe"

	"github.com/egonelbre/exp/sync2/spin"
)

var _ MPSC = (*LinkedMPSC)(nil)

// LinkedMPSC is a MPSC queue based on http://www.1024cores.net/home/lock-free-algorithms/queues/non-intrusive-mpsc-node-based-queue
type LinkedMPSC struct {
	_    [8]uint64
	stub linkedMPSC
	_    [7]uint64
	head unsafe.Pointer
	_    [7]uint64
	tail unsafe.Pointer
	_    [7]uint64
}

type linkedMPSC struct {
	next  unsafe.Pointer // *linkedMPSC
	value Value
}

func NewLinkedMPSC() *LinkedMPSC {
	q := &LinkedMPSC{}
	q.head = unsafe.Pointer(&q.stub)
	q.tail = unsafe.Pointer(&q.stub)
	return q
}

func (q *LinkedMPSC) MultipleProducers() {}

func (q *LinkedMPSC) Send(value Value) bool {
	n := &linkedMPSC{value: value}
	prev := exchangePointer(&q.head, unsafe.Pointer(n))
	prevn := (*linkedMPSC)(prev)
	atomic.StorePointer(&prevn.next, unsafe.Pointer(n))
	return true
}

func (q *LinkedMPSC) Recv(value *Value) bool {
	var s spin.T256
	for s.Spin() {
		if q.TryRecv(value) {
			return true
		}
	}
	return false
}

func (q *LinkedMPSC) TryRecv(value *Value) bool {
	tail := (*linkedMPSC)(q.tail)
	next := atomic.LoadPointer(&tail.next)
	if next == nil {
		return false

	}
	q.tail = next
	*value = (*linkedMPSC)(next).value
	return true
}
