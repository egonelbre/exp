package queue

import (
	"runtime"
	"sync/atomic"
)

// implementation based on MCRingBuffer
// http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.577.960&rep=rep1&type=pdf

type SPSC struct {
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
	buffer    []SPSCValue
}

type SPSCValue = int64

func NewSPSC(size, batchSize int) *SPSC {
	q := &SPSC{}
	q.Init(size, batchSize)
	return q
}

func (q *SPSC) Init(size, batchSize int) {
	q.blocking = true
	q.batchSize = int64(batchSize)
	q.buffer = make([]SPSCValue, ceil(size, batchSize))
}

func (q *SPSC) SetBlocking(blocking bool) {
	q.blocking = blocking
}

func (q *SPSC) yield() {
	runtime.Gosched()
}

func (q *SPSC) next(i int64) int64 {
	r := i + 1
	if r >= int64(len(q.buffer)) {
		r = 0
	}
	return r
}

func (q *SPSC) Send(v SPSCValue) bool {
	afterNextWrite := q.next(q.nextWrite)
	if afterNextWrite == q.localRead {
		for afterNextWrite == atomic.LoadInt64(&q.read) {
			if !q.blocking {
				return false
			}
			q.yield()
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

func (q *SPSC) FlushSend() {
	atomic.StoreInt64(&q.write, q.nextWrite)
	q.writeBatch = 0
}

func (q *SPSC) Recv(v *SPSCValue) bool {
	if q.nextRead == q.localWrite {
		for q.nextRead == atomic.LoadInt64(&q.write) {
			if !q.blocking {
				return false
			}
			q.yield()
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
