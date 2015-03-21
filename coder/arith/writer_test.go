package arith

import (
	"bytes"
	"math/rand"
	"testing"
)

// Tests a model provided by the constructor
func WriterReaderTest(t *testing.T, N int, model func() Model) {
	data := make([]byte, N)
	for i := range data {
		data[i] = byte(rand.Int())
	}

	w := NewWriter(model())
	n, err := w.Write(data)
	if n != N || err != nil {
		t.Fatalf("failed write (%v,%v): %v", n, N, err)
	}

	if err := w.Close(); err != nil {
		t.Fatalf("failed close: %v", err)
	}

	got := make([]byte, N)

	r := NewReader(w.Bytes(), model())
	n, err = r.Read(got)
	if n != N || err != nil {
		t.Fatalf("failed read (%v,%v): %v", n, N, err)
	}

	if !bytes.Equal(got, data) {
		t.Fatalf("different content: %v, %v", got, data)
	}
}

func TestShiftWriter(t *testing.T) {
	WriterReaderTest(t, 3, func() Model { return &Shift{Prob(0.5), 3} })
	WriterReaderTest(t, 21, func() Model { return &Shift{Prob(0.1), 2} })
	WriterReaderTest(t, 312532, func() Model { return &Shift{Prob(0.9), 1} })
}
