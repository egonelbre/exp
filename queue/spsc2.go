package queue

import (
	"sync"
)

// implementation based on MCRingBuffer
// http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.577.960&rep=rep1&type=pdf

type SPSC2 struct {
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
	mu        sync.Mutex
	reader    sync.Cond
	writer    sync.Cond
	batchSize int64
	buffer    []SPSC2Value
}

type SPSC2Value = int64

func NewSPSC2(size, batchSize int) *SPSC2 {
	q := &SPSC2{}
	q.Init(size, batchSize)
	return q
}

func (q *SPSC2) Init(size, batchSize int) {
	q.reader.L = &q.mu
	q.writer.L = &q.mu
	q.batchSize = int64(batchSize)
	q.buffer = make([]SPSCValue, ceil(size+1, batchSize))
}

func (q *SPSC2) next(i int64) int64 {
	r := i + 1
	if r >= int64(len(q.buffer)) {
		r = 0
	}
	return r
}

func (q *SPSC2) Send(v SPSCValue) bool {
	afterNextWrite := q.next(q.nextWrite)
	if afterNextWrite == q.localRead {
		q.mu.Lock()
		if afterNextWrite == q.read {
			q.writer.Wait()
		}
		q.localRead = q.read
		q.mu.Unlock()
	}

	q.buffer[q.nextWrite] = v
	q.nextWrite = afterNextWrite
	q.writeBatch++
	if q.writeBatch >= q.batchSize {
		q.FlushSend()
	}
	return true
}

func (q *SPSC2) FlushSend() {
	q.mu.Lock()
	q.write = q.nextWrite
	q.writeBatch = 0
	q.mu.Unlock()
	q.reader.Signal()
}

func (q *SPSC2) Recv(v *SPSCValue) bool {
	if q.nextRead == q.localWrite {
		q.mu.Lock()
		if q.nextRead == q.write {
			q.reader.Wait()
		}
		q.localWrite = q.write
		q.mu.Unlock()
	}

	*v = q.buffer[q.nextRead]
	q.buffer[q.nextRead] = 0 // clear value, only needed for pointers

	q.nextRead = q.next(q.nextRead)
	q.readBatch++
	if q.readBatch >= q.batchSize {
		q.mu.Lock()
		q.read = q.nextRead
		q.readBatch = 0
		q.mu.Unlock()
		q.writer.Signal()
	}

	return true
}
