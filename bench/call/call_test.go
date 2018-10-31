package call

import (
	"testing"

	"github.com/egonelbre/exp/bench/call/asm"
)

func BenchmarkNopGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Nop()
	}
}

func BenchmarkNopCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CNop()
	}
}

func BenchmarkAsm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asm.Nop()
	}
}

func BenchmarkArgs1(b *testing.B) {
	n := NewDeviceCreateInfo
	x0 := n()
	for i := 0; i < b.N; i++ {
		Args1(x0)
	}
	Release(x0)
}

func BenchmarkArgs2(b *testing.B) {
	n := NewDeviceCreateInfo
	x0, x1 := n(), n()
	for i := 0; i < b.N; i++ {
		Args2(x0, x1)
	}
	Release(x0, x1)
}

func BenchmarkArgs3(b *testing.B) {
	n := NewDeviceCreateInfo
	x0, x1, x2 := n(), n(), n()
	for i := 0; i < b.N; i++ {
		Args3(x0, x1, x2)
	}
	Release(x0, x1, x2)
}

func BenchmarkArgs4(b *testing.B) {
	n := NewDeviceCreateInfo
	x0, x1, x2, x3 := n(), n(), n(), n()
	for i := 0; i < b.N; i++ {
		Args4(x0, x1, x2, x3)
	}
	Release(x0, x1, x2, x3)
}

func BenchmarkArgs8(b *testing.B) {
	n := NewDeviceCreateInfo
	x0, x1, x2, x3, x4, x5, x6, x7 := n(), n(), n(), n(), n(), n(), n(), n()
	for i := 0; i < b.N; i++ {
		Args8(x0, x1, x2, x3, x4, x5, x6, x7)
	}
	Release(x0, x1, x2, x3, x4, x5, x6, x7)
}

func BenchmarkCArgs1(b *testing.B) {
	n := NewDeviceCreateInfo
	x0 := n()
	for i := 0; i < b.N; i++ {
		CArgs1(x0)
	}
	Release(x0)
}

func BenchmarkCArgs2(b *testing.B) {
	n := NewDeviceCreateInfo
	x0, x1 := n(), n()
	for i := 0; i < b.N; i++ {
		CArgs2(x0, x1)
	}
	Release(x0, x1)
}

func BenchmarkCArgs3(b *testing.B) {
	n := NewDeviceCreateInfo
	x0, x1, x2 := n(), n(), n()
	for i := 0; i < b.N; i++ {
		CArgs3(x0, x1, x2)
	}
	Release(x0, x1, x2)
}

func BenchmarkCArgs4(b *testing.B) {
	n := NewDeviceCreateInfo
	x0, x1, x2, x3 := n(), n(), n(), n()
	for i := 0; i < b.N; i++ {
		CArgs4(x0, x1, x2, x3)
	}
	Release(x0, x1, x2, x3)
}

func BenchmarkCArgs8(b *testing.B) {
	n := NewDeviceCreateInfo
	x0, x1, x2, x3, x4, x5, x6, x7 := n(), n(), n(), n(), n(), n(), n(), n()
	for i := 0; i < b.N; i++ {
		CArgs8(x0, x1, x2, x3, x4, x5, x6, x7)
	}
	Release(x0, x1, x2, x3, x4, x5, x6, x7)
}
