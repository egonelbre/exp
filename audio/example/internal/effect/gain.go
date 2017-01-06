package effect

import (
	"github.com/egonelbre/exp/audio"
	"github.com/egonelbre/exp/audio/example/internal/atomic2"
)

type Gain struct {
	Value   atomic2.Float64
	current float64
}

func (gain *Gain) Process32(buf *audio.Buffer32) (int, error) {
	target := float32(gain.Value.Get())
	current := float32(gain.current)

	if target == current {
		for i, v := range buf.Data {
			buf.Data[i] = v * current
		}
		return len(buf.Data), nil
	}

	switch buf.ChannelCount {
	case 1:
		for i, v := range buf.Data {
			buf.Data[i] = v * current
			current = (current + target) * 0.5
		}

	case 2:
		// not sure what is the fastest way to code this
		for i, v := range buf.Data {
			buf.Data[i] = v * current
			if i&1 == 1 {
				current = (current + target) * 0.5
			}
		}
	default:
		for i := 0; i < len(buf.Data); i += int(buf.ChannelCount) {
			for k := 0; k < int(buf.ChannelCount); k++ {
				buf.Data[i+k] *= current
			}
			current = (current + target) * 0.5
		}
	}

	if atomic2.AlmostEqual32(current, target) {
		gain.current = gain.Value.Get()
		return len(buf.Data), nil
	}
	gain.current = float64(current)

	return len(buf.Data), nil
}

func (gain *Gain) Process64(buf *audio.Buffer64) (int, error) {
	target := gain.Value.Get()
	current := gain.current

	if target == current {
		for i, v := range buf.Data {
			buf.Data[i] = v * current
		}
		return len(buf.Data), nil
	}

	switch buf.ChannelCount {
	case 1:
		for i, v := range buf.Data {
			buf.Data[i] = v * current
			current = (current + target) * 0.5
		}

	case 2:
		// not sure what is the fastest way to code this
		for i, v := range buf.Data {
			buf.Data[i] = v * current
			if i&1 == 1 {
				current = (current + target) * 0.5
			}
		}
	default:
		for i := 0; i < len(buf.Data); i += int(buf.ChannelCount) {
			for k := 0; k < int(buf.ChannelCount); k++ {
				buf.Data[i+k] *= current
			}
			current = (current + target) * 0.5
		}
	}

	if atomic2.AlmostEqual64(current, target) {
		gain.current = target
		return len(buf.Data), nil
	}
	gain.current = current

	return len(buf.Data), nil
}
