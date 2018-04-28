package queue

import "testing"

func TestChan(t *testing.T) {
	test(t, func(size int) Queue { return NewChan(size) })
}

func BenchmarkChan(b *testing.B) {
	bench(b, func(size int) Queue { return NewChan(size) })
}
