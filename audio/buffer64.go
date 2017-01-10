package audio

import "time"

type BufferF64 struct {
	Format Format
	Data   []float64
}

func NewBufferF64(format Format, duration time.Duration) *BufferF64 {
	return NewBufferF64Frames(format, format.FrameCount(duration))
}

func NewBufferF64Frames(format Format, frames int) *BufferF64 {
	n := format.ChannelCount * frames
	return &BufferF64{
		Format: format,
		Data:   make([]float64, n, n),
	}
}

func (b *BufferF64) SampleRate() int   { return b.Format.SampleRate }
func (b *BufferF64) ChannelCount() int { return b.Format.ChannelCount }

func (b *BufferF64) Empty() bool { return len(b.Data) == 0 }

func (b *BufferF64) FrameCount() int {
	return len(b.Data) / b.ChannelCount()
}

func (b *BufferF64) Duration() time.Duration {
	return time.Duration(int(time.Second) * b.FrameCount() / b.SampleRate())
}

func (b *BufferF64) ShallowCopy() Buffer {
	return &BufferF64{
		Format: b.Format,
		Data:   b.Data,
	}
}

func (b *BufferF64) DeepCopy() Buffer {
	x := &BufferF64{
		Format: b.Format,
		Data:   make([]float64, len(b.Data), len(b.Data)),
	}
	copy(x.Data, b.Data)
	return x
}

func (b *BufferF64) Slice(lowFrame, highFrame int) {
	b.Data = b.Data[lowFrame*b.ChannelCount() : highFrame*b.ChannelCount()]
}
