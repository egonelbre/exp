package vecprocess

import (
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func BenchmarkAsmAxpyInc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsmAxpyInc(alpha, xs[i&3:], ys, K, 2, 4, 0, 0)
	}
}

func BenchmarkAsmAxpyPointer_Align11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsmAxpyPointer_Align11(alpha, &xs[i&3], 2, &ys[0], 4, K)
	}
}

func BenchmarkAsmAxpyPointer_Align16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsmAxpyPointer_Align16(alpha, &xs[i&3], 2, &ys[0], 4, K)
	}
}

var axpyFuncs = []struct {
	name string
	fn   func(alpha float32, xs *float32, incx uintptr, ys *float32, incy uintptr, n uintptr)
}{
	{name: "AsmAxpyPointer_Align11", fn: AsmAxpyPointer_Align11},
	{name: "AsmAxpyPointer_Align16", fn: AsmAxpyPointer_Align16},
}

func FuzzAsm(f *testing.F) {
	f.Add(int64(0), uint8(0))
	f.Fuzz(func(t *testing.T, seed1 int64, bn byte) {
		n := int(bn)
		xs, ys := make([]float32, n), make([]float32, n)

		const Scale = 10000
		rng := rand.New(rand.NewSource(seed1))
		for i := range xs {
			xs[i] = rng.Float32()*Scale - Scale/2
			ys[i] = rng.Float32()*Scale - Scale/2
		}

		expectys := slices.Clone(ys)
		AxpyBasic(2.3, xs, 1, expectys, 1, uintptr(len(xs)))

		for _, axpy := range axpyFuncs {
			rs := slices.Clone(ys)
			axpy.fn(2.3, unsafe.SliceData(xs), 1, unsafe.SliceData(rs), 1, uintptr(len(xs)))
			if !slices.Equal(rs, expectys) {
				t.Errorf("%q wrong result\n\tgot=%v\n\texp=%v\n\tal=%v\n\txs=%v\n\tys=%v", axpy.name, rs, expectys, 2.3, xs, ys)
			}
		}
	})
}

func TestAsm(t *testing.T) {
	test := func(alpha float32, xs, ys []float32, expect []float32) {
		t.Run("", func(t *testing.T) {
			for _, axpy := range axpyFuncs {
				t.Run(axpy.name, func(t *testing.T) {
					lxs := slices.Clone(xs)
					lys := slices.Clone(ys)

					axpy.fn(alpha, unsafe.SliceData(lxs), 1, unsafe.SliceData(lys), 1, uintptr(len(lxs)))

					if !slices.Equal(lys, expect) {
						t.Errorf("%v != %v", lys, expect)
					}
					if !slices.Equal(lxs, xs) {
						t.Errorf("xs modified")
					}
				})
			}
		})
	}

	test(1.3, []float32{1}, []float32{2}, []float32{3.3})
	test(1.3, []float32{}, []float32{}, []float32{})
	test(2, []float32{8}, []float32{1}, []float32{17})
}
