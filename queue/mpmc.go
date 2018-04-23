package queue

import (
	"sync"
	"sync/atomic"
)

// Based on https://docs.google.com/document/d/1yIAYmbvL3JxOKOjuCyon7JhW4cSv1wy5hC0ApeGMV9s/pub

type Value = int

type MPMC struct {
	sendx  uint64
	recvx  uint64
	buffer []mpmcelement

	mu    sync.Mutex
	sendq sync.Cond
	recvq sync.Cond
}

func NewMPMC(size int) *MPMC {
	mpmc := &MPMC{
		sendx:  0,
		recvx:  1 << 32,
		buffer: make([]mpmcelement, size, size),
	}
	mpmc.sendq.L = &mpmc.mu
	mpmc.recvq.L = &mpmc.mu
	return mpmc
}

type mpmcelement struct {
	lap uint32
	val Value
}

func (q *MPMC) cap() uint32 {
	return uint32(len(q.buffer))
}

func (q *MPMC) full() bool {
	x := atomic.LoadUint64(&q.sendx)
	lap, pos := uint32(x>>32), uint32(x)
	e := &q.buffer[pos]
	elap := atomic.LoadUint32(&e.lap)
	return int32(lap-elap) > 0
}

func (q *MPMC) empty() bool {
	x := atomic.LoadUint64(&q.recvx)
	lap, pos := uint32(x>>32), uint32(x)
	e := &q.buffer[pos]
	elap := atomic.LoadUint32(&e.lap)
	return int32(lap-elap) > 0
}

func (q *MPMC) Send(value Value) bool {
	for {
		if q.trySend(&value) {
			q.recvq.Signal()
			return true
		} else {
			q.mu.Lock()
			if q.full() {
				q.sendq.Wait()
			}
			q.mu.Unlock()
		}
	}
	return false
}

func (q *MPMC) RecvValue() (Value, bool) {
	var value Value
	ok := q.Recv(&value)
	return value, ok
}

func (q *MPMC) TryRecvValue() (Value, bool) {
	var value Value
	ok := q.TryRecv(&value)
	return value, ok
}

func (q *MPMC) Recv(value *Value) bool {
	for {
		if q.tryRecv(value) {
			q.sendq.Signal()
			return true
		} else {
			q.mu.Lock()
			if q.empty() {
				q.recvq.Wait()
			}
			q.mu.Unlock()
		}
	}
	return false
}

func (q *MPMC) TrySend(value Value) bool  { return q.trySend(&value) }
func (q *MPMC) TryRecv(value *Value) bool { return q.tryRecv(value) }

func (q *MPMC) trySend(value *Value) bool {
	for {
		x := atomic.LoadUint64(&q.sendx)
		lap, pos := uint32(x>>32), uint32(x)
		e := &q.buffer[pos]
		elap := atomic.LoadUint32(&e.lap)
		if lap == elap {
			var newx uint64

			// The element is ready for writing on this lap.
			// Try to claim the right to write to this element.
			if pos+1 < q.cap() {
				newx = x + 1 // just increase the pos
			} else {
				newx = uint64(lap+2) << 32
			}

			if !atomic.CompareAndSwapUint64(&q.sendx, x, newx) {
				continue // lose the race, retry
			}

			// We own the element, do non-atomic write.
			e.val = *value

			// Make the element available for reading.
			atomic.StoreUint32(&e.lap, e.lap+1)
			return true
		} else if int32(lap-elap) > 0 {
			// The element is not yet read on the previous lap,
			// the chan is full.
			return false
		} else {
			// The element has already been written on this lap,
			// this means that q.sendx has been changed as well,
			// retry.
		}
	}
}
func (q *MPMC) tryRecv(result *Value) bool {
	var empty Value
	for {
		x := atomic.LoadUint64(&q.recvx)
		lap, pos := uint32(x>>32), uint32(x)
		e := &q.buffer[pos]
		elap := atomic.LoadUint32(&e.lap)
		if lap == elap {
			var newx uint64

			// The element is ready for writing on this lap.
			// Try to claim the right to write to this element.
			if pos+1 < q.cap() {
				newx = x + 1 // just increase the pos
			} else {
				newx = uint64(lap+2) << 32
			}
			if !atomic.CompareAndSwapUint64(&q.recvx, x, newx) {
				continue // lose the race, retry
			}

			// We own the element, do non-atomic read
			*result = e.val
			e.val = empty

			// Make the element available for writing
			atomic.StoreUint32(&e.lap, e.lap+1)
			return true
		} else if int32(lap-elap) > 0 {
			// The element is not yet written on the lap,
			// the chan is empty.
			return false
		} else {
			// The element has already been read on this lap,
			// this means that q.recvx has been changed as well,
			// retry.
		}
	}
}
