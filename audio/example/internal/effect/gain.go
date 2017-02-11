package effect

import (
	"github.com/egonelbre/exp/audio"
	"github.com/egonelbre/exp/audio/example/internal/atomic2"
	"github.com/egonelbre/exp/audio/slice"
)

type Gain struct {
	Value   atomic2.Float64
	current float64
}

func NewGain(value float64) *Gain {
	gain := &Gain{}
	gain.Value.Set(value)
	return gain
}

func (gain *Gain) Process(buf audio.Buffer) error {
	target := gain.Value.Get()
	current := gain.current

	channelCount := buf.ChannelCount()
	if target == current {
		switch buf := buf.(type) {
		case *audio.BufferF32:
			for k := 0; k < channelCount; k++ {
				slice.Scale32(buf.Channel(k), float32(current))
			}
		case *audio.BufferF64:
			for k := 0; k < channelCount; k++ {
				slice.Scale64(buf.Channel(k), current)
			}
		default:
			return audio.ErrUnknownBuffer
		}
		return nil
	}

	var active float64

	switch buf := buf.(type) {
	case *audio.BufferF32:
		for k := 0; k < channelCount; k++ {
			active = current
			data := buf.Channel(k)
			for i := range data {
				data[i] *= float32(active)
				active = (active + target) * 0.5
			}
		}
	case *audio.BufferF64:
		for k := 0; k < channelCount; k++ {
			active = current
			data := buf.Channel(k)
			for i := range data {
				data[i] *= float64(active)
				active = (active + target) * 0.5
			}
		}
	default:
		return audio.ErrUnknownBuffer
	}

	current = active

	if atomic2.AlmostEqual64(current, target) {
		gain.current = target
	} else {
		gain.current = current
	}

	return nil
}
