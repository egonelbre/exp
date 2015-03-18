package bit

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestWriteReadRandomBit(t *testing.T) {
	var buf bytes.Buffer
	w := NewWriter(&buf)

	values := make([]int, 1251)
	for i := range values {
		values[i] = rand.Intn(1)
		w.WriteBit(values[i])
	}
	w.Close()

	r := NewReader(&buf)
	for i, exp := range values {
		got, err := r.ReadBit()
		if err != nil {
			t.Errorf("r.ReadBits: %v", err)
			return
		}
		if got != exp {
			t.Errorf("%v: %d got %v expected %v", values, i, got, exp)
			return
		}
	}
}

func TestWriteReadBits(t *testing.T) {
	testReadWrite(t, []pair{
		{0xAC, 8},
		{1, 3},
		{2, 3},
		{3, 3},
		{4, 3},
		{5, 3},
		{0xEF, 8},
		{0x7F, 7},
		{0xCAFEBABE, 32},
	})
}

func TestWriteReadRandomBitsSmall(t *testing.T) { testRandomBits(t, 5) }
func TestWriteReadRandomBitsLarge(t *testing.T) { testRandomBits(t, 1251) }

func testRandomBits(t *testing.T, N int) {
	values := make([]pair, N)
	for i := range values {
		width := uint(rand.Intn(61) + 1)
		value := uint64(rand.Intn(1<<width - 1))
		values[i] = pair{value, width}
	}
	testReadWrite(t, values)
}

type pair struct {
	value uint64
	width uint
}

func testReadWrite(t *testing.T, values []pair) {
	var buf bytes.Buffer
	w := NewWriter(&buf)
	for _, p := range values {
		w.WriteBits(p.value, p.width)
	}

	err := w.Close()
	if err != nil {
		t.Errorf("w.Close: %v", err)
		return
	}

	r := NewReader(&buf)
	for i, exp := range values {
		got, err := r.ReadBits(exp.width)
		if err != nil {
			t.Errorf("r.ReadBit: %v", err)
			return
		}
		if got != exp.value {
			t.Errorf("%v: %d got %v expected %v", values, i, got, exp.value)
			return
		}
	}
}
