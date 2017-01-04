package audio

import "errors"

type Format struct {
	SampleRate uint32
	Channels   int16
}

type Buffer32 struct {
	Format
	Data []float32
}

type Buffer64 struct {
	Format
	Data []float64
}

func NewBuffer32(format Format, sampleCount int) *Buffer32 {
	return &Buffer32{
		Format: format,
		Data:   make([]float32, sampleCount, sampleCount),
	}
}

func NewBuffer64(format Format, sampleCount int) *Buffer64 {
	return &Buffer64{
		Format: format,
		Data:   make([]float64, sampleCount, sampleCount),
	}
}

var (
	ErrIncompatibleFormat  = errors.New("incompatible formats")
	ErrDifferentBufferSize = errors.New("different buffer sizes")
)

func (format Format) Same(other Format) error {
	if Debug {
		if !format.Equals(other) {
			return ErrIncompatibleFormat
		}
	}
	return nil
}

func (format Format) Equals(other Format) bool {
	return format == other
}

func (buffer *Buffer32) Process32(output *Buffer32) error {
	if err := buffer.Format.Same(output.Format); err != nil {
		return err
	}
	if len(buffer.Data) != len(output.Data) {
		return ErrDifferentBufferSize
	}

	copy(output.Data, buffer.Data)

	return nil
}

func (buffer *Buffer32) Process64(output *Buffer64) error {
	if err := buffer.Format.Same(output.Format); err != nil {
		return err
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
	if err := buffer.Format.Same(output.Format); err != nil {
		return err
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
	if err := buffer.Format.Same(output.Format); err != nil {
		return err
	}
	if len(buffer.Data) != len(output.Data) {
		return ErrDifferentBufferSize
	}

	copy(output.Data, buffer.Data)

	return nil
}
