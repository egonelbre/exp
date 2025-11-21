// We spend 41% of Sign in inverseNTT, 16% in ntt, and 8% in nttMul.
// All the field operations already inline. SIMD would do amazing here
// but can you see other wins?

// Code based on https://godbolt.org/z/YPK5PbnaP

package mldsa

import (
	"math/bits"
)

const n = 256

// RingElement is a polynomial, an element of R_q.
type RingElement [n]FieldElement

// NTTElement is an NTT representation, an element of T_q.
type NTTElement [n]FieldElement

// NTT maps a ringElement to its nttElement representation.
//
// It implements NTT, according to FIPS 203, Algorithm 9.
func NTT(f RingElement) NTTElement {
	m := byte(0)
	for len := 128; len >= 1; len /= 2 {
		for start := 0; start < 256; start += 2 * len {
			m++
			zeta := zetas[m]
			// Bounds check elimination hint.
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

// InverseNTT maps a nttElement back to the ringElement it represents.
//
// It implements NTT⁻¹, according to FIPS 203, Algorithm 10.
func InverseNTT(f NTTElement) RingElement {
	m := byte(255)
	for len := 1; len < 256; len *= 2 {
		for start := 0; start < 256; start += 2 * len {
			zeta := zetas[m]
			m--
			// Bounds check elimination hint.
			xs, ys := f[start:start+len], f[start+len:start+len+len]
			for j := 0; j < len; j++ {
				t := xs[j]
				xs[j] = FieldAdd(t, ys[j])
				// -z * (t - ys[j]) = z * (ys[j] - t)
				ys[j] = FieldMontgomeryMulSub(zeta, ys[j], t)
			}
		}
	}
	for i := range f {
		f[i] = FieldMontgomeryMul(f[i], 16382) // 16382 = 256⁻¹ * R mod q
	}
	return RingElement(f)
}

// NTTMul multiplies two nttElements.
func NTTMul(a, b NTTElement) (p NTTElement) {
	for i := range p {
		p[i] = FieldMontgomeryMul(a[i], b[i])
	}
	return p
}

// zetas are the values ζ^BitRev₈(k) mod q for each index k, converted to the
// Montgomery domain.
var zetas = [256]FieldElement{4193792, 25847, 5771523, 7861508, 237124, 7602457, 7504169, 466468, 1826347, 2353451, 8021166, 6288512, 3119733, 5495562, 3111497, 2680103, 2725464, 1024112, 7300517, 3585928, 7830929, 7260833, 2619752, 6271868, 6262231, 4520680, 6980856, 5102745, 1757237, 8360995, 4010497, 280005, 2706023, 95776, 3077325, 3530437, 6718724, 4788269, 5842901, 3915439, 4519302, 5336701, 3574422, 5512770, 3539968, 8079950, 2348700, 7841118, 6681150, 6736599, 3505694, 4558682, 3507263, 6239768, 6779997, 3699596, 811944, 531354, 954230, 3881043, 3900724, 5823537, 2071892, 5582638, 4450022, 6851714, 4702672, 5339162, 6927966, 3475950, 2176455, 6795196, 7122806, 1939314, 4296819, 7380215, 5190273, 5223087, 4747489, 126922, 3412210, 7396998, 2147896, 2715295, 5412772, 4686924, 7969390, 5903370, 7709315, 7151892, 8357436, 7072248, 7998430, 1349076, 1852771, 6949987, 5037034, 264944, 508951, 3097992, 44288, 7280319, 904516, 3958618, 4656075, 8371839, 1653064, 5130689, 2389356, 8169440, 759969, 7063561, 189548, 4827145, 3159746, 6529015, 5971092, 8202977, 1315589, 1341330, 1285669, 6795489, 7567685, 6940675, 5361315, 4499357, 4751448, 3839961, 2091667, 3407706, 2316500, 3817976, 5037939, 2244091, 5933984, 4817955, 266997, 2434439, 7144689, 3513181, 4860065, 4621053, 7183191, 5187039, 900702, 1859098, 909542, 819034, 495491, 6767243, 8337157, 7857917, 7725090, 5257975, 2031748, 3207046, 4823422, 7855319, 7611795, 4784579, 342297, 286988, 5942594, 4108315, 3437287, 5038140, 1735879, 203044, 2842341, 2691481, 5790267, 1265009, 4055324, 1247620, 2486353, 1595974, 4613401, 1250494, 2635921, 4832145, 5386378, 1869119, 1903435, 7329447, 7047359, 1237275, 5062207, 6950192, 7929317, 1312455, 3306115, 6417775, 7100756, 1917081, 5834105, 7005614, 1500165, 777191, 2235880, 3406031, 7838005, 5548557, 6709241, 6533464, 5796124, 4656147, 594136, 4603424, 6366809, 2432395, 2454455, 8215696, 1957272, 3369112, 185531, 7173032, 5196991, 162844, 1616392, 3014001, 810149, 1652634, 4686184, 6581310, 5341501, 3523897, 3866901, 269760, 2213111, 7404533, 1717735, 472078, 7953734, 1723600, 6577327, 1910376, 6712985, 7276084, 8119771, 4546524, 5441381, 6144432, 7959518, 6094090, 183443, 7403526, 1612842, 4834730, 7826001, 3919660, 8332111, 7018208, 3937738, 1400424, 7534263, 1976782}

const (
	q       = 8380417    // 2²³ - 2¹³ + 1
	qNegInv = 4236238847 // -q⁻¹ mod R (q * qNegInv ≡ -1 mod R)
)

// FieldElement is an element n of ℤ_q in the Montgomery domain, represented as
// an integer x in [0, q) such that x ≡ n * R (mod q) where R = 2³².
type FieldElement uint32

func FieldReduceOnce(a uint32) FieldElement {
	return FieldElement(min(a, a-q))
}

// FieldReduceOnceRef reduces a value a < 2q.
func FieldReduceOnceRef(a uint32) FieldElement {
	x, b := bits.Sub64(uint64(a), uint64(q), 0)
	return FieldElement(x + b*q)
}

// FieldAdd returns a + b mod q.
func FieldAdd(a, b FieldElement) FieldElement {
	x := uint32(a + b)
	return FieldReduceOnce(x)
}

// FieldSub returns a - b mod q.
func FieldSub(a, b FieldElement) FieldElement {
	x := uint32(a - b + q)
	return FieldReduceOnce(x)
}

// FieldMontgomeryMul returns a * b * R⁻¹ mod q.
func FieldMontgomeryMul(a, b FieldElement) FieldElement {
	x := uint64(a) * uint64(b)
	return FieldMontgomeryReduce(x)
}

// FieldMontgomeryReduce returns x * R⁻¹ mod q for x < q * R.
func FieldMontgomeryReduce(x uint64) FieldElement {
	t := uint32(x) * qNegInv
	u := (x + uint64(t)*q) >> 32
	return FieldReduceOnce(uint32(u))
}

// FieldMontgomeryMulSub returns a * (b - c). This operation is fused to save a
// fieldReduceOnce after the subtraction.
func FieldMontgomeryMulSub(a, b, c FieldElement) FieldElement {
	x := uint64(a) * uint64(b-c+q)
	return FieldMontgomeryReduce(x)
}
