package slice

import "github.com/egonelbre/exp/audio"

func Zero64(data []float64) {
	for i := range data {
		data[i] = 0
	}
}

func Add64(dst, src []float64) {
	for i := range dst {
		dst[i] += src[i]
	}
}

func Scale64(data []float64, v float64) {
	for i := range data {
		data[i] *= v
	}
}

func ScaleLinearLerp64(data []float64, from, to float64) {
	inc := (to - from) / float64(len(data))
	for i := range data {
		data[i] *= from
		from += inc
	}
}

func Interleave64(buf audio.Buffer, nchan int, dst []float64) (frameCount int) {
	channelCount := buf.ChannelCount()
	if channelCount != nchan {
		// TODO mixing/splitting when needed
		panic("channel count does not match")
	}

	maxFrames := len(dst) / nchan
	if buf.FrameCount() < maxFrames {
		maxFrames = buf.FrameCount()
	}

	switch buf := buf.(type) {
	case *audio.BufferF32:
		for k := 0; k < channelCount; k++ {
			src := buf.Channel(k)
			for di, si := k, 0; di < maxFrames; di, si = di+nchan, si+1 {
				dst[di] = float64(src[si])
			}
		}
	case *audio.BufferF64:
		for k := 0; k < channelCount; k++ {
			src := buf.Channel(k)
			for di, si := k, 0; di < maxFrames; di, si = di+nchan, si+1 {
				dst[di] = float64(src[si])
			}
		}
	default:
		panic("missing")
	}

	return maxFrames
}

func Equal64(a, b []float64) bool {
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
