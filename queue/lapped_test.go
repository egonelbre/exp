package queue

import "testing"

func TestLapped(t *testing.T) {
	test(t, func(size int) Queue { return NewLapped(size) })
}

func BenchmarkLapped(b *testing.B) {
	bench(b, func(size int) Queue { return NewLapped(size) })
}
