package main

import (
	"fmt"
	"math"
)

type Node interface {
	Process()
}

type Parameter[T any] struct {
	Name         string
	Value        T
	DefaultValue T
}

func (p *Parameter[T]) Process() {}

type SinCos struct {
	In *float64

	Sin float64
	Cos float64
}

func (sincos *SinCos) Process() {
	sincos.Sin, sincos.Cos = math.Sincos(*sincos.In)
}

type Probe[T any] struct {
	In *T
}

func (probe Probe[T]) Process() {
	fmt.Println(*probe.In)
}

func main() {
	hello := &Parameter[float64]{
		Name:         "Hello",
		Value:        10,
		DefaultValue: 5,
	}

	sincos := &SinCos{In: &hello.Value}

	probeSin := Probe[float64]{In: &sincos.Sin}
	probeCos := Probe[float64]{In: &sincos.Cos}

	nodes := []Node{
		hello,
		sincos,
		probeSin,
		probeCos,
	}

	for _, n := range nodes {
		n.Process()
	}
}
