package queue

import (
	"sync/atomic"
	"unsafe"

	"github.com/egonelbre/exp/sync2/spin"
)

var _ MPSC = (*MPSCnsiDV)(nil)

// MPSCnsiDV is a MPSC queue based on http://www.1024cores.net/home/lock-free-algorithms/queues/intrusive-mpsc-node-based-queue
// TODO: create intrusive API
type MPSCnsiDV struct {
	_    [8]uint64
	stub Node
	_    [7]uint64
	head unsafe.Pointer
	_    [7]uint64
	tail unsafe.Pointer
	_    [7]uint64
}

func NewMPSCnsiDV() *MPSCnsiDV {
	q := &MPSCnsiDV{}
	q.head = unsafe.Pointer(&q.stub)
	q.tail = unsafe.Pointer(&q.stub)
	return q
}

func (q *MPSCnsiDV) MultipleProducers() {}

func (q *MPSCnsiDV) Send(value Value) bool {
	q.send(&Node{Value: value})
	return true
}

func (q *MPSCnsiDV) TrySend(value Value) bool { return q.Send(value) }

func (q *MPSCnsiDV) send(n *Node) {
	n.next = nil
	prev := atomic.SwapPointer(&q.head, unsafe.Pointer(n))
	prevn := (*Node)(prev)
	atomic.StorePointer(&prevn.next, unsafe.Pointer(n))
}

func (q *MPSCnsiDV) Recv(value *Value) bool {
	var s spin.T256
	for s.Spin() {
		if q.TryRecv(value) {
			return true
		}
	}
	return false
}

func (q *MPSCnsiDV) TryRecv(value *Value) bool {
	tail := (*Node)(q.tail)
	next := atomic.LoadPointer(&tail.next)
	if tail == &q.stub {
		if next == nil {
			return false
		}
		q.tail = next
		tail = (*Node)(next)
		next = atomic.LoadPointer(&tail.next)
	}
	if next != nil {
		q.tail = next
		*value = tail.Value
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
		*value = tail.Value
		return true
	}

	return false
}
