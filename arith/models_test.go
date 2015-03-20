package arith

import (
	"math/rand"
	"testing"
)

// Tests a model provided by the constructor
func ModelTest(t *testing.T, N int, model func() Model) {
	encm, decm := model(), model()
	mask := uint(1<<encm.NBits() - 1)

	bits := make([]uint, N)
	for i := range bits {
		bits[i] = uint(rand.Int()) & mask
	}

	enc := NewEncoder()
	for _, bit := range bits {
		encm.Encode(enc, bit)
	}
	enc.Close()

	data := enc.Bytes()

	dec := NewDecoder(data)
	for i, v := range bits {
		x := decm.Decode(dec) & mask
		if x != v {
			t.Fatalf("fail %v: %v got %v exp %v", bits, i, x, v)
		}
	}
}

func TestShiftModel(t *testing.T) {
	ModelTest(t, 3, func() Model { return &Shift{Prob(0.5), 3} })
	ModelTest(t, 21, func() Model { return &Shift{Prob(0.1), 2} })
	ModelTest(t, 312532, func() Model { return &Shift{Prob(0.9), 1} })
}

func TestShift2Model(t *testing.T) {
	ModelTest(t, 3, func() Model { return &Shift2{Prob(0.5), 3, Prob(0.1), 1} })
	ModelTest(t, 21, func() Model { return &Shift2{Prob(0.25), 2, Prob(0.25), 1} })
	ModelTest(t, 312532, func() Model { return &Shift2{Prob(0.9), 1, Prob(0.5), 1} })
}

func TestTreeModel(t *testing.T) {
	B := func() Model { return &Shift{Prob(0.43), 3} }
	ModelTest(t, 3, func() Model { return NewTree(1, B) })
	ModelTest(t, 3, func() Model { return NewTree(2, B) })
	ModelTest(t, 21, func() Model { return NewTree(2, B) })
	ModelTest(t, 4124, func() Model { return NewTree(3, B) })
	ModelTest(t, 41244, func() Model { return NewTree(3, B) })

	ModelTest(t, 41244, func() Model {
		m := NewEmptyTree(2)
		m[0] = &Shift{Prob(0.43), 3}
		m[1] = &Shift{Prob(0.9), 2}
		m[2] = &Shift2{Prob(0.5), 3, Prob(0.1), 1}
		return m
	})
}
