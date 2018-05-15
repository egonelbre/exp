package vector_test

import (
	"runtime"
	"testing"
)

type Struct struct{ X, Y, Z float32 }
type Array [3]float32
type ArrayStruct struct{ _0, _1, _2 float32 }

func (a Struct) Add(b Struct) Struct { return Struct{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a Array) Add(b Array) Array {
	r := Array{}
	for i, bv := range b {
		r[i] = a[i] + bv
	}
	return r
}
func (a Array) Add2(b Array) Array {
	r := Array{}
	for i := range b {
		r[i] = a[i] + b[i]
	}
	return r
}
func (a Array) Add3(b Array) Array {
	r := Array{}
	r[0] = a[0] + b[0]
	r[1] = a[1] + b[1]
	r[2] = a[2] + b[2]
	return r
}

func (a ArrayStruct) Len() int { return 3 }
func (a ArrayStruct) At(i int) float32 {
	switch i {
	case 0:
		return a._0
	case 1:
		return a._1
	case 2:
		return a._2
	default:
		return 0 // panic("")
	}
}
func (a *ArrayStruct) Set(i int, v float32) {
	switch i {
	case 0:
		a._0 = v
	case 1:
		a._1 = v
	case 2:
		a._2 = v
	default:
		// panic("")
	}
}

func (a ArrayStruct) Add(b ArrayStruct) ArrayStruct {
	r := ArrayStruct{}
	for i := 0; i < a.Len(); i++ {
		r.Set(i, a.At(i)+b.At(i))
	}
	return r
}
func (a ArrayStruct) Add3(b ArrayStruct) ArrayStruct {
	r := ArrayStruct{}
	r.Set(0, a.At(0)+b.At(0))
	r.Set(1, a.At(1)+b.At(1))
	r.Set(2, a.At(2)+b.At(2))
	return r
}

func BenchmarkStruct(t *testing.B) {
	a, b := Struct{1, 1, 1}, Struct{1, 1, 1}
	for i := 0; i < t.N; i++ {
		a = a.Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

func BenchmarkArray(t *testing.B) {
	a, b := Array{1, 1, 1}, Array{1, 1, 1}
	for i := 0; i < t.N; i++ {
		a = a.Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkArray2(t *testing.B) {
	a, b := Array{1, 1, 1}, Array{1, 1, 1}
	for i := 0; i < t.N; i++ {
		a = a.Add2(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkArray3(t *testing.B) {
	a, b := Array{1, 1, 1}, Array{1, 1, 1}
	for i := 0; i < t.N; i++ {
		a = a.Add3(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

func BenchmarkArrayStruct(t *testing.B) {
	a, b := ArrayStruct{1, 1, 1}, ArrayStruct{1, 1, 1}
	for i := 0; i < t.N; i++ {
		a = a.Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

func BenchmarkArrayStruct3(t *testing.B) {
	a, b := ArrayStruct{1, 1, 1}, ArrayStruct{1, 1, 1}
	for i := 0; i < t.N; i++ {
		a = a.Add3(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
