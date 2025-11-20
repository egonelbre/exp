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
	pf := (*[n]uint32)(unsafe.Pointer(&f[0]))

	m := byte(0)
	for len := 128; len >= 1; len /= 2 {
		for start := 0; start < 256; start += 2 * len {
			m++
			zeta := zetas[m]
			zeta8 := simd.BroadcastUint32x8(uint32(zeta))

			f, flen := pf[start:start+len], pf[start+len:start+len+len]

			j := 0
			len8 := (len / 8) * 8
			for ; j < len8; j += 8 {
				tlen := simd.LoadUint32x8Slice(flen[j:])
				t := FieldMontgomeryMul8(zeta8, tlen)
				fj := simd.LoadUint32x8Slice(f[j:])
				FieldSub8(fj, t).StoreSlice(flen[j:])
				FieldAdd8(fj, t).StoreSlice(f[j:])
			}

			// TODO: does this benefit from unroll by 4 as well?

			for ; j < len; j++ {
				t := FieldMontgomeryMul(zeta, FieldElement(flen[j]))
				flen[j] = uint32(FieldSub(FieldElement(f[j]), t))
				f[j] = uint32(FieldAdd(FieldElement(f[j]), t))
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
	ap := (*[n]uint32)(unsafe.Pointer(&a[0]))
	bp := (*[n]uint32)(unsafe.Pointer(&b[0]))
	pp := (*[n]uint32)(unsafe.Pointer(&p[0]))
	for i := 0; i < len(p); i += 8 {
		x := simd.LoadUint32x8Slice(ap[i:])
		y := simd.LoadUint32x8Slice(bp[i:])
		r := FieldMontgomeryMul8(x, y)
		r.StoreSlice(pp[i:])
	}
	return p
}

func FieldReduceOnce4(a simd.Uint32x4) FieldElement4 {
	qv := simd.BroadcastUint32x4(q)
	c := a.Add(qv).Sub(qv)
	cminusq := c.Sub(qv)
	return c.Min(cminusq)
}

func FieldReduceOnce8(a simd.Uint32x8) FieldElement8 {
	qv := simd.BroadcastUint32x8(q)
	c := a.Add(qv).Sub(qv)
	cminusq := c.Sub(qv)
	return c.Min(cminusq)
}

func FieldAdd4(a, b FieldElement4) FieldElement4 {
	return FieldReduceOnce4(a.Add(b))
}

func FieldAdd8(a, b FieldElement8) FieldElement8 {
	return FieldReduceOnce8(a.Add(b))
}

func FieldSub4(a, b FieldElement4) FieldElement4 {
	q4 := simd.BroadcastUint32x4(q)
	x := a.Sub(b).Add(q4)
	return FieldReduceOnce4(x)
}

func FieldSub8(a, b FieldElement8) FieldElement8 {
	q8 := simd.BroadcastUint32x8(q)
	x := a.Sub(b).Add(q8)
	return FieldReduceOnce8(x)
}

func FieldMontgomeryMul4(a, b FieldElement4) FieldElement4 {
	aw, bw := Uint32x4_Spread(a), Uint32x4_Spread(b)
	r := aw.MulEvenWiden(bw)
	return FieldMontgomeryReduce4(r)
}

func FieldMontgomeryMul8(a, b FieldElement8) FieldElement8 {
	even, odd := Uint32x8_MulEvenOddWiden(a, b)

	reven := fieldMontgomeryReduce4Unpacked(even)
	rodd := fieldMontgomeryReduce4Unpacked(odd)

	roddShifted := rodd.ShiftAllLeft(32)
	r := reven.Or(roddShifted).AsUint32x8()

	return FieldReduceOnce8(r)
}

func FieldMontgomeryMulSub4(a, b, c FieldElement4) FieldElement4 {
	return FieldMontgomeryMul4(a, b.Sub(c).Add(simd.BroadcastUint32x4(q)))
}

func FieldMontgomeryMulSub8(a, b, c FieldElement8) FieldElement8 {
	return FieldMontgomeryMul8(a, b.Sub(c).Add(simd.BroadcastUint32x8(q)))
}

func FieldMontgomeryReduce4(x simd.Uint64x4) FieldElement4 {
	return FieldReduceOnce4(Uint64x4_ConvertToUint32(fieldMontgomeryReduce4Unpacked(x)))
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

// Uint32x4_Spread spreads 4 elements into 8, interleaving with zeros.
func Uint32x4_Spread(a simd.Uint32x4) simd.Uint32x8 {
	zeros := simd.Uint32x4{}
	var r simd.Uint32x8
	r = r.SetLo(zeros.InterleaveLo(a))
	r = r.SetHi(zeros.InterleaveHi(a))
	return r
}

// Uint64x4_ConvertToUint32 converts 4 uint64s to 4 uint32s by keeping the low 32 bits.
func Uint64x4_ConvertToUint32(in simd.Uint64x4) simd.Uint32x4 {
	gatherLowIndices := simd.LoadUint32x8(&[8]uint32{0, 2, 4, 6, 0, 0, 0, 0})
	permuted := in.AsUint32x8().Permute(gatherLowIndices)
	return permuted.GetLo()
}

// Uint32x8_MulEvenOddWiden multiplies even and odd elements separately, widening to uint64.
func Uint32x8_MulEvenOddWiden(a, b simd.Uint32x8) (even, odd simd.Uint64x4) {
	even = a.MulEvenWiden(b)

	//ashift := a.AsUint64x4().ShiftAllRight(32).AsUint32x8()
	//bshift := b.AsUint64x4().ShiftAllRight(32).AsUint32x8()
	ashift := a.PermuteConstantGrouped(0xB1)
	bshift := b.PermuteConstantGrouped(0xB1)
	odd = ashift.MulEvenWiden(bshift)
	return even, odd
}

// Uint32x4_Interleave interleaves two Uint32x4 vectors into one Uint32x8.
func Uint32x4_Interleave(even, odd simd.Uint32x4) (r simd.Uint32x8) {
	r = r.SetLo(even.InterleaveLo(odd))
	r = r.SetHi(even.InterleaveHi(odd))
	return r
}
