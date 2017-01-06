package main

import (
	"errors"
	"io"
	"math"
	"time"

	"github.com/egonelbre/exp/audio"
	"github.com/egonelbre/exp/audio/generate"
)

type Reader struct {
	freqs []float64
	head  int

	frameDuration time.Duration
}

func NewReader(freqs ...float64) *Reader {
	return &Reader{
		freqs: freqs,
		head:  0,

		frameDuration: time.Second,
	}
}

func (reader *Reader) NewFrame() audio.Frame {
	return &Frame{0, reader.frameDuration, reader.frameDuration, 0}
}

func (reader *Reader) Duration() time.Duration {
	return time.Duration(len(reader.freqs)) * reader.frameDuration
}

func (reader *Reader) Position() time.Duration {
	return time.Duration(reader.head) * reader.frameDuration
}

func (reader *Reader) Seek(offset time.Duration, whence int) (time.Duration, error) {
	//TODO: handle whence parameter
	var err error

	head := (int)(offset / reader.frameDuration)
	if head < 0 {
		head = 0
	} else if head > len(reader.freqs) {
		head = len(reader.freqs)
		//TODO: use the same conventions as io.Seeker
		err = io.EOF
	}
	reader.head = head

	return reader.Position(), err
}

func (reader *Reader) Read(frame audio.Frame) error {
	f, ok := frame.(*Frame)
	if !ok {
		return audio.ErrInvalidFrame
	}

	if reader.head >= len(reader.freqs) {
		return io.EOF
	}

	f.offset = reader.Position()
	f.position = 0
	f.duration = reader.frameDuration
	f.frequency = reader.freqs[reader.head]

	reader.head++
	if reader.head >= len(reader.freqs) {
		return io.EOF
	}

	return nil
}

type Frame struct {
	offset    time.Duration
	position  time.Duration
	duration  time.Duration
	frequency float64
}

func (frame *Frame) Offset() time.Duration   { return frame.offset }
func (frame *Frame) Position() time.Duration { return frame.position }
func (frame *Frame) Duration() time.Duration { return frame.duration }

func (frame *Frame) Seek(offset time.Duration, whence int) (time.Duration, error) {
	// TODO: implement whence
	frame.position = offset
	return offset, errors.New("seek not supported")
}

func (frame *Frame) Done() bool {
	return frame.position >= frame.duration
}

func (frame *Frame) Process32(buf *audio.Buffer32) (int, error) {
	inv := 2.0 * math.Pi / float64(buf.SampleRate)
	speed := frame.frequency * inv
	phase := speed * float64(buf.Format.SamplesPerChannel(frame.position))

	samplesPerChannel := buf.Format.SamplesPerChannel(frame.duration - frame.position)
	samplesLeft := samplesPerChannel * buf.Format.ChannelCount
	if samplesLeft < len(buf.Data) {
		buf = buf.Slice(0, samplesLeft)
	}

	generate.Mono32(buf, func() float32 {
		r := math.Cos(phase)
		phase += speed
		return float32(r)
	})

	frame.position += buf.Duration()
	return len(buf.Data), nil
}

func (frame *Frame) Process64(buf *audio.Buffer64) (int, error) {
	inv := 2.0 * math.Pi / float64(buf.SampleRate)
	speed := frame.frequency * inv
	phase := speed * float64(buf.Format.SamplesPerChannel(frame.position))

	samplesPerChannel := buf.Format.SamplesPerChannel(frame.duration - frame.position)
	samplesLeft := samplesPerChannel * buf.Format.ChannelCount
	if samplesLeft < len(buf.Data) {
		buf = buf.Slice(0, samplesLeft)
	}

	generate.Mono64(buf, func() float64 {
		r := math.Cos(phase)
		phase += speed
		return r
	})
	frame.position += buf.Duration()
	return len(buf.Data), nil
}
