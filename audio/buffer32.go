package audio

import "time"

type Buffer32 struct {
	Format
	Data []float32
}

func NewBuffer32(format Format, duration time.Duration) *Buffer32 {
	return NewBuffer32Samples(format, format.SamplesPerChannel(duration))
}

func NewBuffer32Samples(format Format, samplesPerChannel int) *Buffer32 {
	n := format.ChannelCount * samplesPerChannel
	return &Buffer32{
		Format: format,
		Data:   make([]float32, n, n),
	}
}

func (buffer *Buffer32) Clone() *Buffer32 {
	return &Buffer32{
		Format: buffer.Format,
		Data:   make([]float32, len(buffer.Data), len(buffer.Data)),
	}
}

func (buffer *Buffer32) Duration() time.Duration {
	return buffer.Format.BufferDuration(len(buffer.Data))
}

// TODO: get rid of this allocation
func (buffer *Buffer32) Slice(start, end int) *Buffer32 {
	return &Buffer32{
		Format: buffer.Format,
		Data:   buffer.Data[start:end],
	}
}

func (buffer *Buffer32) Process32(output *Buffer32) (int, error) {
	if Debug && buffer.Format != output.Format {
		return 0, ErrIncompatibleFormat
	}
	if len(buffer.Data) != len(output.Data) {
		return 0, ErrDifferentBufferSize
	}

	copy(output.Data, buffer.Data)

	return len(output.Data), nil
}

func (buffer *Buffer32) Process64(output *Buffer64) (int, error) {
	if Debug && buffer.Format != output.Format {
		return 0, ErrIncompatibleFormat
	}
	if len(buffer.Data) != len(output.Data) {
		return 0, ErrDifferentBufferSize
	}

	for i, v := range buffer.Data {
		output.Data[i] = float64(v)
	}

	return len(output.Data), nil
}
