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

func (sine *Sine) Process32(buf *audio.Buffer32) (int, error) {
	inv := 2.0 * math.Pi / float64(buf.SampleRate)
	target := sine.Frequency.Get() * inv
	speed := sine.frequency * inv
	phase := sine.phase

	if atomic2.AlmostEqual64(target, speed) {
		generate.Mono32(buf, func() float32 {
			r := math.Cos(phase)
			phase += speed
			return float32(r)
		})
		sine.phase = phase
		return len(buf.Data), nil
	}

	generate.Mono32(buf, func() float32 {
		r := math.Cos(phase)
		phase += speed
		speed = (speed + target) * 0.5
		return float32(r)
	})
	sine.phase = phase
	sine.frequency = speed / inv

	return len(buf.Data), nil
}

func (sine *Sine) Process64(buf *audio.Buffer64) (int, error) {
	inv := 2.0 * math.Pi / float64(buf.SampleRate)
	target := sine.Frequency.Get() * inv
	speed := sine.frequency * inv
	phase := sine.phase

	if atomic2.AlmostEqual64(target, speed) {
		generate.Mono64(buf, func() float64 {
			r := math.Cos(phase)
			phase += speed
			return r
		})
		sine.phase = phase
		return len(buf.Data), nil
	}

	generate.Mono64(buf, func() float64 {
		r := math.Cos(phase)
		phase += speed
		speed = (speed + target) * 0.5
		return r
	})

	sine.phase = phase
	sine.frequency = speed / inv

	return len(buf.Data), nil
}
