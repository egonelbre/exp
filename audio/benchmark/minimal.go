package benchmark

import (
	"errors"
	"math"
)

type Processor interface {
	ProcessInline(buf *Buffer32) (int, error)
	ProcessDirect(buf *Buffer32) (int, error)
	ProcessInterface(buf Buffer) (int, error)
	ProcessInterfaceInline(buf Buffer) (int, error)
}

type Buffer interface {
	FrameCount() int
	SampleCount() int
	SampleRate() int
	ChannelCount() int
}

type Format struct {
	SampleRateValue   int
	ChannelCountValue int
}

func (f Format) SampleRate() int   { return f.SampleRateValue }
func (f Format) ChannelCount() int { return f.ChannelCountValue }

type Buffer32 struct {
	Format
	Data []float32
}

func (buf *Buffer32) FrameCount() int  { return len(buf.Data) / buf.ChannelCountValue }
func (buf *Buffer32) SampleCount() int { return len(buf.Data) }

type Generator struct {
	Frequency float64
	frequency float64
	phase     float64
}

func (gen *Generator) ProcessInline(buf *Buffer32) (int, error) {
	inv := 2.0 * math.Pi / float64(buf.SampleRateValue)
	target := gen.Frequency * inv
	speed := gen.frequency * inv
	phase := gen.phase

	if target == speed {
		switch buf.ChannelCountValue {
		case 1:
			for i := 0; i < len(buf.Data); i++ {
				buf.Data[i] = float32(phase)
				phase += speed
			}
		case 2:
			for i := 0; i < len(buf.Data); i += 2 {
				buf.Data[i] = float32(phase)
				buf.Data[i+1] = buf.Data[i]
				phase += speed
			}
		default:
			return 0, errors.New("unsupported channel count")
		}
		gen.phase = phase
		return len(buf.Data), nil
	}

	switch buf.ChannelCountValue {
	case 1:
		for i := 0; i < len(buf.Data); i++ {
			buf.Data[i] = float32(phase)
			phase += speed
			speed = (speed + target) * 0.5
		}
	case 2:
		for i := 0; i < len(buf.Data); i += 2 {
			buf.Data[i] = float32(phase)
			buf.Data[i+1] = buf.Data[i]
			phase += speed
			speed = (speed + target) * 0.5
		}
	default:
		return 0, errors.New("unsupported channel count")
	}
	gen.phase = phase
	gen.frequency = speed / inv

	return len(buf.Data), nil
}

func (gen *Generator) ProcessDirect(buf *Buffer32) (int, error) {
	inv := 2.0 * math.Pi / float64(buf.SampleRateValue)
	target := gen.Frequency * inv
	speed := gen.frequency * inv
	phase := gen.phase

	if target == speed {
		MonoDirect32(buf, func() float32 {
			r := float32(phase)
			phase += speed
			return float32(r)
		})
		gen.phase = phase
		return len(buf.Data), nil
	}

	MonoDirect32(buf, func() float32 {
		r := float32(phase)
		phase += speed
		speed = (speed + target) * 0.5
		return float32(r)
	})
	gen.phase = phase
	gen.frequency = speed / inv

	return len(buf.Data), nil
}

func (gen *Generator) ProcessInterface(buf Buffer) (int, error) {
	inv := 2.0 * math.Pi / float64(buf.SampleRate())
	target := gen.Frequency * inv
	speed := gen.frequency * inv
	phase := gen.phase

	if target == speed {
		MonoInterface32(buf, func() float32 {
			r := float32(phase)
			phase += speed
			return float32(r)
		})
		gen.phase = phase
		return buf.SampleCount(), nil
	}

	MonoInterface32(buf, func() float32 {
		r := float32(phase)
		phase += speed
		speed = (speed + target) * 0.5
		return float32(r)
	})
	gen.phase = phase
	gen.frequency = speed / inv

	return buf.SampleCount(), nil
}

func (gen *Generator) ProcessInterfaceInline(buf Buffer) (int, error) {
	inv := 2.0 * math.Pi / float64(buf.SampleRate())
	target := gen.Frequency * inv
	speed := gen.frequency * inv
	phase := gen.phase

	if target == speed {
		switch buf := buf.(type) {
		case *Buffer32:
			switch buf.ChannelCountValue {
			case 1:
				for i := 0; i < len(buf.Data); i++ {
					buf.Data[i] = float32(phase)
					phase += speed
				}
			case 2:
				for i := 0; i < len(buf.Data); i += 2 {
					buf.Data[i] = float32(phase)
					buf.Data[i+1] = buf.Data[i]
					phase += speed
				}
			default:
				return 0, errors.New("unsupported channel count")
			}
		default:
			return 0, errors.New("invalid buffer type")
		}
		gen.phase = phase
		return buf.SampleCount(), nil
	}
	switch buf := buf.(type) {
	case *Buffer32:
		switch buf.ChannelCountValue {
		case 1:
			for i := 0; i < len(buf.Data); i++ {
				buf.Data[i] = float32(phase)
				phase += speed
				speed = (speed + target) * 0.5
			}
		case 2:
			for i := 0; i < len(buf.Data); i += 2 {
				buf.Data[i] = float32(phase)
				buf.Data[i+1] = buf.Data[i]
				phase += speed
				speed = (speed + target) * 0.5
			}
		default:
			return 0, errors.New("unsupported channel count")
		}
	default:
		return 0, errors.New("invalid buffer type")
	}
	gen.phase = phase
	gen.frequency = speed / inv

	return buf.SampleCount(), nil
}

func MonoInterface32(out Buffer, sample func() float32) error {
	switch out := out.(type) {
	case *Buffer32:
		switch out.ChannelCountValue {
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
			for i := 0; i < len(out.Data); i += out.ChannelCountValue {
				v := sample()
				for k := 0; k < out.ChannelCountValue; k++ {
					out.Data[i+k] = v
				}
			}
		}
		return nil
	default:
		return errors.New("unsupported buffer")
	}
}

func MonoDirect32(out *Buffer32, sample func() float32) {
	switch out.ChannelCountValue {
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
		for i := 0; i < len(out.Data); i += out.ChannelCountValue {
			v := sample()
			for k := 0; k < out.ChannelCountValue; k++ {
				out.Data[i+k] = v
			}
		}
	}
}
