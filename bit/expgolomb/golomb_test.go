package expgolomb

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/egonelbre/exp/bit"
)

func TestIntRandomSmall(t *testing.T) { testInt(t, 5) }
func TestIntRandomLarge(t *testing.T) { testInt(t, 1251) }

func testInt(t *testing.T, n int) {
	values := make([]int, n)
	for i := range values {
		if i%2 == 0 {
			values[i] = rand.Int()
		} else {
			values[i] = -rand.Int()
		}
	}

	var buf bytes.Buffer
	w := bit.NewWriter(&buf)
	for _, v := range values {
		WriteInt(w, v)
	}

	err := w.Close()
	if err != nil {
		t.Errorf("w.Close: %v", err)
		return
	}

	r := bit.NewReader(&buf)
	for i, exp := range values {
		got, err := ReadInt(r), r.Error()
		if err != nil {
			t.Errorf("ReadInt: %v", err)
			return
		}
		if got != exp {
			t.Errorf("%v: %d got %v expected %v", values, i, got, exp)
			return
		}
	}
}

func TestUintRandomSmall(t *testing.T) { testUint(t, 5) }
func TestUintRandomLarge(t *testing.T) { testUint(t, 1251) }

func testUint(t *testing.T, n int) {
	values := make([]uint, n)
	for i := range values {
		values[i] = uint(rand.Int())
	}

	var buf bytes.Buffer
	w := bit.NewWriter(&buf)
	for _, v := range values {
		WriteUint(w, v)
	}

	err := w.Close()
	if err != nil {
		t.Errorf("w.Close: %v", err)
		return
	}

	r := bit.NewReader(&buf)
	for i, exp := range values {
		got, err := ReadUint(r), r.Error()
		if err != nil {
			t.Errorf("ReadUint: %v", err)
			return
		}
		if got != exp {
			t.Errorf("%v: %d got %v expected %v", values, i, got, exp)
			return
		}
	}
}
