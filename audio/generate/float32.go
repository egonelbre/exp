package generate

import "github.com/egonelbre/exp/audio"

func MonoF32(out audio.Buffer, sample func() float32) error {
	channelCount := out.ChannelCount()
	switch out := out.(type) {
	case *audio.BufferF32:
		switch channelCount {
		case 1:
			for i := 0; i < len(out.Data); i++ {
				out.Data[i] = sample()
			}
		case 2:
			for i := 0; i < len(out.Data); i += 2 {
				v := sample()
				out.Data[i] = v
				out.Data[i+1] = v
			}
		default:
			for i := 0; i < len(out.Data); i += channelCount {
				v := sample()
				for k := 0; k < channelCount; k++ {
					out.Data[i+k] = v
				}
			}
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
				out.Data[i] = left
				out.Data[i+1] = right
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
