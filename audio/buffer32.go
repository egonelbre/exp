package audio

import "time"

type BufferF32 struct {
	format Format
	offset uint32
	frames uint32
	stride uint32
	data   []float32
}

func NewBufferF32(format Format, duration time.Duration) *BufferF32 {
	return NewBufferF32Frames(format, format.FrameCount(duration))
}

func NewBufferF32Frames(format Format, frames int) *BufferF32 {
	samples := format.ChannelCount * frames
	return &BufferF32{
		format: format,
		offset: 0,
		stride: uint32(frames),
		frames: uint32(frames),
		data:   make([]float32, samples, samples),
	}
}

func (b *BufferF32) Channel(index int) []float32 {
	start := int(b.offset) + index*int(b.stride)
	return b.data[start : start+int(b.frames)]
}

func (b *BufferF32) SampleRate() int   { return b.format.SampleRate }
func (b *BufferF32) ChannelCount() int { return b.format.ChannelCount }

func (b *BufferF32) Empty() bool     { return b.frames == 0 }
func (b *BufferF32) FrameCount() int { return int(b.frames) }

func (b *BufferF32) Duration() time.Duration {
	return time.Duration(int(time.Second) * b.FrameCount() / b.SampleRate())
}

func (b *BufferF32) ShallowCopy() Buffer {
	x := *b
	return &x
}

func (b *BufferF32) DeepCopy() Buffer {
	x := *b
	x.data = make([]float32, len(b.data), len(b.data))
	copy(x.data, b.data)
	return &x
}

func (b *BufferF32) Slice(low, high int) {
	b.offset += uint32(low)
	b.frames = uint32(high - low)
}

func (b *BufferF32) CutLeading(low int) {
	b.offset += uint32(low)
	b.frames -= uint32(low)
}
