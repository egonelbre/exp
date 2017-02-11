package slice

import (
	"testing"

	"github.com/egonelbre/exp/audio"
)

func TestInterleave32_Mono(t *testing.T) {
	buf := audio.NewBufferF32Frames(
		audio.Format{ChannelCount: 1, SampleRate: 44100}, 4)

	copy(buf.Channel(0), []float32{1, 2, 3, 4})

	out := make([]float32, 4)
	Interleave32(buf, 1, out)

	exp := []float32{1, 2, 3, 4}
	if !Equal32(out, exp) {
		t.Errorf("got %v want %v", out, exp)
	}
}

func TestInterleave32_Stereo(t *testing.T) {
	buf := audio.NewBufferF32Frames(
		audio.Format{ChannelCount: 2, SampleRate: 44100}, 4)

	copy(buf.Channel(0), []float32{1, 2, 3, 4})
	copy(buf.Channel(1), []float32{5, 6, 7, 8})

	out := make([]float32, 8)
	Interleave32(buf, 2, out)

	exp := []float32{1, 5, 2, 6, 3, 7, 4, 8}
	if !Equal32(out, exp) {
		t.Errorf("got %v want %v", out, exp)
	}
}
