package queue

import (
	"sync/atomic"
	"unsafe"

	"github.com/egonelbre/exp/sync2/spin"
)

var _ MPSC = (*MPSCns_dv)(nil)

// MPSCns_dv is a MPSC queue based on http://www.1024cores.net/home/lock-free-algorithms/queues/non-intrusive-mpsc-node-based-queue
type MPSCns_dv struct {
	_    [8]uint64
	stub Node
	_    [7]uint64
	head unsafe.Pointer
	_    [7]uint64
	tail unsafe.Pointer
	_    [7]uint64
}

func NewMPSCns_dv() *MPSCns_dv {
	q := &MPSCns_dv{}
	q.head = unsafe.Pointer(&q.stub)
	q.tail = unsafe.Pointer(&q.stub)
	return q
}

func (q *MPSCns_dv) MultipleProducers() {}

func (q *MPSCns_dv) Send(value Value) bool {
	n := &Node{Value: value}
	prev := atomic.SwapPointer(&q.head, unsafe.Pointer(n))
	prevn := (*Node)(prev)
	atomic.StorePointer(&prevn.next, unsafe.Pointer(n))
	return true
}

func (q *MPSCns_dv) TrySend(value Value) bool { return q.Send(value) }

func (q *MPSCns_dv) Recv(value *Value) bool {
	var s spin.T256
	for s.Spin() {
		if q.TryRecv(value) {
			return true
		}
	}
	return false
}

func (q *MPSCns_dv) TryRecv(value *Value) bool {
	tail := (*Node)(q.tail)
	next := atomic.LoadPointer(&tail.next)
	if next == nil {
		return false

	}
	q.tail = next
	*value = (*Node)(next).Value
	return true
}
