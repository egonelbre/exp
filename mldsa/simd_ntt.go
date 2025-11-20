package mldsa

import (
	"simd"
	"unsafe"
)

var (
	qNegInv32Vec simd.Uint32x8
	q32Vec       simd.Uint32x8
	shift32Vec   simd.Uint64x4
)

func init() {
	var qNegInvArr [8]uint32
	var qArr [8]uint32
	for i := 0; i < 8; i++ {
		qNegInvArr[i] = uint32(qNegInv)
		qArr[i] = uint32(q)
	}

	qNegInv32Vec = simd.LoadUint32x8(&qNegInvArr)
	q32Vec = simd.LoadUint32x8(&qArr)

	var shiftArr [4]uint64
	for i := 0; i < 4; i++ {
		shiftArr[i] = 32
	}
	shift32Vec = simd.LoadUint64x4(&shiftArr)
}

// reduce returns (x + t*q) >> 32, where t = x*qNegInv mod 2^32.
// The result is < 2q.
func reduce(x simd.Uint64x4) simd.Uint64x4 {
	// t = x * qNegInv mod 2^32
	// We perform this using 32-bit multiplication on the even elements (low 32 bits of x).
	x32 := x.AsUint32x8()
	tVec := x32.Mul(qNegInv32Vec)

	// t * q
	// We use MulEvenWiden to multiply t (in even positions) by q (in even positions) to get 64-bit result.
	tq := tVec.MulEvenWiden(q32Vec)

	// (x + t*q)
	sum := x.Add(tq)

	// >> 32
	// ShiftRight(shift32Vec) shifts each element right by 32 bits.
	return sum.ShiftRight(shift32Vec)
}

// vecMontgomeryMul returns a * b * R^-1 mod q.
func vecMontgomeryMul(a, b simd.Uint32x8) simd.Uint32x8 {
	// even = a_even * b_even
	even := a.MulEvenWiden(b)

	// odd = a_odd * b_odd
	// PermuteConstantGrouped(0xB1) swaps pairs in each 128-bit lane.
	// 0xB1 = 10 11 00 01.
	// [0, 1, 2, 3] -> [1, 0, 3, 2]
	aOdd := a.PermuteConstantGrouped(0xB1)
	bOdd := b.PermuteConstantGrouped(0xB1)

	odd := aOdd.MulEvenWiden(bOdd)

	// reduce
	resEven := reduce(even)
	resOdd := reduce(odd)

	// Combine
	// resEven has results in low 32 bits of each 64-bit lane.
	// [r0, 0, r2, 0, r4, 0, r6, 0] (viewed as 32-bit)
	// resOdd has results in low 32 bits.
	// [r1, 0, r3, 0, r5, 0, r7, 0]

	resEven32 := resEven.AsUint32x8()
	resOdd32 := resOdd.AsUint32x8()

	// We want resOdd at 1, 3, 5, 7.
	// Currently at 0, 2, 4, 6.
	// Swap pairs to move them to 1, 3, 5, 7.
	// [r1, 0, r3, 0] -> [0, r1, 0, r3]
	resOddShifted := resOdd32.PermuteConstantGrouped(0xB1)

	return resEven32.Or(resOddShifted)
}

// vecFieldAdd returns a + b mod q.
func vecFieldAdd(a, b simd.Uint32x8) simd.Uint32x8 {
	// c = a + b
	c := a.Add(b)
	// if c >= q, return c - q, else c.
	// c - q
	cMinusQ := c.Sub(q32Vec)
	// Min(c, c-q) will pick c-q if c >= q (since c-q < c),
	// and c if c < q (since c-q wraps to large value > c).
	return c.Min(cMinusQ)
}

// vecFieldSub returns a - b mod q.
func vecFieldSub(a, b simd.Uint32x8) simd.Uint32x8 {
	// c = a - b + q
	// We can do a + q - b to avoid underflow before adding q?
	// a + q might overflow 32-bit? q is small (8380417), a < q. a+q < 2^32. Safe.
	c := a.Add(q32Vec).Sub(b)

	// if c >= q, return c - q, else c.
	cMinusQ := c.Sub(q32Vec)
	return c.Min(cMinusQ)
}

// NTT_AVX2 computes the NTT of f in place.
func NTT_AVX2(f *NTTElement) {
	m := 0
	// Layers with len >= 8
	for len := 128; len >= 8; len /= 2 {
		for start := 0; start < 256; start += 2 * len {
			m++
			zeta := zetas[m]
			// Broadcast zeta
			var zetaArr [8]uint32
			for k := 0; k < 8; k++ {
				zetaArr[k] = uint32(zeta)
			}
			zetaVec := simd.LoadUint32x8(&zetaArr)

			for j := 0; j < len; j += 8 {
				idx1 := start + j
				idx2 := start + len + j

				v1 := simd.LoadUint32x8((*[8]uint32)(unsafe.Pointer(&f[idx1])))
				v2 := simd.LoadUint32x8((*[8]uint32)(unsafe.Pointer(&f[idx2])))

				t := vecMontgomeryMul(zetaVec, v2)

				v2New := vecFieldSub(v1, t)
				v1New := vecFieldAdd(v1, t)

				v1New.Store((*[8]uint32)(unsafe.Pointer(&f[idx1])))
				v2New.Store((*[8]uint32)(unsafe.Pointer(&f[idx2])))
			}
		}
	}

	// Scalar fallback for len < 8
	for len := 4; len >= 1; len /= 2 {
		for start := 0; start < 256; start += 2 * len {
			m++
			zeta := zetas[m]
			for j := 0; j < len; j++ {
				t := FieldMontgomeryMul(zeta, f[start+len+j])
				f[start+len+j] = FieldSub(f[start+j], t)
				f[start+j] = FieldAdd(f[start+j], t)
			}
		}
	}
}

// InverseNTT_AVX2 computes the inverse NTT of f in place.
func InverseNTT_AVX2(f *NTTElement) {
	m := 255

	// Scalar fallback for len < 8
	for len := 1; len < 8; len *= 2 {
		for start := 0; start < 256; start += 2 * len {
			zeta := zetas[m]
			m--
			for j := 0; j < len; j++ {
				t := f[start+j]
				f[start+j] = FieldAdd(t, f[start+len+j])
				f[start+len+j] = FieldMontgomeryMulSub(zeta, f[start+len+j], t)
			}
		}
	}

	// Layers with len >= 8
	for len := 8; len < 256; len *= 2 {
		for start := 0; start < 256; start += 2 * len {
			zeta := zetas[m]
			m--

			var zetaArr [8]uint32
			for k := 0; k < 8; k++ {
				zetaArr[k] = uint32(zeta)
			}
			zetaVec := simd.LoadUint32x8(&zetaArr)

			for j := 0; j < len; j += 8 {
				idx1 := start + j
				idx2 := start + len + j

				v1 := simd.LoadUint32x8((*[8]uint32)(unsafe.Pointer(&f[idx1])))
				v2 := simd.LoadUint32x8((*[8]uint32)(unsafe.Pointer(&f[idx2])))

				// t = f[j] (v1)
				// f[j] = t + flen[j] => v1New = v1 + v2
				v1New := vecFieldAdd(v1, v2)

				// flen[j] = zeta * (flen[j] - t) => v2New = zeta * (v2 - v1)
				diff := vecFieldSub(v2, v1)
				v2New := vecMontgomeryMul(zetaVec, diff)

				v1New.Store((*[8]uint32)(unsafe.Pointer(&f[idx1])))
				v2New.Store((*[8]uint32)(unsafe.Pointer(&f[idx2])))
			}
		}
	}

	// Scaling by 256^-1
	var scaleArr [8]uint32
	for k := 0; k < 8; k++ {
		scaleArr[k] = 16382 // 256^-1 mod q * R
	}
	scaleVec := simd.LoadUint32x8(&scaleArr)

	for i := 0; i < 256; i += 8 {
		v := simd.LoadUint32x8((*[8]uint32)(unsafe.Pointer(&f[i])))
		v = vecMontgomeryMul(v, scaleVec)
		v.Store((*[8]uint32)(unsafe.Pointer(&f[i])))
	}
}

// NTTMul_AVX2 multiplies two nttElements.
func NTTMul_AVX2(a, b *NTTElement) NTTElement {
	var p NTTElement
	// Process 8 elements at a time.
	// 256 / 8 = 32 iterations.
	for i := 0; i < 256; i += 8 {
		aVec := simd.LoadUint32x8((*[8]uint32)(unsafe.Pointer(&a[i])))
		bVec := simd.LoadUint32x8((*[8]uint32)(unsafe.Pointer(&b[i])))
		pVec := vecMontgomeryMul(aVec, bVec)
		pVec.Store((*[8]uint32)(unsafe.Pointer(&p[i])))
	}
	return p
}
