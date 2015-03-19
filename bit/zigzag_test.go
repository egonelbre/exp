package bit

import (
	"math/rand"
	"testing"
)

func TestZigRandom(t *testing.T) {
	for i := 0; i < 1521; i += 1 {
		x := rand.Int63() * (rand.Int63n(2)*2 - 1)
		rx := ZEncode(x)
		rrx := ZDecode(rx)
		if rrx != x {
			t.Errorf("failed %v: %v / %v", x, rx, rrx)
		}
	}
}
