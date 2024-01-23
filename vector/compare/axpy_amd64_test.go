package compare

import (
	"math/rand"
	"slices"
	"strconv"
	"testing"
	"unsafe"
)

func BenchmarkAmd64(b *testing.B) {
	for _, decl := range amdAxpyDecls {
		b.Run(decl.name, func(b *testing.B) {
			fn := decl.fn
			for i := 0; i < b.N; i++ {
				fn(alpha, &xs[i&3], 2, &ys[0], 4, K)
			}
		})
	}
}

func FuzzAmd64(f *testing.F) {
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

		for _, axpy := range amdAxpyDecls {
			lxs := slices.Clone(xs)
			lys := slices.Clone(ys)
			axpy.fn(alpha, unsafe.SliceData(lxs), 1, unsafe.SliceData(lys), 1, uintptr(len(xs)))
			if !equalFloats(lys, expectys) {
				t.Errorf("%q wrong result\n\tgot=%v\n\texp=%v\n\tal=%v\n\txs=%v\n\tys=%v", axpy.name, lys, expectys, alpha, xs, ys)
			}
			if !equalFloats(lxs, xs) {
				t.Errorf("%q xs modified\n\tgot=%v\n\texp=%v\n\tal=%v\n\txs=%v\n\tys=%v", axpy.name, lxs, xs, alpha, xs, ys)
			}
		}
	})
}

func TestAmd64(t *testing.T) {
	for _, axpy := range amdAxpyDecls {
		t.Run(axpy.name, func(t *testing.T) {
			for i, test := range axpyTestCases {
				t.Run(strconv.Itoa(i), func(t *testing.T) {
					lxs := slices.Clone(test.xs)
					lys := slices.Clone(test.ys)

					axpy.fn(test.alpha, unsafe.SliceData(lxs), test.incx, unsafe.SliceData(lys), test.incy, test.N())

					if !equalFloats(lys, test.expect) {
						t.Errorf("wrong result\n\tgot=%v\n\texp=%v\n\tal=%v\n\txs=%v\n\tys=%v", lys, test.expect, test.alpha, test.xs, test.ys)
					}
					if !equalFloats(lxs, test.xs) {
						t.Errorf("xs modified")
					}
				})
			}
		})
	}
}
