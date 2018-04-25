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

var (
	fact = [base]int{
		39916800,
		3628800,
		362880,
		40320,
		5040,
		720,
		120,
		24,
		6,
		2,
		1,
		1,
	}

	masks [base]uint64
)

func init() {
	/*
		fact[0] = 1
		for i := 1; i < base; i++ {
			fact[i] = fact[i-1] * i
		}
		for i := 0; i < base/2; i++ {
			fact[i], fact[base-i-1] = fact[base-i-1], fact[i]
		}
	*/

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
	var state = [base]byte{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
	}
	var inverse = [base]byte{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
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

	total := 0
	for i, v := range perm {
		total += int(v) * fact[i]
	}

	return total
}

func CodeShuffleUnroll(perm [base]byte) int {
	var state = [base]byte{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
	}
	var inverse = [base]byte{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
	}

	total := 0

	j0 := inverse[perm[0]]
	inverse[state[0]], state[j0] = j0, state[0]
	total += int(j0-0) * 39916800

	j1 := inverse[perm[1]]
	inverse[state[1]], state[j1] = j1, state[1]
	total += int(j1-1) * 3628800

	j2 := inverse[perm[2]]
	inverse[state[2]], state[j2] = j2, state[2]
	total += int(j2-2) * 362880

	j3 := inverse[perm[3]]
	inverse[state[3]], state[j3] = j3, state[3]
	total += int(j3-3) * 40320

	j4 := inverse[perm[4]]
	inverse[state[4]], state[j4] = j4, state[4]
	total += int(j4-4) * 5040

	j5 := inverse[perm[5]]
	inverse[state[5]], state[j5] = j5, state[5]
	total += int(j5-5) * 720

	j6 := inverse[perm[6]]
	inverse[state[6]], state[j6] = j6, state[6]
	total += int(j6-6) * 120

	j7 := inverse[perm[7]]
	inverse[state[7]], state[j7] = j7, state[7]
	total += int(j7-7) * 24

	j8 := inverse[perm[8]]
	inverse[state[8]], state[j8] = j8, state[8]
	total += int(j8-8) * 6

	j9 := inverse[perm[9]]
	inverse[state[9]], state[j9] = j9, state[9]
	total += int(j9-9) * 2

	j10 := inverse[perm[10]]
	inverse[state[10]], state[j10] = j10, state[10]
	total += int(j10-10) * 1

	return total
}

func DecodeShuffle(v int) [base]byte {
	var state = [base]byte{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
	}

	var perm [base]byte
	for i := range perm {
		perm[i] = byte(v / fact[i])
		v = v % fact[i]
	}

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
