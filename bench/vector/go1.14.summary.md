# Shorthands

Definitions for the funcs:

```
func (a *V) PP_Add(b *V)        { a.X += b.X; a.Y += b.Y; a.Z += b.Z }
func (a *V) PN_Add(b V)         { a.X += b.X; a.Y += b.Y; a.Z += b.Z }
func (a *V) PPP_Add(b *V) *V    { return &V{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V)  NNN_Add(b V) V      { return V{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a *V) PPN_Add(b *V) V     { return V{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a *V) PNP_Add(b V) *V     { return &V{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V)  NPP_Add(b *V) *V    { return &V{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V)  NNP_Add(b V) *V     { return &V{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V)  NPN_Add(b *V) V     { return V{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a *V) PNN_Add(b V) V      { return V{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }

func (a *V) PPZ_Add(b *V, r *V) { *r = V{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V)  NPZ_Add(b *V, r *V) { *r = V{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a *V) PNZ_Add(b V, r *V)  { *r = V{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }
func (a V)  NNZ_Add(b V, r *V)  { *r = V{a.X + b.X, a.Y + b.Y, a.Z + b.Z} }

func (a *V) PPY_Add(b *V, r *V) { r.X = a.X + b.X; r.Y = a.Y + b.Y; r.Z = a.Z + b.Z }
func (a V)  NPY_Add(b *V, r *V) { r.X = a.X + b.X; r.Y = a.Y + b.Y; r.Z = a.Z + b.Z }
func (a *V) PNY_Add(b V, r *V)  { r.X = a.X + b.X; r.Y = a.Y + b.Y; r.Z = a.Z + b.Z }
func (a V)  NNY_Add(b V, r *V)  { r.X = a.X + b.X; r.Y = a.Y + b.Y; r.Z = a.Z + b.Z }
```
# float32

## struct

```
Add_PP_V32s3-32   2.67ns ± 0%
Add_PN_V32s3-32   2.65ns ± 0%
Add_PPP_V32s3-32  2.65ns ± 0%
Add_NNN_V32s3-32  0.72ns ± 0%
Add_PPN_V32s3-32  2.66ns ± 1%
Add_PNP_V32s3-32  2.65ns ± 0%
Add_NPP_V32s3-32  0.73ns ± 1%
Add_NNP_V32s3-32  0.72ns ± 0%
Add_NPN_V32s3-32  0.73ns ± 1%
Add_PNN_V32s3-32  2.64ns ± 0%

Add_PPZ_V32s3-32  2.64ns ± 0%
Add_NPZ_V32s3-32  2.67ns ± 0%
Add_PNZ_V32s3-32  2.65ns ± 0%
Add_NNZ_V32s3-32  2.65ns ± 1%

Add_PPY_V32s3-32  2.64ns ± 0%
Add_NPY_V32s3-32  2.65ns ± 0%
Add_PNY_V32s3-32  2.64ns ± 0%
Add_NNY_V32s3-32  2.64ns ± 0%
```

## array

```
Add_PP_V32a3-32   2.65ns ± 0%
Add_PN_V32a3-32   2.71ns ± 0%
Add_PPP_V32a3-32  2.65ns ± 0%
Add_NNN_V32a3-32  10.7ns ± 1%
Add_PPN_V32a3-32  8.16ns ± 0%
Add_PNP_V32a3-32  2.74ns ± 0%
Add_NPP_V32a3-32  7.30ns ± 1%
Add_NNP_V32a3-32  7.24ns ± 0%
Add_NPN_V32a3-32  9.88ns ± 0%
Add_PNN_V32a3-32  8.56ns ± 0%

Add_PPZ_V32a3-32  2.66ns ± 1%
Add_NPZ_V32a3-32  7.32ns ± 0%
Add_PNZ_V32a3-32  2.72ns ± 0%
Add_NNZ_V32a3-32  7.33ns ± 1%

Add_PPY_V32a3-32  2.66ns ± 0%
Add_NPY_V32a3-32  7.36ns ± 1%
Add_PNY_V32a3-32  2.75ns ± 0%
Add_NNY_V32a3-32  8.48ns ± 0%
```

# float64

## struct

```
Add_PP_V64s3-32   2.65ns ± 0%
Add_PN_V64s3-32   2.65ns ± 0%
Add_PPP_V64s3-32  2.65ns ± 0%
Add_NNN_V64s3-32  0.73ns ± 0%
Add_PPN_V64s3-32  2.66ns ± 0%
Add_PNP_V64s3-32  2.65ns ± 0%
Add_NPP_V64s3-32  0.73ns ± 0%
Add_NNP_V64s3-32  0.72ns ± 0%
Add_NPN_V64s3-32  0.73ns ± 0%
Add_PNN_V64s3-32  2.65ns ± 0%

Add_PPZ_V64s3-32  2.65ns ± 0%
Add_NPZ_V64s3-32  2.65ns ± 0%
Add_PNZ_V64s3-32  2.65ns ± 0%
Add_NNZ_V64s3-32  2.66ns ± 0%

Add_PPY_V64s3-32  2.66ns ± 0%
Add_NPY_V64s3-32  2.70ns ± 0%
Add_PNY_V64s3-32  2.66ns ± 0%
Add_NNY_V64s3-32  2.66ns ± 0%
```

## array

```
Add_PP_V64a3-32   2.66ns ± 0%
Add_PN_V64a3-32   2.67ns ± 0%
Add_PPP_V64a3-32  2.65ns ± 0%
Add_NNN_V64a3-32  9.91ns ± 0%
Add_PPN_V64a3-32  8.67ns ± 0%
Add_PNP_V64a3-32  2.71ns ± 0%
Add_NPP_V64a3-32  9.14ns ± 0%
Add_NNP_V64a3-32  9.15ns ± 0%
Add_NPN_V64a3-32  9.84ns ± 0%
Add_PNN_V64a3-32  8.23ns ± 0%

Add_PPZ_V64a3-32  2.66ns ± 1%
Add_NPZ_V64a3-32  9.18ns ± 0%
Add_PNZ_V64a3-32  2.69ns ± 0%
Add_NNZ_V64a3-32  9.17ns ± 0%

Add_PPY_V64a3-32  2.66ns ± 0%
Add_NPY_V64a3-32  8.69ns ± 0%
Add_PNY_V64a3-32  2.68ns ± 0%
Add_NNY_V64a3-32  7.30ns ± 0%
```