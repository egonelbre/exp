package audio

import "time"

type Buffer64 struct {
	Format
	Data []float64
}

func NewBuffer64(format Format, duration time.Duration) *Buffer64 {
	return NewBuffer64Samples(format, format.SamplesPerChannel(duration))
}

func NewBuffer64Samples(format Format, samplesPerChannel int) *Buffer64 {
	n := format.ChannelCount * samplesPerChannel
	return &Buffer64{
		Format: format,
		Data:   make([]float64, n, n),
	}
}

func (buffer *Buffer64) Clone() *Buffer64 {
	return &Buffer64{
		Format: buffer.Format,
		Data:   make([]float64, len(buffer.Data), len(buffer.Data)),
	}
}

func (buffer *Buffer64) Duration() time.Duration {
	return buffer.Format.BufferDuration(len(buffer.Data))
}

// TODO: get rid of this allocation
func (buffer *Buffer64) Slice(start, end int) *Buffer64 {
	return &Buffer64{
		Format: buffer.Format,
		Data:   buffer.Data[start:end],
	}
}

func (buffer *Buffer64) Process32(output *Buffer32) (int, error) {
	if Debug && buffer.Format != output.Format {
		return 0, ErrIncompatibleFormat
	}
	if len(buffer.Data) != len(output.Data) {
		return 0, ErrDifferentBufferSize
	}

	for i, v := range buffer.Data {
		output.Data[i] = float32(v)
	}

	return len(output.Data), nil
}

func (buffer *Buffer64) Process64(output *Buffer64) (int, error) {
	if Debug && buffer.Format != output.Format {
		return 0, ErrIncompatibleFormat
	}
	if len(buffer.Data) != len(output.Data) {
		return 0, ErrDifferentBufferSize
	}

	copy(output.Data, buffer.Data)
	return len(output.Data), nil
}
