package bench_test

import (
	"testing"
	"strconv"
)

var testBytesValues = func() (xs [][]byte) {
	for i := 0; i < 1000; i++ {
		xs = append(xs, []byte(strconv.Itoa(i)+"-"+strconv.Itoa(i)))
	}
	return xs
}()

func BenchmarkByteInt(b *testing.B) {
	count := map[string]int{}
	for i := 0; i < b.N; i++ {
		for _, v := range testBytesValues {
			count[string(v)]++
		}
	}
}

func BenchmarkBytePointer(b *testing.B) {
	count := map[string]*int{}
	for i := 0; i < b.N; i++ {
		for _, v := range testBytesValues {
			if p, ok := count[string(v)]; ok {
				*p++
				continue
			}
			p := new(int)
			*p = 1
			count[string(v)] = p
		}
	}
}