package combiner

import "testing"

func newBounded(exe Executor) Combiner { return NewBounded(exe, 16) }
func TestBounded(t *testing.T)         { test(t, newBounded) }
func BenchmarkBounded(b *testing.B)    { bench(b, newBounded) }

func newBoundedU(exe Executor) Combiner { return NewBoundedU(exe, 16) }
func TestBoundedU(t *testing.T)         { test(t, newBoundedU) }
func BenchmarkBoundedU(b *testing.B)    { bench(b, newBoundedU) }
