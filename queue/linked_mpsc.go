package queue

import (
	"sync/atomic"
	"unsafe"

	"github.com/egonelbre/exp/sync2/spin"
)

var linkedStub linkedNode

var _ MPSC = (*LinkedMPSC)(nil)

// LinkedMPSC is a MPSC queue based on http://www.1024cores.net/home/lock-free-algorithms/queues/non-intrusive-mpsc-node-based-queue
type LinkedMPSC struct {
	_    [8]uint64
	stub linkedNode
	_    [7]uint64
	head unsafe.Pointer
	_    [7]uint64
	tail unsafe.Pointer
	_    [7]uint64
}

type linkedNode struct {
	next  unsafe.Pointer // *linkedNode
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
	n := &linkedNode{value: value}
	var prev unsafe.Pointer
	for {
		prev = atomic.LoadPointer(&q.head)
		if atomic.CompareAndSwapPointer(&q.head, prev, unsafe.Pointer(n)) {
			break
		}
	}
	prevn := (*linkedNode)(prev)
	prevn.next = unsafe.Pointer(n)
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
	tail := (*linkedNode)(q.tail)
	next := tail.next
	if next == nil {
		return false

	}
	q.tail = next
	*value = (*linkedNode)(next).value
	return true
}
