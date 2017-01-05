package audio

import (
	"errors"
	"math"
	"time"
)

type Format struct {
	SampleRate int
	Channels   int
}

type Buffer32 struct {
	Format
	Data []float32
}

type Buffer64 struct {
	Format
	Data []float64
}

func NewBuffer32(format Format, duration time.Duration) *Buffer32 {
	samples := math.Ceil(duration.Seconds() / float64(format.SampleRate))
	return NewBuffer32Samples(format, int(samples))
}

func NewBuffer64(format Format, duration time.Duration) *Buffer64 {
	samples := math.Ceil(duration.Seconds() / float64(format.SampleRate))
	return NewBuffer64Samples(format, int(samples))
}

func NewBuffer32Samples(format Format, sampleCount int) *Buffer32 {
	n := format.Channels * sampleCount
	return &Buffer32{
		Format: format,
		Data:   make([]float32, n, n),
	}
}

func NewBuffer64Samples(format Format, sampleCount int) *Buffer64 {
	n := format.Channels * sampleCount
	return &Buffer64{
		Format: format,
		Data:   make([]float64, n, n),
	}
}

var (
	ErrIncompatibleFormat  = errors.New("incompatible formats")
	ErrDifferentBufferSize = errors.New("different buffer sizes")
)

func (buffer *Buffer32) Process32(output *Buffer32) error {
	if Debug && buffer.Format != output.Format {
		return ErrIncompatibleFormat
	}
	if len(buffer.Data) != len(output.Data) {
		return ErrDifferentBufferSize
	}

	copy(output.Data, buffer.Data)

	return nil
}

func (buffer *Buffer32) Process64(output *Buffer64) error {
	if Debug && buffer.Format != output.Format {
		return ErrIncompatibleFormat
	}
	if len(buffer.Data) != len(output.Data) {
		return ErrDifferentBufferSize
	}

	for i, v := range buffer.Data {
		output.Data[i] = float64(v)
	}

	return nil
}

func (buffer *Buffer64) Process32(output *Buffer32) error {
	if Debug && buffer.Format != output.Format {
		return ErrIncompatibleFormat
	}
	if len(buffer.Data) != len(output.Data) {
		return ErrDifferentBufferSize
	}

	for i, v := range buffer.Data {
		output.Data[i] = float32(v)
	}

	return nil
}

func (buffer *Buffer64) Process64(output *Buffer64) error {
	if Debug && buffer.Format != output.Format {
		return ErrIncompatibleFormat
	}
	if len(buffer.Data) != len(output.Data) {
		return ErrDifferentBufferSize
	}

	copy(output.Data, buffer.Data)

	return nil
}
