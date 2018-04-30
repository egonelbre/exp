package queue

import "github.com/pltr/onering"

// SPSC wrapper

type SPSCrs_one struct {
	onering.SPSC
	capacity int
}

func NewSPSCrs_one(size int) *SPSCrs_one {
	q := &SPSCrs_one{}
	if size <= 1 {
		size = 2
	}
	x := int(nextPowerOfTwo(uint32(size)))
	q.Init(uint(x))
	q.capacity = x
	return q
}

func (q *SPSCrs_one) Cap() int { return q.capacity - 1 }

func (q *SPSCrs_one) Send(v Value) bool {
	q.Put(v)
	return true
}

func (q *SPSCrs_one) Recv(v *Value) bool {
	return q.Get(v)
}

func (q *SPSCrs_one) BatchRecv(fn func(Value)) bool {
	q.Consume(fn)
	return true
}

// SPMC wrapper

type SPMCrs_one struct {
	onering.SPMC
	capacity int
}

func NewSPMCrs_one(size int) *SPMCrs_one {
	q := &SPMCrs_one{}
	if size <= 1 {
		size = 2
	}
	x := int(nextPowerOfTwo(uint32(size)))
	q.Init(uint(x))
	q.capacity = x
	return q
}

func (q *SPMCrs_one) MultipleConsumers() {}

func (q *SPMCrs_one) Cap() int { return q.capacity }

func (q *SPMCrs_one) Send(v Value) bool {
	q.Put(v)
	return true
}

func (q *SPMCrs_one) Recv(v *Value) bool {
	return q.Get(v)
}

// MPSC wrapper

type MPSCrs_one struct {
	onering.MPSC
	capacity int
}

func NewMPSCrs_one(size int) *MPSCrs_one {
	q := &MPSCrs_one{}
	if size <= 1 {
		size = 2
	}
	x := int(nextPowerOfTwo(uint32(size)))
	q.Init(uint(x))
	q.capacity = x
	return q
}

func (q *MPSCrs_one) MultipleProducers() {}

func (q *MPSCrs_one) Cap() int { return q.capacity - 1 }

func (q *MPSCrs_one) Send(v Value) bool {
	q.Put(v)
	return true
}

func (q *MPSCrs_one) Recv(v *Value) bool {
	return q.Get(v)
}

func (q *MPSCrs_one) BatchRecv(fn func(Value)) bool {
	q.Consume(fn)
	return true
}
