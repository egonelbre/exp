package osc

import (
	"math"

	"github.com/egonelbre/exp/audio"
	"github.com/egonelbre/exp/audio/example/internal/atomic2"
	"github.com/egonelbre/exp/audio/generate"
)

type Sine struct {
	Frequency atomic2.Float64

	frequency float64
	phase     float64
}

func (sine *Sine) Process(buf audio.Buffer) error {
	inv := 2.0 * math.Pi / float64(buf.SampleRate())
	target := sine.Frequency.Get() * inv
	speed := sine.frequency * inv
	phase := sine.phase

	if atomic2.AlmostEqual64(target, speed) {
		err := generate.MonoF64(buf, func() float64 {
			r := math.Cos(phase)
			phase += speed
			return r
		})
		sine.phase = phase
		return err
	}

	err := generate.MonoF64(buf, func() float64 {
		r := math.Cos(phase)
		phase += speed
		speed = (speed + target) * 0.5
		return r
	})
	sine.phase = phase
	sine.frequency = speed / inv

	return err
}
