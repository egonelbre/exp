package compare

import (
	"slices"
	"strconv"
	"testing"
)

const N = 10000000
const K = N/5 - 1

var alpha = float32(1.0)
var xs = makeFloat32(N)
var ys = makeFloat32(N)

type goAxpyDecl struct {
	name string
	fn   func(alpha float32, xs []float32, incx uintptr, ys []float32, incy uintptr, n uintptr)
}

type axpyTestCase struct {
	alpha  float32
	xs     []float32
	incx   uintptr
	ys     []float32
	incy   uintptr
	expect []float32
}

var axpyTestCases = []axpyTestCase{
	0: {
		alpha: 1.3,
		xs:    []float32{1}, incx: 1,
		ys: []float32{2}, incy: 1,
		expect: []float32{3.3},
	},
	1: {
		alpha: 1.3,
		xs:    []float32{1}, incx: 1,
		ys: []float32{2}, incy: 1,
		expect: []float32{3.3},
	},
	2: {
		alpha: 1.3,
		xs:    []float32{}, incx: 1,
		ys: []float32{}, incy: 1,
		expect: []float32{},
	},
	3: {
		alpha: 2,
		xs:    []float32{8}, incx: 1,
		ys: []float32{1}, incy: 1,
		expect: []float32{17},
	},
	4: {
		alpha: 3,
		xs:    []float32{1, 2, 3, 4}, incx: 1,
		ys: []float32{1, 1, 1, 1}, incy: 1,
		expect: []float32{4, 7, 10, 13},
	},
	5: {
		alpha: 3,
		xs:    []float32{1, 2, 3, 5}, incx: 1,
		ys: []float32{-7, -11, -13, -17}, incy: 1,
		expect: []float32{-4, -5, -4, -2},
	},
	6: {
		alpha: 2.3,
		xs:    []float32{-4691.1084, 1844.69, 1986.0142, 3274.4463, 4433.3447}, incx: 1,
		ys: []float32{2613.8955, -355.9663, 898.40283, 2144.5698, 4446.6465}, incy: 1,
		expect: []float32{-8175.6533, 3886.8203, 5466.2354, 9675.796, 14643.339},
	},
	7: {
		alpha: 3,
		xs:    []float32{1, 2, 3, 5, 1, 2, 3, 5}, incx: 1,
		ys: []float32{-7, -11, -13, -17, -7, -11, -13, -17}, incy: 1,
		expect: []float32{-4, -5, -4, -2, -4, -5, -4, -2},
	},
	8: {
		alpha: 3,
		xs:    []float32{0, 1, 2, 3, 5, 1, 2, 3, 5}, incx: 1,
		ys: []float32{0, -7, -11, -13, -17, -7, -11, -13, -17}, incy: 1,
		expect: []float32{0, -4, -5, -4, -2, -4, -5, -4, -2},
	},
}

var goAxypDecls = []goAxpyDecl{
	{name: "AxpyBasic", fn: AxpyBasic},
	{name: "AxpyUnsafe", fn: AxpyUnsafe},
	{name: "AxpyUnsafeX", fn: AxpyUnsafeX},
	{name: "AxpyUnsafeInline", fn: AxpyUnsafeInline},
	{name: "AxpyPointer", fn: AxpyPointer},
	{name: "AxpyPointerLoop", fn: AxpyPointerLoop},
	{name: "AxpyPointerLoopX", fn: AxpyPointerLoopX},
	{name: "AxpyBasicR4", fn: AxpyBasicR4},
	{name: "AxpyBasicXR4", fn: AxpyBasicXR4},
	{name: "AxpyUnsafeR4", fn: AxpyUnsafeR4},
	{name: "AxpyUnsafeXR4", fn: AxpyUnsafeXR4},
	{name: "AxpyUnsafeR8", fn: AxpyUnsafeR8},
	{name: "AxpyUnsafeXR8", fn: AxpyUnsafeXR8},
	{name: "AxpyUnsafeInlineR4", fn: AxpyUnsafeInlineR4},
	{name: "AxpyUnsafeInlineXR4", fn: AxpyUnsafeInlineXR4},
	{name: "AxpyUnsafeInlineR8", fn: AxpyUnsafeInlineR8},
	{name: "AxpyUnsafeInlineXR8", fn: AxpyUnsafeInlineXR8},
	{name: "AxpyPointerR4", fn: AxpyPointerR4},
	{name: "AxpyPointerLoopR4", fn: AxpyPointerLoopR4},
	{name: "AxpyPointerLoopXR4", fn: AxpyPointerLoopXR4},
	{name: "AxpyPointerLoopInterleaveR4", fn: AxpyPointerLoopInterleaveR4},
	{name: "AxpyPointerLoopInterleaveXR4", fn: AxpyPointerLoopInterleaveXR4},
	{name: "AxpyPointerR4Alt", fn: AxpyPointerR4Alt},
}

func BenchmarkGo(b *testing.B) {
	for _, axpy := range goAxypDecls {
		b.Run(axpy.name, func(b *testing.B) {
			fn := axpy.fn
			for i := 0; i < b.N; i++ {
				fn(alpha, xs[i&3:], 2, ys, 4, K)
			}
		})
	}
}

func TestGo(t *testing.T) {
	for _, axpy := range goAxypDecls {
		t.Run(axpy.name, func(t *testing.T) {
			for i, test := range axpyTestCases {
				t.Run(strconv.Itoa(i), func(t *testing.T) {
					lxs := slices.Clone(test.xs)
					lys := slices.Clone(test.ys)

					axpy.fn(test.alpha, lxs, test.incx, lys, test.incy, uintptr(len(lxs)))

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

func makeFloat32(n int) []float32 {
	r := make([]float32, n)
	for i := range r {
		r[i] = float32(i)
	}
	return r
}
