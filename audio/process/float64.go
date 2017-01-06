package process

import "github.com/egonelbre/exp/audio"

func Mono64(out *audio.Buffer64, sample func(float64) float64) {
	switch out.ChannelCount {
	case 1:
		for i := 0; i < len(out.Data); i++ {
			out.Data[i] = sample(out.Data[i])
		}
	case 2:
		for i := 0; i < len(out.Data); i += 2 {
			in := (out.Data[i] + out.Data[i+1])
			v := sample(in * 0.5)
			out.Data[i] = v
			out.Data[i+1] = v
		}
	default:
		scale := 1.0 / float64(out.ChannelCount)
		for i := 0; i < len(out.Data); i += int(out.ChannelCount) {
			t := float64(0.0)
			for k := 0; k < int(out.ChannelCount); k++ {
				t += out.Data[i+k]
			}

			v := sample(t * scale)

			for k := 0; k < int(out.ChannelCount); k++ {
				out.Data[i+k] = v
			}
		}
	}
}

func Stereo64(out *audio.Buffer64, sample func(float64, float64) (float64, float64)) {
	switch out.ChannelCount {
	case 1:
		for i := 0; i < len(out.Data); i++ {
			left, right := sample(out.Data[i], out.Data[i])
			out.Data[i] = (left + right) * 0.5
		}
	case 2:
		for i := 0; i < len(out.Data); i += 2 {
			out.Data[i], out.Data[i+1] = sample(out.Data[i], out.Data[i+1])
		}
	default:
		for i := 0; i < len(out.Data); i += int(out.ChannelCount) {
			out.Data[i], out.Data[i+1] = sample(out.Data[i], out.Data[i+1])
		}
	}
}
