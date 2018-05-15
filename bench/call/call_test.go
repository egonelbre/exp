package call

import (
	"testing"

	"github.com/egonelbre/exp/bench/call/asm"
)

func BenchmarkCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CNop()
	}
}

func BenchmarkGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Nop()
	}
}

func BenchmarkAsm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asm.Nop()
	}
}
