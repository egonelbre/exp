package bit

import (
	"math/rand"
	"testing"
)

func TestAbsRandom(t *testing.T) {
	for i := 0; i < 1521; i += 1 {
		x := rand.Int63() * (rand.Int63n(2)*2 - 1)
		rx := AbsEncode(x)
		rrx := AbsDecode(rx)
		if rrx != x {
			t.Errorf("failed %v: %v / %v", x, rx, rrx)
		}
	}
}
