package combiner

import (
	"testing"
)

func newBasic(exe Executor) Combiner { return NewBasic(exe) }
func TestBasic(t *testing.T)         { test(t, newBasic) }
func BenchmarkBasic(b *testing.B)    { bench(b, newBasic) }
