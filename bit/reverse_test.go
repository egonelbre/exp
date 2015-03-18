package bit

import (
	"math/rand"
	"testing"
)

func TestReverseRandom(t *testing.T) {
	for i := 0; i < 1521; i += 1 {
		width := uint(rand.Intn(61) + 1)
		x := uint64(rand.Intn(1<<width - 1))
		rx := Reverse(x, width)
		rrx := Reverse(rx, width)
		if rrx != x {
			t.Errorf("failed <%v,%v>:  %v -> %v -> %v", x, width, x, rx, rrx)
		}
	}
}

func TestReversePowers(t *testing.T) {
	for width := uint(2); width < 8; width += 1 {
		exp := uint64(1 << (width - 1))
		got := Reverse(1, width)
		if exp != got {
			t.Errorf("fail <1,%v>: got %v exp %v", width, got, exp)
		}
	}
}

func TestReverseOracle(t *testing.T) {
	for i := 0; i < 1521; i += 1 {
		width := uint(rand.Intn(61) + 1)
		x := uint64(rand.Intn(1<<width - 1))
		rx := Reverse(x, width)
		rx2 := slowReverse(x, width)
		if rx != rx2 {
			t.Errorf("failed <%v,%v>: got %v exp %v", x, width, rx, rx2)
		}
	}
}
