package generate

import "github.com/egonelbre/exp/audio"

func Mono32(out *audio.Buffer32, sample func() float32) {
	switch out.Channels {
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
		for i := 0; i < len(out.Data); i += int(out.Channels) {
			v := sample()
			for k := 0; k < int(out.Channels); k++ {
				out.Data[i+k] = v
			}
		}
	}
}

func Stereo32(out *audio.Buffer32, sample func() (float32, float32)) {
	switch out.Channels {
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
		for i := 0; i < len(out.Data); i += int(out.Channels) {
			left, right := sample()
			out.Data[i] = left
			out.Data[i+1] = right
			for k := 2; k < int(out.Channels); k++ {
				out.Data[i+k] = 0
			}
		}
	}
}
