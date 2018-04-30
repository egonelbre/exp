package queue

import (
	"testing"

	"github.com/pltr/onering"
)

var _ SPSC = (*OneringSPSC)(nil)

type OneringSPSC struct {
	onering.SPSC
	capacity int
}

func NewOneringSPSC(size int) *OneringSPSC {
	q := &OneringSPSC{}
	if size <= 1 {
		size = 2
	}
	x := int(nextPowerOfTwo(uint32(size)))
	q.Init(uint(x))
	q.capacity = x
	return q
}

func (q *OneringSPSC) Cap() int { return q.capacity - 1 }

func (q *OneringSPSC) Send(v Value) bool {
	q.Put(v)
	return true
}

func (q *OneringSPSC) Recv(v *Value) bool {
	return q.Get(v)
}

func (q *OneringSPSC) BatchRecv(fn func(Value)) bool {
	q.Consume(fn)
	return true
}

func TestOneringSPSC(t *testing.T) {
	test(t, func(size int) Queue { return NewOneringSPSC(size) })
}

func BenchmarkOneringSPSC(b *testing.B) {
	bench(b, func(size int) Queue { return NewOneringSPSC(size) })
}

var _ SPMC = (*OneringSPMC)(nil)

type OneringSPMC struct {
	onering.SPMC
	capacity int
}

func NewOneringSPMC(size int) *OneringSPMC {
	q := &OneringSPMC{}
	if size <= 1 {
		size = 2
	}
	x := int(nextPowerOfTwo(uint32(size)))
	q.Init(uint(x))
	q.capacity = x
	return q
}

func (q *OneringSPMC) MultipleConsumers() {}

func (q *OneringSPMC) Cap() int { return q.capacity }

func (q *OneringSPMC) Send(v Value) bool {
	q.Put(v)
	return true
}

func (q *OneringSPMC) Recv(v *Value) bool {
	return q.Get(v)
}

func TestOneringSPMC(t *testing.T) {
	test(t, func(size int) Queue { return NewOneringSPMC(size) })
}

func BenchmarkOneringSPMC(b *testing.B) {
	bench(b, func(size int) Queue { return NewOneringSPMC(size) })
}

var _ MPSC = (*OneringMPSC)(nil)

type OneringMPSC struct {
	onering.MPSC
	capacity int
}

func NewOneringMPSC(size int) *OneringMPSC {
	q := &OneringMPSC{}
	if size <= 1 {
		size = 2
	}
	x := int(nextPowerOfTwo(uint32(size)))
	q.Init(uint(x))
	q.capacity = x
	return q
}

func (q *OneringMPSC) MultipleProducers() {}

func (q *OneringMPSC) Cap() int { return q.capacity - 1 }

func (q *OneringMPSC) Send(v Value) bool {
	q.Put(v)
	return true
}

func (q *OneringMPSC) Recv(v *Value) bool {
	return q.Get(v)
}

func (q *OneringMPSC) BatchRecv(fn func(Value)) bool {
	q.Consume(fn)
	return true
}

func TestOneringMPSC(t *testing.T) {
	test(t, func(size int) Queue { return NewOneringMPSC(size) })
}

func BenchmarkOneringMPSC(b *testing.B) {
	bench(b, func(size int) Queue { return NewOneringMPSC(size) })
}
