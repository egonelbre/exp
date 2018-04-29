package queue

import (
	"sync/atomic"

	"github.com/egonelbre/exp/sync2/spin"
)

var _ SPSC = (*SeqwSPSCSpinning)(nil)

// SeqwSPSCSpinning is a MPMC queue based on http://www.1024cores.net/home/lock-free-algorithms/queues/bounded-mpmc-queue
type SeqwSPSCSpinning struct {
	_ [8]int64

	buffer []seqwSPSCSpinning
	mask   int64
	_      [4]int64

	sendx int64
	_     [7]int64

	recvx int64
	_     [7]int64
}

type seqwSPSCSpinning struct {
	sequence int64
	value    Value
	_        [6]int64
}

func NewSeqwSPSCSpinning(size int) *SeqwSPSCSpinning {
	if size <= 1 {
		size = 2
	}
	size = int(nextPowerOfTwo(uint32(size)))

	q := &SeqwSPSCSpinning{}
	q.buffer = make([]seqwSPSCSpinning, size)
	q.mask = int64(size) - 1
	for i := range q.buffer {
		q.buffer[i].sequence = int64(i)
	}

	return q
}

func (q *SeqwSPSCSpinning) Cap() int { return len(q.buffer) }

func (q *SeqwSPSCSpinning) Send(v Value) bool {
	var s spin.T256
	for s.Spin() {
		if q.TrySend(v) {
			return true
		}
	}
	return false
}

func (q *SeqwSPSCSpinning) TrySend(v Value) bool {
	var cell *seqwSPSCSpinning
	pos := q.sendx
	for {
		cell = &q.buffer[pos&q.mask]
		seq := atomic.LoadInt64(&cell.sequence)
		df := seq - pos
		if df == 0 {
			q.sendx = pos + 1
			break
		} else if df < 0 {
			// full
			return false
		}
	}

	cell.value = v
	atomic.StoreInt64(&cell.sequence, pos+1)
	return true
}

func (q *SeqwSPSCSpinning) Recv(v *Value) bool {
	var s spin.T256
	for s.Spin() {
		if q.TryRecv(v) {
			return true
		}
	}
	return false
}

func (q *SeqwSPSCSpinning) TryRecv(v *Value) bool {
	var cell *seqwSPSCSpinning
	pos := q.recvx
	for {
		cell = &q.buffer[pos&q.mask]
		seq := atomic.LoadInt64(&cell.sequence)
		df := seq - (pos + 1)
		if df == 0 {
			q.recvx = pos + 1
			break
		} else if df < 0 {
			// empty
			return false
		}
	}

	*v = cell.value
	atomic.StoreInt64(&cell.sequence, pos+q.mask+1)
	return true
}
