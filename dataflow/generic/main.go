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

func MkPick2[N Node2[A, B], A, B any](in N) Pick2[N, A, B] {
	return Pick2[N, A, B]{Value: in}
}

func (p Pick2[N, A, B]) Out1() B { return p.Value.Out2() }

type Parameter[T any] struct {
	Name         string
	Value        T
	DefaultValue T
}

func (p *Parameter[T]) Out1() T { return p.Value }

type SinCos[In Node1[A], A float64] struct {
	Value In

	Sin, Cos A
}

func MkSinCos[In Node1[A], A float64](v In) *SinCos[In, A] {
	return &SinCos[In, A]{Value: v}
}

func (v *SinCos[In, A]) Update() {
	sn, cs := math.Sincos(float64(v.Value.Out1()))
	v.Sin, v.Cos = A(sn), A(cs)
}

func (v *SinCos[In, A]) Out1() A { return v.Sin }
func (v *SinCos[In, A]) Out2() A { return v.Cos }

type Probe[In Node1[A], A any] struct {
	Value In
}

func MkProbe[In Node1[A], A any](v In) Probe[In, A] {
	return Probe[In, A]{Value: v}
}

func (p *Probe[In, A]) Out1() A {
	v := p.Value.Out1()
	fmt.Println(v)
	return v
}

func main() {
	sinCos := MkSinCos(
		&Parameter[float64]{
			Name:         "Hello",
			Value:        10,
			DefaultValue: 5,
		},
	)

	probeSin := MkProbe(sinCos)
	probeCos := MkProbe(MkPick2(sinCos))

	sinCos.Update()
	probeSin.Out1()
	probeCos.Out1()
}
