package call

import (
	"runtime"
	"testing"
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

func BenchmarkAdd2(b *testing.B) {
	total := int32(0)
	for i := 0; i < b.N; i++ {
		total = Add2(total, int32(i))
	}
	runtime.KeepAlive(total)
}

func BenchmarkCAdd2(b *testing.B) {
	total := int32(0)
	for i := 0; i < b.N; i++ {
		total = CAdd2(total, int32(i))
	}
	runtime.KeepAlive(total)
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

func BenchmarkArgs4(b *testing.B) {
	n := NewDeviceCreateInfo
	x0, x1, x2, x3 := n(), n(), n(), n()
	for i := 0; i < b.N; i++ {
		Args4(x0, x1, x2, x3)
	}
	Release(x0, x1, x2, x3)
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

func BenchmarkCArgs4(b *testing.B) {
	n := NewDeviceCreateInfo
	x0, x1, x2, x3 := n(), n(), n(), n()
	for i := 0; i < b.N; i++ {
		CArgs4(x0, x1, x2, x3)
	}
	Release(x0, x1, x2, x3)
}
