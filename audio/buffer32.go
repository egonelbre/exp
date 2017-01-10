package audio

import "time"

type BufferF32 struct {
	Format Format
	Data   []float32
}

func NewBufferF32(format Format, duration time.Duration) *BufferF32 {
	return NewBufferF32Frames(format, format.FrameCount(duration))
}

func NewBufferF32Frames(format Format, frames int) *BufferF32 {
	n := format.ChannelCount * frames
	return &BufferF32{
		Format: format,
		Data:   make([]float32, n, n),
	}
}

func (b *BufferF32) SampleRate() int   { return b.Format.SampleRate }
func (b *BufferF32) ChannelCount() int { return b.Format.ChannelCount }

func (b *BufferF32) Empty() bool { return len(b.Data) == 0 }

func (b *BufferF32) FrameCount() int {
	return len(b.Data) / b.ChannelCount()
}

func (b *BufferF32) Duration() time.Duration {
	return time.Duration(int(time.Second) * b.FrameCount() / b.SampleRate())
}

func (b *BufferF32) ShallowCopy() Buffer {
	return &BufferF32{
		Format: b.Format,
		Data:   b.Data,
	}
}

func (b *BufferF32) DeepCopy() Buffer {
	x := &BufferF32{
		Format: b.Format,
		Data:   make([]float32, len(b.Data), len(b.Data)),
	}
	copy(x.Data, b.Data)
	return x
}

func (b *BufferF32) Slice(lowFrame, highFrame int) {
	b.Data = b.Data[lowFrame*b.ChannelCount() : highFrame*b.ChannelCount()]
}
