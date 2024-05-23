package main

import (
	"fmt"
	"math"
	"math/bits"
)

var numbers = []int64{
	math.MinInt64,
	-0xFFFFFFF,
	-16854351,
	-100,
	100,
	684312516,
	0xFFFFFFF,
	0xFFFFF000,
	0x7FFFFFFF_FFFFFFFF,
	math.MaxInt64,
}

func main() {
	testVariantAlpha()
	testVariantBeta()
}

func testVariantAlpha() {
	for low := int64(1); low < 63; low++ {
		for high := low+1; high < 63; high++ {
			constant := 1 | int64(1)<<low | int64(1) << high

			if !isMultiplyShiftDuplet64(constant) {
				fmt.Println("INVALID GENERATOR", constant)
				continue
			}

			for _, num := range numbers {
				exp := num * constant

				part := uint64(constant - 1)
				got := multiplyVariantAlpha(num,
					byte(bits.TrailingZeros64(part)),
					byte(63 - bits.LeadingZeros64(part)),
				)

				if exp != got {
					fmt.Printf("wrong result %v * %v = %v, but got %v\n", constant, num, exp, got)
					fmt.Printf("    con: %064b\n", constant)
					fmt.Printf("    gen: %v %v\n", high, low)
					fmt.Printf("    bit: %v %v\n", 63-bits.LeadingZeros64(part), bits.TrailingZeros64(part))

					fmt.Println()
					fmt.Printf("    %[1]v %[1]b\n", num<<low)
					fmt.Printf("    %[1]v %[1]b\n", num + num<<low)
					fmt.Printf("    %[1]v %[1]b\n", (num + num<<low) << (high-low))
					fmt.Printf("    %[1]v %[1]b\n", num + (num + num<<low) << (high-low))
				}
			}
		}
	}
}

func testVariantBeta() {
	for low := int64(1); low < 63; low++ {
		for mid := low+1; low+mid < 63; mid++ {
			high := low + mid
			constant := 1 | int64(1)<<low | int64(1) << mid | int64(1) << high

			part := uint64(constant - 1)

			if !isMultiplyShiftTriplet64(constant) {
				fmt.Printf("INVALID GENERATOR %25[1]d  %[1]b\n", constant)
				fmt.Printf("    con: %064b\n", constant)
				fmt.Printf("    gen: h%v l%v m%v\n", high, low, mid)
				fmt.Printf("    bit: h%v l%v\n", 63-bits.LeadingZeros64(part), bits.TrailingZeros64(part))
				continue
			}

			for _, num := range numbers {
				exp := num * constant

				got := multiplyVariantBeta(num,
					byte(bits.TrailingZeros64(part)),
					byte(63 - bits.LeadingZeros64(part)),
				)

				if exp != got {
					fmt.Printf("wrong result %v * %v = %v, but got %v\n", constant, num, exp, got)
					fmt.Printf("    con: %064b\n", constant)
					fmt.Printf("    gen: %v %v\n", mid, low)
					fmt.Printf("    bit: %v %v\n", 63-bits.LeadingZeros64(part), bits.TrailingZeros64(part))
				}
			}
		}
	}
}

func multiplyVariantAlpha(x int64, low, high byte) int64 {
	return x + (x + x<<(high-low)) << low
}

func multiplyVariantBeta(x int64, low, high byte) int64 {
	mid := high - low
	x = x + x<<low
	return x + x<<mid
}

// isMultiplyShiftDuplet checks for c.
// x * c; where c == 1 + 2^M + 2^(M+N), M != N,
// this can be reduced to x + (x + x<<M)<<N.
// e.g. 7, 11, 13, 19, 21, 25, 35, 37, 41, 49, 67, 69, 73, 81, 97, 131, 133, 137, 145, 161, 193...
func isMultiplyShiftDuplet64(c int64) bool {
	part := uint64(c-1)
	return c >= 1 && part&1 == 0 && bits.OnesCount64(part) == 2
}

// isMultiplyShiftTriplet checks for c, where
// x * c; where c == 1 + 2^N + 2^M + 2^(M+N), N < M.
// This can be reduced to x := x + x<<N; x = x + x<<M.
// e.g. 15, 23, 27, 29, 39, 43, 51, 53, 57, 71, 75, 83, 89, 99, 101, 135, 139, 147, 163, 169, 177, 195...
func isMultiplyShiftTriplet64(in int64) bool {
	part := uint64(in-1)
	isCandidate := in >= 3 && part&1 == 0 && bits.OnesCount64(part) == 3
	if !isCandidate {
		return false
	}

	// check whether (M+N) - N == M
	highbit := 63 - bits.LeadingZeros64(part)
	lowbit := bits.TrailingZeros64(part)
	midbit := highbit - lowbit
	return part&(1<<midbit) != 0
}
