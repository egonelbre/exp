package compare

import (
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func BenchmarkAsm(b *testing.B) {
	for _, decl := range axpyDecls {
		b.Run(decl.name, func(b *testing.B) {
			fn := decl.fn
			for i := 0; i < b.N; i++ {
				fn(alpha, &xs[i&3], 2, &ys[0], 4, K)
			}
		})
	}
}

func FuzzAsm(f *testing.F) {
	f.Add(int64(0), uint8(0))
	f.Fuzz(func(t *testing.T, seed1 int64, bn byte) {
		n := int(bn) % 8
		xs, ys := make([]float32, n), make([]float32, n)

		const Scale = 10000
		rng := rand.New(rand.NewSource(seed1))
		for i := range xs {
			xs[i] = rng.Float32()*Scale - Scale/2
			ys[i] = rng.Float32()*Scale - Scale/2
		}

		expectys := slices.Clone(ys)
		alpha := float32(2.3)
		AxpyBasic(alpha, xs, 1, expectys, 1, uintptr(len(xs)))

		for _, axpy := range axpyDecls {
			lxs := slices.Clone(xs)
			lys := slices.Clone(ys)
			axpy.fn(alpha, unsafe.SliceData(lxs), 1, unsafe.SliceData(lys), 1, uintptr(len(xs)))
			if !slices.Equal(lys, expectys) {
				t.Errorf("%q wrong result\n\tgot=%v\n\texp=%v\n\tal=%v\n\txs=%v\n\tys=%v", axpy.name, lys, expectys, alpha, xs, ys)
			}
			if !slices.Equal(lxs, xs) {
				t.Errorf("%q xs modified\n\tgot=%v\n\texp=%v\n\tal=%v\n\txs=%v\n\tys=%v", axpy.name, lxs, xs, alpha, xs, ys)
			}
		}
	})
}

func TestAsm(t *testing.T) {
	test := func(alpha float32, xs, ys []float32, expect []float32) {
		t.Run("", func(t *testing.T) {
			for _, axpy := range axpyDecls {
				t.Run(axpy.name, func(t *testing.T) {
					lxs := slices.Clone(xs)
					lys := slices.Clone(ys)

					axpy.fn(alpha, unsafe.SliceData(lxs), 1, unsafe.SliceData(lys), 1, uintptr(len(lxs)))

					if !slices.Equal(lys, expect) {
						t.Errorf("wrong result\n\tgot=%v\n\texp=%v\n\tal=%v\n\txs=%v\n\tys=%v", lys, expect, alpha, xs, ys)
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

	test(3,
		[]float32{1, 2, 3, 4},
		[]float32{1, 1, 1, 1},
		[]float32{4, 7, 10, 13})

	test(3,
		[]float32{1, 2, 3, 5},
		[]float32{-7, -11, -13, -17},
		[]float32{-4, -5, -4, -2})

	test(2.3,
		[]float32{-4691.1084, 1844.69, 1986.0142, 3274.4463, 4433.3447},
		[]float32{2613.8955, -355.9663, 898.40283, 2144.5698, 4446.6465},
		[]float32{-8175.6533, 3886.8203, 5466.2354, 9675.796, 14643.339})
}
