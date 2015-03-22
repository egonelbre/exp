package bit

import (
	"math/rand"
	"testing"
)

func TestZigRandom(t *testing.T) {
	for i := 0; i < 15315; i += 1 {
		x := rand.Int63()
		if rand.Intn(2) == 1 {
			x = -x
		}
		rx := ZEncode(x)
		rrx := ZDecode(rx)
		if rrx != x {
			t.Fatalf("failed %v: %v / %v", x, rx, rrx)
		}
	}
}
