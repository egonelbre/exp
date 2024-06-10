package div5_test

import (
	"math/big"
	"math/bits"
	"testing"
)

func LutDivByPower5(u64 uint64, n int) bool {
	if n <= 0 || n*2 > len(lut) {
		return n == 0
	}
	lo0 := lut[n*2-2]
	hi0 := lut[n*2-1]
	hi1, lo1 := bits.Mul64(u64, lo0)
	hi1 += u64 * hi0
	return hi1 < hi0 || hi1 == hi0 && lo1 < lo0
}

func StdDivByPower5(m uint64, k int) bool {
	if m == 0 {
		return true
	}
	for i := 0; i < k; i++ {
		if m%5 != 0 {
			return false
		}
		m /= 5
	}
	return true
}

func ReciDivByPower5(m uint64, k int) bool {
	if m == 0 {
		return true
	}
	for i := 0; i < k; i++ {
		quo, div := ReciDiv5x(m)
		if !div {
			return false
		}
		m = quo
	}
	return true
}

func ReciDiv5x(v uint64) (quo uint64, div bool) {
	quo, _ = bits.Mul64(v, 0xCCCC_CCCC_CCCC_CCCD)
	quo >>= 2
	return quo, v == quo+quo<<2
}

func ReciDiv5(v uint64) (quo, rem uint64, divisible bool) {
	quo, _ = bits.Mul64(v, 0xCCCC_CCCC_CCCC_CCCD)
	quo >>= 2
	rem = v - (quo + quo<<2)
	return quo, rem, rem == 0
}

func BasicDiv5(v uint64) (quo, rem uint64, divisible bool) {
	return v / 5, v % 5, v%5 == 0
}

func BenchmarkLut(b *testing.B) {
	for range b.N {
		for v := 0; v < 0xFFF_FFFF; v += 5 {
			for k := 1; k < 24; k++ {
				div = LutDivByPower5(uint64(v), k)
			}
		}
	}
}

func BenchmarkReci(b *testing.B) {
	for range b.N {
		for v := 0; v < 0xFFF_FFFF; v += 5 {
			for k := 1; k < 24; k++ {
				div = ReciDivByPower5(uint64(v), k)
			}
		}
	}
}

func BenchmarkStd(b *testing.B) {
	for range b.N {
		for v := 0; v < 0xFFF_FFFF; v += 5 {
			for k := 1; k < 24; k++ {
				div = ReciDivByPower5(uint64(v), k)
			}
		}
	}
}

var div bool

var (
	lut [27 * 2]uint64
)

func init() {
	div := uint64(5)
	for n := 0; n < len(lut); n += 2 {
		hi, lo := castspell(div)
		lut[n] = lo
		lut[n+1] = hi
		div *= 5
	}
}

func castspell(div uint64) (uint64, uint64) {
	var recip, bigdiv big.Int
	bigdiv.SetUint64(div)
	recip.SetBit(&recip, 128, 1)
	recip.Div(&recip, &bigdiv)
	recip.Add(&recip, big.NewInt(1))
	if bits.UintSize != 64 {
		panic("not implemented on 32-bit")
	}
	words := recip.Bits()
	var hi, lo uint64
	if len(words) > 0 {
		lo = uint64(words[0])
	}
	if len(words) > 1 {
		hi = uint64(words[1])
	}
	return hi, lo
}
