// We spend 41% of Sign in inverseNTT, 16% in ntt, and 8% in nttMul.
// All the field operations already inline. SIMD would do amazing here
// but can you see other wins?

package mldsa

import (
	"simd"
	"unsafe"
)

type FieldElement8 = simd.Uint32x8
type FieldElement4 = simd.Uint32x4

func NTT8(f RingElement) NTTElement {
	m := byte(0)

	for len := 128; len >= 8; len /= 2 {
		for start := 0; start < 256; start += 2 * len {
			m++
			zeta := BroadcastUint32x8(zetas[m])

			xs, ys := f[start:start+len], f[start+len:start+len+len]
			for j := 0; j < len; j += 8 {
				y := LoadUint32x8Slice(ys[j:])
				t := FieldMontgomeryMul8(zeta, y)

				x := LoadUint32x8Slice(xs[j:])
				StoreUint32x8Slice(FieldSub8(x, t), ys[j:])
				StoreUint32x8Slice(FieldAdd8(x, t), xs[j:])
			}
		}
	}

	for len := 4; len >= 1; len /= 2 {
		for start := 0; start < 256; start += 2 * len {
			m++
			zeta := zetas[m]
			xs, ys := f[start:start+len], f[start+len:start+len+len]
			for j := 0; j < len; j++ {
				t := FieldMontgomeryMul(zeta, ys[j])
				ys[j] = FieldSub(xs[j], t)
				xs[j] = FieldAdd(xs[j], t)
			}
		}
	}

	return NTTElement(f)
}

func InverseNTT8(f NTTElement) RingElement {
	pf := (*[n]uint32)(unsafe.Pointer(&f[0]))
	m := byte(255)
	for len := 1; len < 256; len *= 2 {
		for start := 0; start < 256; start += 2 * len {
			zeta := zetas[m]
			m--
			zeta8 := simd.BroadcastUint32x8(uint32(zeta))

			f, flen := pf[start:start+len], pf[start+len:start+len+len]

			j := 0
			len8 := (len / 8) * 8
			for ; j < len8; j += 8 {
				t := simd.LoadUint32x8Slice(f[j:])
				tlen := simd.LoadUint32x8Slice(flen[j:])
				FieldAdd8(t, tlen).StoreSlice(f[j:])
				// -z * (t - flen[j]) = z * (flen[j] - t)
				FieldMontgomeryMulSub8(zeta8, tlen, t).StoreSlice(flen[j:])
			}

			// TODO: does this benefit from unroll by 4 as well?

			for ; j < len; j += 1 {
				t := FieldElement(f[j])
				f[j] = uint32(FieldAdd(t, FieldElement(flen[j])))
				// -z * (t - flen[j]) = z * (flen[j] - t)
				flen[j] = uint32(FieldMontgomeryMulSub(FieldElement(zeta), FieldElement(flen[j]), t))
			}
		}
	}

	c16382 := simd.BroadcastUint32x8(16382)
	for i := 0; i < len(f); i += 8 {
		t := simd.LoadUint32x8Slice(pf[i:])
		FieldMontgomeryMul8(t, c16382).StoreSlice(pf[i:])
	}
	return RingElement(f)
}

func NTTMul8(a, b NTTElement) (p NTTElement) {
	for i := 0; i < len(p); i += 8 {
		x := LoadUint32x8Slice(a[i:])
		y := LoadUint32x8Slice(b[i:])
		r := FieldMontgomeryMul8(x, y)
		StoreUint32x8Slice(r, p[i:])
	}
	return p
}

func FieldReduceOnce8(a simd.Uint32x8) FieldElement8 {
	qv := simd.BroadcastUint32x8(q)
	c := a.Add(qv).Sub(qv)
	cminusq := c.Sub(qv)
	return c.Min(cminusq)
}

func FieldAdd8(a, b FieldElement8) FieldElement8 {
	return FieldReduceOnce8(a.Add(b))
}

func FieldSub8(a, b FieldElement8) FieldElement8 {
	q8 := simd.BroadcastUint32x8(q)
	x := a.Sub(b).Add(q8)
	return FieldReduceOnce8(x)
}

func FieldMontgomeryMul8(a, b FieldElement8) FieldElement8 {
	even, odd := Uint32x8_MulEvenOddWiden(a, b)

	reven := fieldMontgomeryReduce4Unpacked(even)
	rodd := fieldMontgomeryReduce4Unpacked(odd)

	roddShifted := rodd.ShiftAllLeft(32)
	r := reven.Or(roddShifted).AsUint32x8()

	return FieldReduceOnce8(r)
}

func FieldMontgomeryMulSub8(a, b, c FieldElement8) FieldElement8 {
	return FieldMontgomeryMul8(a, b.Sub(c).Add(simd.BroadcastUint32x8(q)))
}

func fieldMontgomeryReduce4Unpacked(x simd.Uint64x4) simd.Uint64x4 {
	qNegInv8 := simd.BroadcastUint32x8(qNegInv)
	t := x.AsUint32x8().Mul(qNegInv8)

	// multiply and widen to 64bits, because we later need to
	// recombine the even and odd pairs.
	q8 := simd.BroadcastUint32x8(q)
	tq := t.MulEvenWiden(q8)

	// (tq + x) >> 32
	return tq.Add(x).ShiftAllRight(32)
}

// Utilities

// Uint32x8_MulEvenOddWiden multiplies even and odd elements separately, widening to uint64.
func Uint32x8_MulEvenOddWiden(a, b simd.Uint32x8) (even, odd simd.Uint64x4) {
	even = a.MulEvenWiden(b)

	ashift := a.AsUint64x4().ShiftAllRight(32).AsUint32x8()
	bshift := b.AsUint64x4().ShiftAllRight(32).AsUint32x8()
	odd = ashift.MulEvenWiden(bshift)

	return even, odd
}
