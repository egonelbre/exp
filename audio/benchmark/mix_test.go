package benchmark_audio

import (
	"testing"
)

var (
	tmp_dst = make([]float32, TestFrameCount)
	tmp_a   = make([]float32, TestFrameCount)
	tmp_b   = make([]float32, TestFrameCount)
)

func BenchmarkMix_Dynamic(b *testing.B) {
	b.SetBytes(TestBytes)
	for i := 0; i < b.N; i++ {
		Mix_Dynamic(tmp_dst, tmp_a, tmp_b)
	}
}

func BenchmarkMix_Baseline(b *testing.B) {
	b.SetBytes(TestBytes)
	for i := 0; i < b.N; i++ {
		Mix_Baseline(tmp_dst, tmp_a, tmp_b)
	}
}

func BenchmarkMixInto_Dynamic(b *testing.B) {
	b.SetBytes(TestBytes)
	for i := 0; i < b.N; i++ {
		MixInto_Dynamic(tmp_a, tmp_b)
	}
}

func BenchmarkMixInto_Baseline(b *testing.B) {
	b.SetBytes(TestBytes)
	for i := 0; i < b.N; i++ {
		MixInto_Baseline(tmp_a, tmp_b)
	}
}
