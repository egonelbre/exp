package audio

const int16scale = 1.0 / 0x8000

type BufferInt16 struct {
	Format
	Data []int16
}

func (buffer *BufferInt16) Process32(output *Buffer32) error {
	if Debug && buffer.Format != output.Format {
		return ErrIncompatibleFormat
	}
	if len(buffer.Data) != len(output.Data) {
		return ErrDifferentBufferSize
	}

	for i, v := range buffer.Data {
		output.Data[i] = float32(v) * int16scale
	}

	return nil
}

func (buffer *BufferInt16) Process64(output *Buffer64) error {
	if Debug && buffer.Format != output.Format {
		return ErrIncompatibleFormat
	}
	if len(buffer.Data) != len(output.Data) {
		return ErrDifferentBufferSize
	}

	for i, v := range buffer.Data {
		output.Data[i] = float64(v) * int16scale
	}

	return nil
}
