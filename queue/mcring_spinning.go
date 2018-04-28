package queue

import (
	"sync/atomic"

	"github.com/egonelbre/exp/sync2/spin"
)

// implementation based on MCRingBuffer
// http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.577.960&rep=rep1&type=pdf

var _ BlockingSPSC = (*MCRingSpinning)(nil)

// MCRingSpinning is a SPSC queue
type MCRingSpinning struct {
	_ [8]uint64
	// volatile
	read  int64
	write int64
	_     [8 - 2]uint64
	// consumer
	localWrite int64
	nextRead   int64
	readBatch  int64
	_          [8 - 3]uint64
	// producer
	localRead  int64
	nextWrite  int64
	writeBatch int64
	_          [8 - 3]uint64
	// constant
	blocking  bool
	batchSize int64
	buffer    []Value
}

func NewMCRingSpinning(batchSize, size int) *MCRingSpinning {
	q := &MCRingSpinning{}
	q.Init(batchSize, size)
	return q
}

func (q *MCRingSpinning) Cap() int { return len(q.buffer) - 1 }

func (q *MCRingSpinning) Init(batchSize, size int) {
	q.blocking = true
	q.batchSize = int64(batchSize)
	q.buffer = make([]Value, ceil(size+1, batchSize))
}

func (q *MCRingSpinning) SetBlocking(blocking bool) {
	q.blocking = blocking
}

func (q *MCRingSpinning) next(i int64) int64 {
	r := i + 1
	if r >= int64(len(q.buffer)) {
		r = 0
	}
	return r
}

func (q *MCRingSpinning) Send(v Value) bool {
	afterNextWrite := q.next(q.nextWrite)
	if afterNextWrite == q.localRead {
		var s spin.T
		for afterNextWrite == atomic.LoadInt64(&q.read) && s.Spin() {
			if !q.blocking {
				return false
			}
		}
		q.localRead = atomic.LoadInt64(&q.read)
	}

	q.buffer[q.nextWrite] = v
	q.nextWrite = afterNextWrite
	q.writeBatch++
	if q.writeBatch >= q.batchSize {
		q.FlushSend()
	}
	return true
}

func (q *MCRingSpinning) FlushSend() {
	atomic.StoreInt64(&q.write, q.nextWrite)
	q.writeBatch = 0
}

func (q *MCRingSpinning) Recv(v *Value) bool {
	if q.nextRead == q.localWrite {
		var s spin.T
		for q.nextRead == atomic.LoadInt64(&q.write) && s.Spin() {
			if !q.blocking {
				return false
			}
		}
		q.localWrite = atomic.LoadInt64(&q.write)
	}

	*v = q.buffer[q.nextRead]
	q.buffer[q.nextRead] = 0 // clear value, only needed for pointers

	q.nextRead = q.next(q.nextRead)
	q.readBatch++
	if q.readBatch >= q.batchSize {
		atomic.StoreInt64(&q.read, q.nextRead)
		q.readBatch = 0
	}

	return true
}
