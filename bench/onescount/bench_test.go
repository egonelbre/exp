package onescount_test

import (
	"math/bits"
	"testing"
)

const m0 = 0x5555555555555555 // 01010101 ...
const m1 = 0x3333333333333333 // 00110011 ...
const m2 = 0x0f0f0f0f0f0f0f0f // 00001111 ...
const m3 = 0x00ff00ff00ff00ff // etc.
const m4 = 0x0000ffff0000ffff
const mz = 0xffffffffffffffff

// GoOnesCount64 returns the number of one bits ("population count") in x.
func GoOnesCount64(x uint64) int {
	// Implementation: Parallel summing of adjacent bits.
	// See "Hacker's Delight", Chap. 5: Counting Bits.
	// The following pattern shows the general approach:
	//
	//   x = x>>1&(m0&m) + x&(m0&m)
	//   x = x>>2&(m1&m) + x&(m1&m)
	//   x = x>>4&(m2&m) + x&(m2&m)
	//   x = x>>8&(m3&m) + x&(m3&m)
	//   x = x>>16&(m4&m) + x&(m4&m)
	//   x = x>>32&(m5&m) + x&(m5&m)
	//   return int(x)
	//
	// Masking (& operations) can be left away when there's no
	// danger that a field's sum will carry over into the next
	// field: Since the result cannot be > 64, 8 bits is enough
	// and we can ignore the masks for the shifts by 8 and up.
	// Per "Hacker's Delight", the first line can be simplified
	// more, but it saves at best one instruction, so we leave
	// it alone for clarity.
	const m = 1<<64 - 1
	x = x>>1&(m0&m) + x&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	x += x >> 8
	x += x >> 16
	x += x >> 32
	return int(x) & (1<<7 - 1)
}

func GoOnesCount256(x, y, z, w uint64) int {
	return GoOnesCount64(x) +
		GoOnesCount64(y) +
		GoOnesCount64(z) +
		GoOnesCount64(w)
}

func BitsOnesCount256(x, y, z, w uint64) int {
	return bits.OnesCount64(x) +
		bits.OnesCount64(y) +
		bits.OnesCount64(z) +
		bits.OnesCount64(w)
}

func OnesCount256(x, y, z, w uint64) int {
	const m = 1<<64 - 1

	x = x>>1&(m0&m) + x&(m0&m)
	y = y>>1&(m0&m) + y&(m0&m)
	z = z>>1&(m0&m) + z&(m0&m)
	w = w>>1&(m0&m) + w&(m0&m)

	x = x>>2&(m1&m) + x&(m1&m)
	y = y>>2&(m1&m) + y&(m1&m)
	z = z>>2&(m1&m) + z&(m1&m)
	w = w>>2&(m1&m) + w&(m1&m)

	x = (x>>4 + x) & (m2 & m)
	y = (y>>4 + y) & (m2 & m)
	z = (z>>4 + z) & (m2 & m)
	w = (w>>4 + w) & (m2 & m)

	q := x + y + z + w

	q += q >> 8
	q += q >> 16
	q += q >> 32

	return int(q) & (1<<9 - 1)
}

func OnesCount256Alt(x, y, z, w uint64) int {
	const m = 1<<64 - 1

	var q uint64

	x = x>>1&(m0&m) + x&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	q += x

	y = y>>1&(m0&m) + y&(m0&m)
	y = y>>2&(m1&m) + y&(m1&m)
	y = (y>>4 + y) & (m2 & m)
	q += y

	z = z>>1&(m0&m) + z&(m0&m)
	z = z>>2&(m1&m) + z&(m1&m)
	z = (z>>4 + z) & (m2 & m)
	q += z

	w = w>>1&(m0&m) + w&(m0&m)
	w = w>>2&(m1&m) + w&(m1&m)
	w = (w>>4 + w) & (m2 & m)
	q += w

	q += q >> 8
	q += q >> 16
	q += q >> 32

	return int(q) & (1<<9 - 1)
}

func BenchmarkGoOnesCount256(b *testing.B) {
	var z int
	for i := 0; i < b.N; i++ {
		k := uint64(i)
		z += GoOnesCount256(k, k, k, k)
	}
	sink(z)
}

func BenchmarkOnesCount256(b *testing.B) {
	var z int
	for i := 0; i < b.N; i++ {
		k := uint64(i)
		z += OnesCount256(k, k, k, k)
	}
	sink(z)
}

func BenchmarkOnesCount256Alt(b *testing.B) {
	var z int
	for i := 0; i < b.N; i++ {
		k := uint64(i)
		z += OnesCount256Alt(k, k, k, k)
	}
	sink(z)
}

func BenchmarkBitsOnesCount256(b *testing.B) {
	var z int
	for i := 0; i < b.N; i++ {
		k := uint64(i)
		z += BitsOnesCount256(k, k, k, k)
	}
	sink(z)
}

func TestOnesCount256(t *testing.T) {
	check := func(a, b, c, d uint64) {
		exp := BitsOnesCount256(a, b, c, d)
		got := OnesCount256(a, b, c, d)
		if exp != got {
			t.Error(a, b, c, d, "exp", exp, "got", got)
		}
	}

	check(mz, mz, mz, mz)
	check(mz, mz-1, mz-2, mz-3)
	check(mz-3, mz-2, mz-1, mz)
}

//go:noinline
func sink(v int) {}
