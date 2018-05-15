package vector_test

import (
	"runtime"
	"testing"
)

var V64s1 = V64s3{1, 1, 1}

type V64s3 struct{ X, Y, Z float64 }

func (a *V64s3) PP_Add(b *V64s3)         { a.X += b.X; a.Y += b.Y; a.Z += b.Z }
func (a *V64s3) PN_Add(b V64s3)          { a.X += b.X; a.Y += b.Y; a.Z += b.Z }
func (a *V64s3) PPP_Add(b *V64s3) *V64s3 { return &V64s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V64s3) NNN_Add(b V64s3) V64s3    { return V64s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a *V64s3) PPN_Add(b *V64s3) V64s3  { return V64s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a *V64s3) PNP_Add(b V64s3) *V64s3  { return &V64s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V64s3) NPP_Add(b *V64s3) *V64s3  { return &V64s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V64s3) NNP_Add(b V64s3) *V64s3   { return &V64s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V64s3) NPN_Add(b *V64s3) V64s3   { return V64s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a *V64s3) PNN_Add(b V64s3) V64s3   { return V64s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }

func (a *V64s3) PPZ_Add(b *V64s3, r *V64s3) { *r = V64s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V64s3) NPZ_Add(b *V64s3, r *V64s3)  { *r = V64s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a *V64s3) PNZ_Add(b V64s3, r *V64s3)  { *r = V64s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V64s3) NNZ_Add(b V64s3, r *V64s3)   { *r = V64s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }

func (a *V64s3) PPY_Add(b *V64s3, r *V64s3) { r.X = a.X + b.X; r.Y = a.Y + b.Y; r.Z = a.Z + b.Z }
func (a V64s3) NPY_Add(b *V64s3, r *V64s3)  { r.X = a.X + b.X; r.Y = a.Y + b.Y; r.Z = a.Z + b.Z }
func (a *V64s3) PNY_Add(b V64s3, r *V64s3)  { r.X = a.X + b.X; r.Y = a.Y + b.Y; r.Z = a.Z + b.Z }
func (a V64s3) NNY_Add(b V64s3, r *V64s3)   { r.X = a.X + b.X; r.Y = a.Y + b.Y; r.Z = a.Z + b.Z }

func Benchmark64_Add_PP_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a.PP_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PN_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a.PN_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PPP_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a = *a.PPP_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NNN_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a = a.NNN_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PPN_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a = a.PPN_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PNP_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a = *a.PNP_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NPP_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a = *a.NPP_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NNP_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a = *a.NNP_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NPN_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a = a.NPN_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PNN_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a = a.PNN_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

func Benchmark64_Add_PPZ_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a.PPZ_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NPZ_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a.NPZ_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PNZ_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a.PNZ_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NNZ_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a.NNZ_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

func Benchmark64_Add_PPY_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a.PPY_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NPY_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a.NPY_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PNY_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a.PNY_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NNY_V64s3(t *testing.B) {
	a, b := V64s1, V64s1
	for i := 0; i < t.N; i++ {
		a.NNY_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

var V64a1 = V64a3{1, 1, 1}

type V64a3 [3]float64

func (a *V64a3) PP_Add(b *V64a3)         { a[0] += b[0]; a[1] += b[1]; a[2] += b[2] }
func (a *V64a3) PN_Add(b V64a3)          { a[0] += b[0]; a[1] += b[1]; a[2] += b[2] }
func (a *V64a3) PPP_Add(b *V64a3) *V64a3 { return &V64a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a V64a3) NNN_Add(b V64a3) V64a3    { return V64a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a *V64a3) PPN_Add(b *V64a3) V64a3  { return V64a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a *V64a3) PNP_Add(b V64a3) *V64a3  { return &V64a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a V64a3) NPP_Add(b *V64a3) *V64a3  { return &V64a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a V64a3) NNP_Add(b V64a3) *V64a3   { return &V64a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a V64a3) NPN_Add(b *V64a3) V64a3   { return V64a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a *V64a3) PNN_Add(b V64a3) V64a3   { return V64a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }

func (a *V64a3) PPZ_Add(b *V64a3, r *V64a3) { *r = V64a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a V64a3) NPZ_Add(b *V64a3, r *V64a3)  { *r = V64a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a *V64a3) PNZ_Add(b V64a3, r *V64a3)  { *r = V64a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a V64a3) NNZ_Add(b V64a3, r *V64a3)   { *r = V64a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }

func (a *V64a3) PPY_Add(b *V64a3, r *V64a3) {
	r[0] = a[0] + b[0]
	r[1] = a[1] + b[1]
	r[2] = a[2] + b[2]
}
func (a V64a3) NPY_Add(b *V64a3, r *V64a3) { r[0] = a[0] + b[0]; r[1] = a[1] + b[1]; r[2] = a[2] + b[2] }
func (a *V64a3) PNY_Add(b V64a3, r *V64a3) { r[0] = a[0] + b[0]; r[1] = a[1] + b[1]; r[2] = a[2] + b[2] }
func (a V64a3) NNY_Add(b V64a3, r *V64a3)  { r[0] = a[0] + b[0]; r[1] = a[1] + b[1]; r[2] = a[2] + b[2] }

func Benchmark64_Add_PP_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a.PP_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PN_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a.PN_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PPP_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a = *a.PPP_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NNN_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a = a.NNN_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PPN_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a = a.PPN_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PNP_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a = *a.PNP_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NPP_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a = *a.NPP_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NNP_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a = *a.NNP_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NPN_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a = a.NPN_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PNN_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a = a.PNN_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

func Benchmark64_Add_PPZ_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a.PPZ_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NPZ_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a.NPZ_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PNZ_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a.PNZ_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NNZ_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a.NNZ_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

func Benchmark64_Add_PPY_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a.PPY_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NPY_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a.NPY_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_PNY_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a.PNY_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func Benchmark64_Add_NNY_V64a3(t *testing.B) {
	a, b := V64a1, V64a1
	for i := 0; i < t.N; i++ {
		a.NNY_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
