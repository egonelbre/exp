package arith

import (
	"math/rand"
	"testing"
)

type pair struct {
	bit  uint
	prob P
}

func TestRandomEncodeDecode(t *testing.T) {
	bits := make([]pair, 2141)
	for i := range bits {
		bits[i].bit = uint(rand.Intn(2))
		bits[i].prob = P(rand.Intn(MaxP))
	}

	enc := NewEncoder()
	for _, v := range bits {
		enc.Encode(v.bit, v.prob)
	}
	enc.Close()
	data := enc.Bytes()

	dec := NewDecoder(data)
	for i, v := range bits {
		x := dec.Decode(v.prob)
		if x != v.bit {
			t.Fatalf("fail %v: %v got %v exp %v", bits, i, x, v.bit)
		}
	}
}

func TestPacking(t *testing.T) {
	const N = 1 << 16
	enc := NewEncoder()
	p := Prob(0.1)
	for i := 0; i < N; i++ {
		if rand.Intn(100) < 10 {
			enc.Encode(1, p)
		} else {
			enc.Encode(0, p)
		}
	}
	enc.Close()

	if len(enc.Bytes()) > N/8 {
		t.Fatalf("No packing: %v vs %v", len(enc.Bytes()), N/8)
	}
	enc.Bytes()
}
