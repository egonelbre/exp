package slice

import "github.com/egonelbre/exp/audio"

func Zero32(data []float32) {
	for i := range data {
		data[i] = 0
	}
}

func Add32(dst, src []float32) {
	for i := range dst {
		dst[i] += src[i]
	}
}

func Scale32(data []float32, v float32) {
	for i := range data {
		data[i] *= v
	}
}

func ScaleLinearLerp32(data []float32, from, to float32) {
	inc := (to - from) / float32(len(data))
	for i := range data {
		data[i] *= from
		from += inc
	}
}

func Interleave32(buf audio.Buffer, nchan int, dst []float32) (frameCount int) {
	channelCount := buf.ChannelCount()
	if channelCount != nchan {
		// TODO mixing/splitting when needed
		panic("channel count does not match")
	}

	maxFrames := len(dst) / nchan
	if buf.FrameCount() < maxFrames {
		maxFrames = buf.FrameCount()
	}

	maxSamples := maxFrames * nchan

	switch buf := buf.(type) {
	case *audio.BufferF32:
		for k := 0; k < channelCount; k++ {
			src := buf.Channel(k)
			for di, si := k, 0; di < maxSamples; di, si = di+nchan, si+1 {
				dst[di] = float32(src[si])
			}
		}
	case *audio.BufferF64:
		for k := 0; k < channelCount; k++ {
			src := buf.Channel(k)
			for di, si := k, 0; di < maxSamples; di, si = di+nchan, si+1 {
				dst[di] = float32(src[si])
			}
		}
	default:
		panic("missing")
	}

	return maxFrames
}

func Equal32(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
