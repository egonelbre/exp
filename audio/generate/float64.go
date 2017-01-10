package generate

import "github.com/egonelbre/exp/audio"

func MonoF64(out audio.Buffer, sample func() float64) error {
	channelCount := out.ChannelCount()
	switch out := out.(type) {
	case *audio.BufferF32:
		for i := 0; i < len(out.Data); i += channelCount {
			v := float32(sample())
			for k := 0; k < channelCount; k++ {
				out.Data[i+k] = v
			}
		}
	case *audio.BufferF64:
		for i := 0; i < len(out.Data); i += channelCount {
			v := sample()
			for k := 0; k < channelCount; k++ {
				out.Data[i+k] = v
			}
		}
	default:
		//TODO: implement the slowest path
		return audio.ErrUnknownBuffer
	}
	return nil
}

func StereoF64(out audio.Buffer, sample func() (float64, float64)) error {
	channelCount := out.ChannelCount()
	switch out := out.(type) {
	case *audio.BufferF32:
		switch channelCount {
		case 1:
			for i := 0; i < len(out.Data); i++ {
				left, right := sample()
				out.Data[i] = float32(left+right) * 0.5
			}
		case 2:
			for i := 0; i < len(out.Data); i += 2 {
				left, right := sample()
				out.Data[i], out.Data[i+1] = float32(left), float32(right)
			}
		default:
			for i := 0; i < len(out.Data); i += channelCount {
				left, right := sample()
				out.Data[i], out.Data[i+1] = float32(left), float32(right)
				for k := 2; k < channelCount; k++ {
					out.Data[i+k] = 0
				}
			}
		}
	case *audio.BufferF64:
		switch channelCount {
		case 1:
			for i := 0; i < len(out.Data); i++ {
				left, right := sample()
				out.Data[i] = (left + right) * 0.5
			}
		case 2:
			for i := 0; i < len(out.Data); i += 2 {
				out.Data[i], out.Data[i+1] = sample()
			}
		default:
			for i := 0; i < len(out.Data); i += channelCount {
				left, right := sample()
				out.Data[i], out.Data[i+1] = left, right
				for k := 2; k < channelCount; k++ {
					out.Data[i+k] = 0
				}
			}
		}
	default:
		//TODO: implement the slowest path
		return audio.ErrUnknownBuffer
	}
	return nil
}
