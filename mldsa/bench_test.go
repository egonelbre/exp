package mldsa

import (
	"math/rand/v2"
	"runtime"
	"testing"
)

/*

go1.26-devel_3fe246ae0f

goos: windows
goarch: amd64
pkg: github.com/egonelbre/exp/mldsa
cpu: AMD Ryzen Threadripper 2950X 16-Core Processor
BenchmarkNTT/Original-32                  354405              3370 ns/op
BenchmarkNTT/SIMD-32                      479840              2585 ns/op
BenchmarkNTTMul/Original-32              2596790               464.8 ns/op
BenchmarkNTTMul/SIMD-32                  5010793               244.4 ns/op
BenchmarkInverseNTT/Original-32           399896              3054 ns/op
BenchmarkInverseNTT/SIMD-32               482515              2468 ns/op
BenchmarkFieldReduceOnce/Original-32            186527607                6.489 ns/op
BenchmarkFieldReduceOnce/SIMD-32                827229642                1.618 ns/op
BenchmarkFieldAdd/Original-32                   153656426                7.333 ns/op
BenchmarkFieldAdd/SIMD-32                       794731459                1.576 ns/op
BenchmarkFieldSub/Original-32                   175429568                6.820 ns/op
BenchmarkFieldSub/SIMD-32                       657936218                1.828 ns/op
BenchmarkFieldMontgomeryMul/Original-32         95925561                10.65 ns/op
BenchmarkFieldMontgomeryMul/SIMD-32             157265240                7.635 ns/op
PASS
ok      github.com/egonelbre/exp/mldsa  22.753s

*/

func BenchmarkNTT(b *testing.B) {
	var x RingElement
	for i := range x {
		x[i] = FieldElement(rand.Uint32())
	}

	b.Run("Original", func(b *testing.B) {
		x := x
		for range b.N {
			x = RingElement(NTT(x))
		}
		runtime.KeepAlive(x)
	})
	b.Run("SIMD", func(b *testing.B) {
		x := x
		for range b.N {
			x = RingElement(NTT8(x))
		}
		runtime.KeepAlive(x)
	})
}

func BenchmarkNTTMul(b *testing.B) {
	var x, y NTTElement
	for i := range x {
		x[i] = FieldElement(rand.Uint32())
		y[i] = FieldElement(rand.Uint32())
	}

	b.Run("Original", func(b *testing.B) {
		x, y := x, y
		for range b.N {
			x = NTTMul(x, y)
		}
		runtime.KeepAlive(x)
	})
	b.Run("SIMD", func(b *testing.B) {
		x, y := x, y
		for range b.N {
			x = NTTMul8(x, y)
		}
		runtime.KeepAlive(x)
	})
}

func BenchmarkInverseNTT(b *testing.B) {
	var x NTTElement
	for i := range x {
		x[i] = FieldElement(rand.Uint32())
	}

	b.Run("Original", func(b *testing.B) {
		x := x
		for range b.N {
			x = NTTElement(InverseNTT(x))
		}
		runtime.KeepAlive(x)
	})
	b.Run("SIMD", func(b *testing.B) {
		x := x
		for range b.N {
			x = NTTElement(InverseNTT8(x))
		}
		runtime.KeepAlive(x)
	})
}

func BenchmarkFieldReduceOnce(b *testing.B) {
	var x [8]FieldElement
	for i := range x {
		x[i] = FieldElement(rand.Uint32())
	}

	b.Run("Original", func(b *testing.B) {
		x := x
		for range b.N {
			for i := range x {
				x[i] = FieldReduceOnce(uint32(x[i]))
			}
		}
		runtime.KeepAlive(x)
	})

	b.Run("SIMD", func(b *testing.B) {
		x := FieldElement8FromArray(x)
		for range b.N {
			x = FieldReduceOnce8(x)
		}
		runtime.KeepAlive(x)
	})
}

func BenchmarkFieldAdd(b *testing.B) {
	var x, y [8]FieldElement
	for i := range x {
		x[i] = FieldElement(rand.Uint32())
		y[i] = FieldElement(rand.Uint32())
	}

	b.Run("Original", func(b *testing.B) {
		x, y := x, y
		for range b.N {
			for i := range x {
				x[i] = FieldAdd(x[i], y[i])
			}
		}
		runtime.KeepAlive(x)
	})

	b.Run("SIMD", func(b *testing.B) {
		x := FieldElement8FromArray(x)
		y := FieldElement8FromArray(y)
		for range b.N {
			x = FieldAdd8(x, y)
		}
		runtime.KeepAlive(x)
	})
}

func BenchmarkFieldSub(b *testing.B) {
	var x, y [8]FieldElement
	for i := range x {
		x[i] = FieldElement(rand.Uint32())
		y[i] = FieldElement(rand.Uint32())
	}

	b.Run("Original", func(b *testing.B) {
		x, y := x, y
		for range b.N {
			for i := range x {
				x[i] = FieldSub(x[i], y[i])
			}
		}
		runtime.KeepAlive(x)
	})

	b.Run("SIMD", func(b *testing.B) {
		x := FieldElement8FromArray(x)
		y := FieldElement8FromArray(y)
		for range b.N {
			x = FieldSub8(x, y)
		}
		runtime.KeepAlive(x)
	})
}

func BenchmarkFieldMontgomeryMul(b *testing.B) {
	var x, y [8]FieldElement
	for i := range x {
		x[i] = FieldElement(rand.Uint32())
		y[i] = FieldElement(rand.Uint32())
	}

	b.Run("Original", func(b *testing.B) {
		x, y := x, y
		for range b.N {
			for i := range x {
				x[i] = FieldMontgomeryMul(x[i], y[i])
			}
		}
		runtime.KeepAlive(x)
	})

	b.Run("SIMD", func(b *testing.B) {
		x := FieldElement8FromArray(x)
		y := FieldElement8FromArray(y)
		for range b.N {
			x = FieldMontgomeryMul8(x, y)
		}
		runtime.KeepAlive(x)
	})
}
