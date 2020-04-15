package parseint_test

import (
	"reflect"
	"strconv"
	"testing"
	"unsafe"
)

func BenchmarkStdAtoi(b *testing.B) {
	var out int
	for i := 0; i < b.N; i++ {
		out, _ = strconv.Atoi("14")
		out, _ = strconv.Atoi("148487491")
	}
	_ = out
}

func BenchmarkStdAtoiBool(b *testing.B) {
	var out int
	for i := 0; i < b.N; i++ {
		out, _ = parseAtoiBool("14")
		out, _ = parseAtoiBool("148487491")
	}
	_ = out
}

func BenchmarkStdAtoiNeg(b *testing.B) {
	var out int
	for i := 0; i < b.N; i++ {
		out = parseAtoiNeg("14")
		out = parseAtoiNeg("148487491")
	}
	_ = out
}

func BenchmarkTailPtr(b *testing.B) {
	var out int
	for i := 0; i < b.N; i++ {
		_ = parseTailPtr("14", &out)
		_ = parseTailPtr("148487491", &out)
	}
	_ = out
}

func BenchmarkTail(b *testing.B) {
	var out int
	for i := 0; i < b.N; i++ {
		out, _ = parseTail("14")
		out, _ = parseTail("148487491")
	}
	_ = out
}

func BenchmarkTailNeg(b *testing.B) {
	var out int
	for i := 0; i < b.N; i++ {
		out = parseTailNeg("14")
		out = parseTailNeg("148487491")
	}
	_ = out
}

func BenchmarkJeff1(b *testing.B) {
	var out int
	for i := 0; i < b.N; i++ {
		out = parseJeff1("14")
		out = parseJeff1("148487491")
	}
	_ = out
}

func BenchmarkUnsafe(b *testing.B) {
	var out int
	for i := 0; i < b.N; i++ {
		out = parseUnsafe("14")
		out = parseUnsafe("148487491")
	}
	_ = out
}

func parseAtoiBool(s string) (int, bool) {
	if len(s) == 0 || len(s) > len("999999999") {
		return 0, false
	}
	n := 0
	for _, c := range []byte(s) {
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return n, true
}

func parseAtoiNeg(s string) int {
	if len(s) == 0 || len(s) > len("999999999") {
		return -1
	}
	n := 0
	for _, c := range []byte(s) {
		if c < '0' || c > '9' {
			return -1
		}
		n = n*10 + int(c-'0')
	}
	return n
}

func parseTailPtr(s string, dst *int) bool {
	if len(s) == 0 || len(s) > len("999999999") {
		*dst = 0
		return false
	}
	n := 0
	for i := 0; i < len(s); i++ {
		d := s[i] - '0'
		if d > 9 {
			*dst = 0
			return false
		}
		n = n*10 + int(d)
	}
	*dst = n
	return true
}

func parseTail(s string) (int, bool) {
	if len(s) == 0 || len(s) > len("999999999") {
		return 0, false
	}
	n := 0
	for i := 0; i < len(s); i++ {
		d := s[i] - '0'
		if d > 9 {
			return 0, false
		}
		n = n*10 + int(d)
	}
	return n, true
}

func parseTailNeg(s string) int {
	if len(s) == 0 || len(s) > len("999999999") {
		return -1
	}
	n := 0
	for i := 0; i < len(s); i++ {
		d := s[i] - '0'
		if d > 9 {
			return -1
		}
		n = n*10 + int(d)
	}
	return n
}

var deltasJeff1 [10]int

func init() {
	v := 0
	for i := range deltasJeff1 {
		deltasJeff1[i] = v
		v = v*10 + '0'
	}

	// we do this so that we return -1 when len(s) == 0
	// skips having to check for the len(s) == 0 condition
	deltasJeff1[0] = 1
}

func parseJeff1(s string) (n int) {
	var c byte
	var i int

next:
	if i < len(s) {
		c = s[i]
		i = i + 1
		n = 10*n + int(c)

		if c-'0' > 9 {
			return -1
		}
		goto next
	}

	if len(s) < len(deltas) {
		return n - deltas[len(s)]
	}
	return -1
}

var deltas [10]int

func init() {
	v := 0
	for i := range deltas {
		deltas[i] = v
		v = v*10 + '0'
	}
	deltas[0] = 1
}

// because life without unsafe and reflect would be too safe.
func parseUnsafe(s string) (n int) {
	sp := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	lp := uintptr(sp) + uintptr(len(s))
next:
	if uintptr(sp) < lp {
		c := *(*byte)(sp)
		n = 10*n + int(c)
		if c-'0' > 9 {
			return -1
		}
		sp = unsafe.Pointer(uintptr(sp) + 1)
		goto next
	}
	if len(s) < len(deltas) {
		return n - deltas[len(s)]
	}
	return -1
}
