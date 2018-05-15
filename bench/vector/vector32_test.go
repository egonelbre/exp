package vector_test

import (
	"runtime"
	"testing"
)

var V32s1 = V32s3{1, 1, 1}

type V32s3 struct{ X, Y, Z float32 }

func (a *V32s3) PP_Add(b *V32s3)         { a.X += b.X; a.Y += b.Y; a.Z += b.Z }
func (a *V32s3) PN_Add(b V32s3)          { a.X += b.X; a.Y += b.Y; a.Z += b.Z }
func (a *V32s3) PPP_Add(b *V32s3) *V32s3 { return &V32s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V32s3) NNN_Add(b V32s3) V32s3    { return V32s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a *V32s3) PPN_Add(b *V32s3) V32s3  { return V32s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a *V32s3) PNP_Add(b V32s3) *V32s3  { return &V32s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V32s3) NPP_Add(b *V32s3) *V32s3  { return &V32s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V32s3) NNP_Add(b V32s3) *V32s3   { return &V32s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V32s3) NPN_Add(b *V32s3) V32s3   { return V32s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a *V32s3) PNN_Add(b V32s3) V32s3   { return V32s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }

func (a *V32s3) PPZ_Add(b *V32s3, r *V32s3) { *r = V32s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V32s3) NPZ_Add(b *V32s3, r *V32s3)  { *r = V32s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a *V32s3) PNZ_Add(b V32s3, r *V32s3)  { *r = V32s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V32s3) NNZ_Add(b V32s3, r *V32s3)   { *r = V32s3{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }

func (a *V32s3) PPY_Add(b *V32s3, r *V32s3) { r.X = a.X + b.X; r.Y = a.Y + b.Y; r.Z = a.Z + b.Z }
func (a V32s3) NPY_Add(b *V32s3, r *V32s3)  { r.X = a.X + b.X; r.Y = a.Y + b.Y; r.Z = a.Z + b.Z }
func (a *V32s3) PNY_Add(b V32s3, r *V32s3)  { r.X = a.X + b.X; r.Y = a.Y + b.Y; r.Z = a.Z + b.Z }
func (a V32s3) NNY_Add(b V32s3, r *V32s3)   { r.X = a.X + b.X; r.Y = a.Y + b.Y; r.Z = a.Z + b.Z }

func BenchmarkAdd_PP_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a.PP_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PN_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a.PN_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PPP_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a = *a.PPP_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NNN_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a = a.NNN_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PPN_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a = a.PPN_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PNP_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a = *a.PNP_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NPP_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a = *a.NPP_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NNP_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a = *a.NNP_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NPN_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a = a.NPN_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PNN_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a = a.PNN_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

func BenchmarkAdd_PPZ_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a.PPZ_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NPZ_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a.NPZ_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PNZ_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a.PNZ_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NNZ_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a.NNZ_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

func BenchmarkAdd_PPY_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a.PPY_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NPY_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a.NPY_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PNY_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a.PNY_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NNY_V32s3(t *testing.B) {
	a, b := V32s1, V32s1
	for i := 0; i < t.N; i++ {
		a.NNY_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

var V32a1 = V32a3{1, 1, 1}

type V32a3 [3]float32

func (a *V32a3) PP_Add(b *V32a3)         { a[0] += b[0]; a[1] += b[1]; a[2] += b[2] }
func (a *V32a3) PN_Add(b V32a3)          { a[0] += b[0]; a[1] += b[1]; a[2] += b[2] }
func (a *V32a3) PPP_Add(b *V32a3) *V32a3 { return &V32a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a V32a3) NNN_Add(b V32a3) V32a3    { return V32a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a *V32a3) PPN_Add(b *V32a3) V32a3  { return V32a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a *V32a3) PNP_Add(b V32a3) *V32a3  { return &V32a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a V32a3) NPP_Add(b *V32a3) *V32a3  { return &V32a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a V32a3) NNP_Add(b V32a3) *V32a3   { return &V32a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a V32a3) NPN_Add(b *V32a3) V32a3   { return V32a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a *V32a3) PNN_Add(b V32a3) V32a3   { return V32a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }

func (a *V32a3) PPZ_Add(b *V32a3, r *V32a3) { *r = V32a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a V32a3) NPZ_Add(b *V32a3, r *V32a3)  { *r = V32a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a *V32a3) PNZ_Add(b V32a3, r *V32a3)  { *r = V32a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }
func (a V32a3) NNZ_Add(b V32a3, r *V32a3)   { *r = V32a3{a[0] + b[0], a[1] + b[1], a[2] + b[2]} }

func (a *V32a3) PPY_Add(b *V32a3, r *V32a3) {
	r[0] = a[0] + b[0]
	r[1] = a[1] + b[1]
	r[2] = a[2] + b[2]
}
func (a V32a3) NPY_Add(b *V32a3, r *V32a3) { r[0] = a[0] + b[0]; r[1] = a[1] + b[1]; r[2] = a[2] + b[2] }
func (a *V32a3) PNY_Add(b V32a3, r *V32a3) { r[0] = a[0] + b[0]; r[1] = a[1] + b[1]; r[2] = a[2] + b[2] }
func (a V32a3) NNY_Add(b V32a3, r *V32a3)  { r[0] = a[0] + b[0]; r[1] = a[1] + b[1]; r[2] = a[2] + b[2] }

func BenchmarkAdd_PP_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a.PP_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PN_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a.PN_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PPP_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a = *a.PPP_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NNN_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a = a.NNN_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PPN_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a = a.PPN_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PNP_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a = *a.PNP_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NPP_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a = *a.NPP_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NNP_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a = *a.NNP_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NPN_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a = a.NPN_Add(&b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PNN_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a = a.PNN_Add(b)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

func BenchmarkAdd_PPZ_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a.PPZ_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NPZ_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a.NPZ_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PNZ_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a.PNZ_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NNZ_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a.NNZ_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}

func BenchmarkAdd_PPY_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a.PPY_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NPY_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a.NPY_Add(&b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_PNY_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a.PNY_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
func BenchmarkAdd_NNY_V32a3(t *testing.B) {
	a, b := V32a1, V32a1
	for i := 0; i < t.N; i++ {
		a.NNY_Add(b, &a)
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
}
