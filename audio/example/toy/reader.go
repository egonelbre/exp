package main

import (
	"io"
	"math"
	"time"

	"github.com/egonelbre/exp/audio"
	"github.com/egonelbre/exp/audio/generate"
)

type Reader struct {
	freqs []float64
	pos   time.Duration
	decay time.Duration
	phase float64
}

func NewReader(freqs ...float64) *Reader {
	return &Reader{
		freqs: freqs,
		pos:   0,
		decay: time.Second,
		phase: 0.0,
	}
}

func (reader *Reader) Duration() time.Duration {
	return time.Duration(len(reader.freqs)) * reader.decay
}

func (reader *Reader) Position() time.Duration {
	return reader.pos
}

/*
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
*/

func (reader *Reader) Read(buf audio.Buffer) (int, error) {
	fmt := audio.Format{buf.SampleRate(), buf.ChannelCount()}
	buf = buf.ShallowCopy()

	// only render as much as we should
	timeLeft := reader.Duration() - reader.Position()
	frameCount := fmt.FrameCount(timeLeft)
	if frameCount < buf.FrameCount() {
		buf.Slice(0, frameCount)
	}

	// current frequency
	head := int(reader.pos / reader.decay)
	pos := reader.pos - time.Duration(head)*reader.decay
	freq := reader.freqs[head]

	// params
	timeStep := time.Duration(float64(time.Second) / float64(fmt.SampleRate))
	inv := 2.0 * math.Pi / float64(buf.SampleRate())
	speed := freq * inv
	phase := reader.phase

	generate.MonoF64(buf, func() float64 {
		r := math.Cos(phase)
		phase += speed
		pos += timeStep
		if pos > reader.decay {
			pos = 0
			head++
			if head < len(reader.freqs) {
				freq = reader.freqs[head]
				speed = freq * inv
			}
		}
		return r
	})

	reader.pos += buf.Duration()
	reader.phase = phase

	if reader.Position() >= reader.Duration() {
		return buf.FrameCount(), io.EOF
	}

	return buf.FrameCount(), nil
}
