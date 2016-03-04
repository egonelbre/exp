package ring

import (
	"runtime"
	"sync/atomic"
)

// Implementation is based on
//   http://www.1024cores.net/home/lock-free-algorithms/queues/bounded-mpmc-queue

const (
	Size     = 128 << 10 // ~128000
	Mask     = Size - 1
	CellSize = 4 << 10 // 4kb
)

type Cell struct {
	sequence int64
	Size     int64
	Data     [CellSize]byte
}

type Ring struct {
	writeat  int64
	readfrom int64
	buffer   [Size]Cell
}

func New() *Ring {
	ring := &Ring{}
	for i := range ring.buffer {
		ring.buffer[i].sequence = int64(i)
	}
	return ring
}

func (ring *Ring) Allocate() *Cell {
	busywait := 16
	for {
		pos := atomic.LoadInt64(&ring.writeat)
		cell := &ring.buffer[pos&Mask]
		seq := atomic.LoadInt64(&cell.sequence)
		if seq-pos == 0 {
			if atomic.CompareAndSwapInt64(&ring.writeat, pos, pos+1) {
				return cell
			}
		}

		if busywait < 0 {
			runtime.Gosched()
		} else {
			busywait--
		}
	}
}

func (cell *Cell) Enqueue() { atomic.AddInt64(&cell.sequence, 1) }

func (ring *Ring) Dequeue() (cell *Cell) {
	busywait := 16
	for {
		pos := atomic.LoadInt64(&ring.readfrom)
		cell = &ring.buffer[pos&Mask]
		seq := atomic.LoadInt64(&cell.sequence)
		dif := seq - (pos + 1)
		if dif == 0 {
			if atomic.CompareAndSwapInt64(&ring.readfrom, pos, pos+1) {
				return cell
			}
		}

		if busywait < 0 {
			runtime.Gosched()
		} else {
			busywait--
		}
	}
}

func (cell *Cell) Release() { atomic.AddInt64(&cell.sequence, Size-1) }
