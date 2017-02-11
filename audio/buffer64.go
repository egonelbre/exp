package audio

import "time"

type BufferF64 struct {
	format Format
	offset uint32
	frames uint32
	stride uint32
	data   []float64
}

func NewBufferF64(format Format, duration time.Duration) *BufferF64 {
	return NewBufferF64Frames(format, format.FrameCount(duration))
}

func NewBufferF64Frames(format Format, frames int) *BufferF64 {
	n := format.ChannelCount * frames
	return &BufferF64{
		format: format,
		offset: 0,
		stride: uint32(frames),
		frames: uint32(frames),
		data:   make([]float64, n, n),
	}
}

func (b *BufferF64) Channel(index int) []float64 {
	start := int(b.offset) + index*int(b.stride)
	return b.data[start : start+int(b.frames)]
}

func (b *BufferF64) SampleRate() int   { return b.format.SampleRate }
func (b *BufferF64) ChannelCount() int { return b.format.ChannelCount }

func (b *BufferF64) Empty() bool     { return b.frames == 0 }
func (b *BufferF64) FrameCount() int { return int(b.frames) }

func (b *BufferF64) Duration() time.Duration {
	return time.Duration(int(time.Second) * b.FrameCount() / b.SampleRate())
}

func (b *BufferF64) ShallowCopy() Buffer {
	x := *b
	return &x
}

func (b *BufferF64) DeepCopy() Buffer {
	x := *b
	x.data = make([]float64, len(b.data), len(b.data))
	copy(x.data, b.data)
	return &x
}

func (b *BufferF64) Slice(low, high int) {
	b.offset += uint32(low)
	b.frames = uint32(high - low)
}

func (b *BufferF64) CutLeading(low int) {
	b.offset += uint32(low)
	b.frames -= uint32(low)
}
