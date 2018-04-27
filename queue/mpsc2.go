package queue

import (
	"sync"
)

type MPSC2 struct {
	_ [8]uint64
	// volatile
	read  int64
	_     [8 - 1]uint64
	write int64
	_     [8 - 1]uint64
	// consumer
	localWrite int64
	nextRead   int64
	readBatch  int64
	_          [8 - 3]uint64
	localRead  int64
	_          [8 - 1]uint64
	// constant
	mu        sync.Mutex
	reader    sync.Cond
	writer    sync.Cond
	batchSize int64
	buffer    []MPSC2Value
}

type MPSC2Value = int64

func NewMPSC2(batchSize, size int) *MPSC2 {
	q := &MPSC2{}
	q.Init(batchSize, size)
	return q
}

func (q *MPSC2) Init(batchSize, size int) {
	q.reader.L = &q.mu
	q.writer.L = &q.mu
	q.batchSize = int64(batchSize)
	q.buffer = make([]MPSC2Value, ceil(size, batchSize))
}

func (q *MPSC2) next(i int64) int64 {
	r := i + 1
	if r >= int64(len(q.buffer)) {
		r = 0
	}
	return r
}

func (q *MPSC2) Send(v MPSC2Value) bool {
	q.mu.Lock()
	nextWrite := q.next(q.write)
	for nextWrite == q.localRead {
		if nextWrite == q.read {
			// channel is full
			q.writer.Wait()
		}
		nextWrite = q.next(q.write)
		q.localRead = q.read
	}
	q.buffer[nextWrite] = v
	q.write = nextWrite
	q.reader.Signal()
	q.mu.Unlock()

	return true
}

func (q *MPSC2) Recv(v *MPSC2Value) bool {
	if q.nextRead == q.localWrite {
		q.mu.Lock()
		if q.nextRead == q.write {
			// channel is empty
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
		q.writer.Broadcast()
		q.mu.Unlock()
	}

	return true
}
