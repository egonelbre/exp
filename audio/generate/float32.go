package generate

import (
	"github.com/egonelbre/exp/audio"
	"github.com/egonelbre/exp/audio/slice"
)

func MonoF32(out audio.Buffer, sample func() float32) error {
	channelCount := out.ChannelCount()
	switch out := out.(type) {
	case *audio.BufferF32:
		main := out.Channel(0)
		for i := range main {
			main[i] = float32(sample())
		}
		for k := 1; k < channelCount; k++ {
			copy(out.Channel(k), main)
		}
	case *audio.BufferF64:
		main := out.Channel(0)
		for i := range main {
			main[i] = float64(sample())
		}
		for k := 1; k < channelCount; k++ {
			copy(out.Channel(k), main)
		}
	default:
		//TODO: implement the slowest path
		return audio.ErrUnknownBuffer
	}
	return nil
}

func StereoF32(out audio.Buffer, sample func() (float32, float32)) error {
	channelCount := out.ChannelCount()
	switch out := out.(type) {
	case *audio.BufferF32:
		if channelCount >= 2 {
			left, right := out.Channel(0), out.Channel(1)
			for i := range left {
				leftsample, rightsample := sample()
				left[i], right[i] = float32(leftsample), float32(rightsample)
			}
			for k := 2; k < channelCount; k++ {
				slice.Zero32(out.Channel(k))
			}
		} else {
			main := out.Channel(0)
			for i := range main {
				leftsample, rightsample := sample()
				main[i] = float32(leftsample+rightsample) * 0.5
			}
		}
	case *audio.BufferF64:
		if channelCount >= 2 {
			left, right := out.Channel(0), out.Channel(1)
			for i := range left {
				leftsample, rightsample := sample()
				left[i], right[i] = float64(leftsample), float64(rightsample)
			}
			for k := 2; k < channelCount; k++ {
				slice.Zero64(out.Channel(k))
			}
		} else {
			main := out.Channel(0)
			for i := range main {
				leftsample, rightsample := sample()
				main[i] = float64(leftsample+rightsample) * 0.5
			}
		}
	default:
		//TODO: implement the slowest path
		return audio.ErrUnknownBuffer
	}
	return nil
}
