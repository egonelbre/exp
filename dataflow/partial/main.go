package main

import (
	"fmt"
	"math"
)

type Node1[A any] interface {
	Out1() A
}

type Node2[A, B any] interface {
	Out1() A
	Out2() B
}

type Node3[A, B, C any] interface {
	Out1() A
	Out2() B
	Out3() C
}

type Pick2[N Node2[A, B], A, B any] struct{ Value N }

func Out2[N Node2[A, B], A, B any](v N) *Pick2[N, A, B] {
	return &Pick2[N, A, B]{Value: v}
}

func (p *Pick2[N, A, B]) Out1() B { return p.Value.Out2() }

type Parameter[T any] struct {
	Name         string
	Value        T
	DefaultValue T
}

func (p *Parameter[T]) Out1() T { return p.Value }

type SinCos[A float64] struct {
	Value Node1[A]

	Sin, Cos A
}

func (v *SinCos[A]) Update() {
	sn, cs := math.Sincos(float64(v.Value.Out1()))
	v.Sin, v.Cos = A(sn), A(cs)
}

func (v *SinCos[A]) Out1() A { return v.Sin }
func (v *SinCos[A]) Out2() A { return v.Cos }

type Probe[A any] struct {
	Value Node1[A]
}

func (p *Probe[A]) Out1() A {
	v := p.Value.Out1()
	fmt.Println(v)
	return v
}

func main() {
	sinCos := &SinCos[float64]{
		Value: &Parameter[float64]{
			Name:         "Hello",
			Value:        10,
			DefaultValue: 5,
		},
	}

	probeSin := &Probe[float64]{Value: sinCos}
	probeCos := &Probe[float64]{Value: Out2(sinCos)}

	sinCos.Update()
	probeSin.Out1()
	probeCos.Out1()
}
