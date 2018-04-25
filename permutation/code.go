package permutation

import (
	"math/bits"
)

const (
	base       = 12
	nybblemask = 1<<4 - 1
	fullmask   = ^uint64(0)
	maskstarts = 1<<(4*0) | 1<<(4*1) | 1<<(4*2) | 1<<(4*3) | 1<<(4*4) | 1<<(4*5) | 1<<(4*6) | 1<<(4*7) | 1<<(4*8) | 1<<(4*9) | 1<<(4*10) | 1<<(4*11)
)

var fact [base]int
var masks [base]uint64

func init() {
	fact[0] = 1
	for i := 1; i < base; i++ {
		fact[i] = fact[i-1] * i
	}
	for i := 0; i < base/2; i++ {
		fact[i], fact[base-i-1] = fact[base-i-1], fact[i]
	}

	for min := 0; min < base; min++ {
		v := min
		for i := 0; i < base; i++ {
			v = v<<4 | min
		}
		masks[min] = ^uint64(v)
	}
}

func Code(perm [base]byte) int { return CodeCopy(perm) }

func CodeCopy(perm [base]byte) int {
	r := 0
	for min := byte(0); min < base-1; min++ {
		z := 0
		n := base - min
		for i, v := range perm[:n] {
			if v == min {
				z = i
				break
			}
		}
		copy(perm[z:], perm[z+1:n])
		r += z * fact[min]
	}
	return r
}

func CodeCopyBit(perm [base]byte) int {
	return codebit(uint64(PackNybbleUnroll(perm)))
}

func CodeCopyBitNonstandard(perm [base]byte) int {
	return codebit(PackNybbleUnroll3(perm))
}

func codebit(perm uint64) int {
	r := 0
	for min, mask := range masks {
		filtered := perm ^ mask
		filtered &= filtered >> 2
		filtered &= filtered >> 1
		filtered &= maskstarts
		z := byte(bits.TrailingZeros64(filtered))
		upper := (perm >> (z + 4)) << z
		lower := perm &^ (fullmask << z)
		perm = upper | lower
		r += int(z/4) * fact[min]
	}
	return r
}

func CodeCount(perm [base]byte) int {
	r := 0
	for min := byte(0); min < base-1; min++ {
		z := 0
		for _, v := range perm {
			if v > min {
				z++
			} else if v == min {
				break
			}
		}
		r += z * fact[min]
	}
	return r
}

func CodeTable(perm [base]byte) int {
	var index [base]byte
	for i, v := range perm {
		index[v] = byte(i)
	}

	r := 0
	for min := byte(0); min < base-1; min++ {
		z := 0
		for _, v := range perm[:index[min]] {
			if v > min {
				z++
			}
		}
		r += z * fact[min]
	}
	return r
}

func CodeTable2(perm [base]byte) int {
	var index [base]byte
	for i, v := range perm {
		index[v] = byte(i)
	}

	r := 0
	for min := byte(0); min < base-1; min++ {
		z := index[min]
		for i, pos := range index {
			if pos > z {
				index[i] = pos - 1
			}
		}
		r += int(z) * fact[min]
	}
	return r
}

func CodeShuffle(perm [base]byte) int {
	// #define SQ_COUNT 25
	// #define PC_COUNT 12
	var state = [25]byte{
		0, 1, 2, 3, 4,
		5, 6, 7, 8, 9,
		10, 11, 12, 13, 14,
		15, 16, 17, 18, 19,
		20, 21, 22, 23, 24,
	}
	var inverse = [25]byte{
		0, 1, 2, 3, 4,
		5, 6, 7, 8, 9,
		10, 11, 12, 13, 14,
		15, 16, 17, 18, 19,
		20, 21, 22, 23, 24,
	}
	for i := range perm {
		/*
		 * this cryptic looking code is an optimized version of
		 * this pseudocode:
		 *
		 *     j = inverse[perm[i]]
		 *     swap entries state[i] and state[j] in inverse
		 *     swap entries i and j in state
		 *     perm[i] = j - i
		 *
		 * two optimizations are performed:
		 *  - inverse[state[k]] == k and state[inverse[k]] == k
		 *  - after iteration i, inverse[perm[i]] and state[i]
		 *    are never read again, so we can omit assigning to
		 *    them.
		 */
		j := inverse[perm[i]]
		inverse[state[i]] = j
		state[j] = state[i]
		perm[i] = j - byte(i)
	}
	return PackNybbleUnroll(perm)
}

func DecodeShuffle(v int) [base]byte {
	var state = [25]byte{
		0, 1, 2, 3, 4,
		5, 6, 7, 8, 9,
		10, 11, 12, 13, 14,
		15, 16, 17, 18, 19,
		20, 21, 22, 23, 24,
	}
	perm := UnpackNybbleUnroll(v)
	for i := byte(0); i < base; i++ {
		j := i + perm[i]
		perm[i] = state[j]
		state[j] = state[i]
	}
	return perm
}

func PackNybble(perm [base]byte) int {
	r := 0
	for _, v := range perm {
		r = r<<4 + int(v)
	}
	return r
}

func PackTight(perm [base]byte) int {
	r := 0
	for _, v := range perm {
		r = r*12 + int(v)
	}
	return r
}

func PackNybbleUnroll(perm [base]byte) int {
	return int(perm[0]) |
		int(perm[1])<<(4*1) |
		int(perm[2])<<(4*2) |
		int(perm[3])<<(4*3) |
		int(perm[4])<<(4*4) |
		int(perm[5])<<(4*5) |
		int(perm[6])<<(4*6) |
		int(perm[7])<<(4*7) |
		int(perm[8])<<(4*8) |
		int(perm[9])<<(4*9) |
		int(perm[10])<<(4*10) |
		int(perm[11])<<(4*11)
}

func UnpackNybbleUnroll(v int) [base]byte {
	const mask = 1<<4 - 1
	return [base]byte{
		byte((v >> (4 * 0)) & mask),
		byte((v >> (4 * 1)) & mask),
		byte((v >> (4 * 2)) & mask),
		byte((v >> (4 * 3)) & mask),
		byte((v >> (4 * 4)) & mask),
		byte((v >> (4 * 5)) & mask),
		byte((v >> (4 * 6)) & mask),
		byte((v >> (4 * 7)) & mask),
		byte((v >> (4 * 8)) & mask),
		byte((v >> (4 * 9)) & mask),
		byte((v >> (4 * 10)) & mask),
		byte((v >> (4 * 11)) & mask),
	}
}

func PackNybbleUnrollReverse(perm [base]byte) int {
	return int(perm[0]<<(4*11)) |
		int(perm[1])<<(4*10) |
		int(perm[2])<<(4*9) |
		int(perm[3])<<(4*8) |
		int(perm[4])<<(4*7) |
		int(perm[5])<<(4*6) |
		int(perm[6])<<(4*5) |
		int(perm[7])<<(4*4) |
		int(perm[8])<<(4*3) |
		int(perm[9])<<(4*2) |
		int(perm[10])<<(4*1) |
		int(perm[11])<<(4*0)
}

func PackNybbleUnroll2(perm [base]byte) int {
	i0 := int(perm[0]) | int(perm[1])<<8 | int(perm[2])<<16 | int(perm[3])<<24
	i1 := int(perm[4]) | int(perm[5])<<8 | int(perm[6])<<16 | int(perm[7])<<24
	i2 := int(perm[8]) | int(perm[9])<<8 | int(perm[10])<<16 | int(perm[11])<<24
	return i0 | i1<<4 | i2<<32
}

func PackNybbleUnroll3(perm [base]byte) uint64 {
	i0 := uint64(perm[0]) | uint64(perm[1])<<8 | uint64(perm[2])<<16 | uint64(perm[3])<<24
	i1 := uint64(perm[4]) | uint64(perm[5])<<8 | uint64(perm[6])<<16 | uint64(perm[7])<<24
	i2 := uint64(perm[8]) | uint64(perm[9])<<8
	i3 := uint64(perm[10]) | uint64(perm[11])<<8
	return i0 | i1<<4 | i2<<32 | i3<<36
}

const (
	b12_0  = 1
	b12_1  = b12_0 * 12
	b12_2  = b12_1 * 12
	b12_3  = b12_2 * 12
	b12_4  = b12_3 * 12
	b12_5  = b12_4 * 12
	b12_6  = b12_5 * 12
	b12_7  = b12_6 * 12
	b12_8  = b12_7 * 12
	b12_9  = b12_8 * 12
	b12_10 = b12_9 * 12
	b12_11 = b12_10 * 12
)

func PackTightUnroll(perm [base]byte) int {
	return int(perm[0])*b12_0 |
		int(perm[1])*b12_1 |
		int(perm[2])*b12_2 |
		int(perm[3])*b12_3 |
		int(perm[4])*b12_4 |
		int(perm[5])*b12_5 |
		int(perm[6])*b12_6 |
		int(perm[7])*b12_7 |
		int(perm[8])*b12_8 |
		int(perm[9])*b12_9 |
		int(perm[10])*b12_10 |
		int(perm[11])*b12_11
}
