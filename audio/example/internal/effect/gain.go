package effect

import (
	"github.com/egonelbre/exp/audio"
	"github.com/egonelbre/exp/audio/example/internal/atomic2"
)

type Gain struct {
	Value   atomic2.Float64
	current float64
}

func (gain *Gain) Process(buf audio.Buffer) error {
	target := gain.Value.Get()
	current := gain.current

	if target == current {
		switch buf := buf.(type) {
		case *audio.BufferF32:
			current := float32(current)
			for i, v := range buf.Data {
				buf.Data[i] = v * current
			}
		case *audio.BufferF64:
			for i, v := range buf.Data {
				buf.Data[i] = v * current
			}
		default:
			return audio.ErrUnknownBuffer
		}
		return nil
	}

	channelCount := buf.ChannelCount()
	switch buf := buf.(type) {
	case *audio.BufferF32:
		for i := 0; i < len(buf.Data); i += channelCount {
			for k := 0; k < channelCount; k++ {
				buf.Data[k] *= float32(current)
			}
			current = (current + target) * 0.5
		}
	case *audio.BufferF64:
		for i := 0; i < len(buf.Data); i += channelCount {
			for k := 0; k < channelCount; k++ {
				buf.Data[k] *= current
			}
			current = (current + target) * 0.5
		}
	default:
		return audio.ErrUnknownBuffer
	}

	if atomic2.AlmostEqual64(current, target) {
		gain.current = target
	} else {
		gain.current = current
	}

	return nil
}
