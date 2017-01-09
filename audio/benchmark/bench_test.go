package benchmark

import "testing"

var scratch = &Buffer32{Format{44100, 2}, make([]float32, 64*2)}
var stack = []Processor{
	&Generator{440, 440, 0},
	&Generator{880, 880, 0},
	&Generator{220, 220, 0},
	&Generator{110, 110, 0},
}

func BenchmarkInterface(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, p := range stack {
			p.ProcessInterface(scratch)
		}
	}
}

func BenchmarkDirect(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, p := range stack {
			p.ProcessDirect(scratch)
		}
	}
}

func BenchmarkInline(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, p := range stack {
			p.ProcessInline(scratch)
		}
	}
}

func BenchmarkInterfaceInline(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, p := range stack {
			p.ProcessInterfaceInline(scratch)
		}
	}
}
