package combiner

import (
	"testing"
)

func newTBB(exe Executor) Combiner { return NewTBB(exe) }
func TestTBB(t *testing.T)         { test(t, newTBB) }
func BenchmarkTBB(b *testing.B)    { bench(b, newTBB) }
