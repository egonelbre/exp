package bit

import (
	"math/rand"
	"testing"
)

func TestScanPowers(t *testing.T) {
	for shift := uint(0); shift < 64; shift += 1 {
		sl := ScanLeft(1 << shift)
		sr := ScanRight(1 << shift)
		if sl != shift || sr != shift {
			t.Errorf("fail %v: got %v, %v", shift, sl, sr)
		}
	}
}

func TestScanLeftSmall(t *testing.T) {
	for i := 0; i < 1521; i += 1 {
		x := uint64(i * 31)
		got := ScanLeft(x)
		exp := slowScanLeft(x)
		if got != exp {
			t.Errorf("failed %b: got %v exp %v", x, got, exp)
			break
		}
	}
}

func TestScanLeftRandom(t *testing.T) {
	for i := 0; i < 1521; i += 1 {
		x := uint64(rand.Int63())
		got := ScanLeft(x)
		exp := slowScanLeft(x)
		if got != exp {
			t.Errorf("failed %b: got %v exp %v", x, got, exp)
			break
		}
	}
}

func TestScanRightSmall(t *testing.T) {
	for i := 0; i < 1521; i += 1 {
		x := uint64(i * 31)
		got := ScanRight(x)
		exp := slowScanRight(x)
		if got != exp {
			t.Errorf("failed %b: got %v exp %v", x, got, exp)
			break
		}
	}
}

func TestScanRightRandom(t *testing.T) {
	for i := 0; i < 1521; i += 1 {
		x := uint64(rand.Int63())
		got := ScanRight(x)
		exp := slowScanRight(x)
		if got != exp {
			t.Errorf("failed %b: got %v exp %v", x, got, exp)
			break
		}
	}
}
