package queue

import (
	"sync"
)

// implementation based on MCRingBuffer
// http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.577.960&rep=rep1&type=pdf

var _ SPSC = (*SPSCrMC)(nil)

// SPSCrMC is a SPSC queue based on MCRingBuffer
type SPSCrMC struct {
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
	buffer    []Value
}

func NewSPSCrMC(batchSize, size int) *SPSCrMC {
	q := &SPSCrMC{}
	q.Init(size, batchSize)
	return q
}

func (q *SPSCrMC) Cap() int { return len(q.buffer) - 1 }

func (q *SPSCrMC) Init(size, batchSize int) {
	q.reader.L = &q.mu
	q.writer.L = &q.mu
	q.batchSize = int64(batchSize)
	q.buffer = make([]Value, ceil(size+1, batchSize))
}

func (q *SPSCrMC) next(i int64) int64 {
	r := i + 1
	if r >= int64(len(q.buffer)) {
		r = 0
	}
	return r
}

func (q *SPSCrMC) Send(v Value) bool {
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

func (q *SPSCrMC) FlushSend() {
	q.mu.Lock()
	q.write = q.nextWrite
	q.writeBatch = 0
	q.mu.Unlock()
	q.reader.Signal()
}

func (q *SPSCrMC) Recv(v *Value) bool {
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
		q.FlushRecv()
	}

	return true
}

func (q *SPSCrMC) FlushRecv() {
	q.mu.Lock()
	q.read = q.nextRead
	q.readBatch = 0
	q.mu.Unlock()
	q.writer.Signal()
}
