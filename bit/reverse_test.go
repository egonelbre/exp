package bit

import (
	"math/rand"
	"testing"
)

func TestReverseBitsRandom(t *testing.T) {
	for i := 0; i < 1521; i += 1 {
		width := uint(rand.Intn(31) + 1)
		x := uint(rand.Intn(1<<width - 1))
		rx := ReverseBits(x, width)
		rrx := ReverseBits(rx, width)
		if rrx != x {
			t.Errorf("failed <%v,%v>:  %v -> %v -> %v", x, width, x, rx, rrx)
		}
	}
}

func TestReverseBitsPowers(t *testing.T) {
	for width := uint(2); width < 8; width += 1 {
		exp := uint(1 << (width - 1))
		got := ReverseBits(1, width)
		if exp != got {
			t.Errorf("fail <1,%v>: got %v exp %v", width, got, exp)
		}
	}
}
