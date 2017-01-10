package audio

/* TODO:
const int16scale = 1.0 / 0x8000

type BufferInt16 struct {
	Format
	Data []int16
}

func (buffer *BufferInt16) ConvertToF32(output *Buffer32) (int, error) {
	if Debug && buffer.Format != output.Format {
		return 0, ErrIncompatibleFormat
	}
	if len(buffer.Data) != len(output.Data) {
		return 0, ErrDifferentBufferSize
	}

	for i, v := range buffer.Data {
		output.Data[i] = float32(v) * int16scale
	}

	return len(output.Data), nil
}

func (buffer *BufferInt16) ConvertToF64(output *Buffer64) (int, error) {
	if Debug && buffer.Format != output.Format {
		return 0, ErrIncompatibleFormat
	}
	if len(buffer.Data) != len(output.Data) {
		return 0, ErrDifferentBufferSize
	}

	for i, v := range buffer.Data {
		output.Data[i] = float64(v) * int16scale
	}

	return len(output.Data), nil
}
*/
