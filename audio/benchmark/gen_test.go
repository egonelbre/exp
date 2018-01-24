package benchmark_audio

import (
	"strconv"
	"testing"
)

const (
	SampleRate     = 44100
	TestFrameCount = SampleRate / 10
	TestBytes      = TestFrameCount * 4
)

var (
	tmp   = make([]float32, 7*TestFrameCount)
	tmp_1 = make([][1]float32, TestFrameCount)
	tmp_2 = make([][2]float32, TestFrameCount)
	tmp_5 = make([][5]float32, TestFrameCount)
	tmp_7 = make([][7]float32, TestFrameCount)
)

func BenchmarkSawtooth_Dynamic(b *testing.B) {
	for _, nch := range []int{1, 2, 5, 7} {
		b.Run("Ch"+strconv.Itoa(nch), func(b *testing.B) {
			b.SetBytes(TestBytes)
			for i := 0; i < b.N; i++ {
				buf := &Buffer{byte(nch), tmp[:TestFrameCount*nch]}
				Sawtooth_Dynamic(buf, SampleRate, 440)
			}
		})
	}
}

func BenchmarkSawtooth_Switch(b *testing.B) {
	for _, nch := range []int{1, 2, 5, 7} {
		b.Run("Ch"+strconv.Itoa(nch), func(b *testing.B) {
			b.SetBytes(TestBytes)
			for i := 0; i < b.N; i++ {
				buf := &Buffer{byte(nch), tmp[:TestFrameCount*nch]}
				Sawtooth_Switch(buf, SampleRate, 440)
			}
		})
	}
}

func BenchmarkSawtooth_Baseline(b *testing.B) {
	b.Run("Ch1", func(b *testing.B) {
		b.SetBytes(TestBytes)
		for i := 0; i < b.N; i++ {
			Sawtooth_Baseline_Ch1(Buffer_Ch1(tmp_1), SampleRate, 440)
		}
	})

	b.Run("Ch2", func(b *testing.B) {
		b.SetBytes(TestBytes)
		for i := 0; i < b.N; i++ {
			Sawtooth_Baseline_Ch2(Buffer_Ch2(tmp_2), SampleRate, 440)
		}
	})

	b.Run("Ch5", func(b *testing.B) {
		b.SetBytes(TestBytes)
		for i := 0; i < b.N; i++ {
			Sawtooth_Baseline_Ch5(Buffer_Ch5(tmp_5), SampleRate, 440)
		}
	})

	b.Run("Ch7", func(b *testing.B) {
		b.SetBytes(TestBytes)
		for i := 0; i < b.N; i++ {
			Sawtooth_Baseline_Ch7(Buffer_Ch7(tmp_7), SampleRate, 440)
		}
	})
}
