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

func (reader *Reader) Process32(buf *audio.Buffer32) (int, error) {
	// only render as much as we should
	timeLeft := reader.Duration() - reader.Position()
	samplesPerChannel := buf.Format.SamplesPerChannel(timeLeft)
	samplesLeft := samplesPerChannel * buf.Format.ChannelCount
	if samplesLeft < len(buf.Data) {
		buf = buf.Slice(0, samplesLeft)
	}

	// current frequency
	head := int(reader.pos / reader.decay)
	pos := reader.pos - time.Duration(head)*reader.decay
	freq := reader.freqs[head]

	// params
	timeStep := buf.Format.BufferDuration(1)
	inv := 2.0 * math.Pi / float64(buf.SampleRate)
	speed := freq * inv
	phase := reader.phase

	generate.Mono32(buf, func() float32 {
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
		return float32(r)
	})

	reader.pos += buf.Duration()
	reader.phase = phase

	if reader.Position() >= reader.Duration() {
		return len(buf.Data), io.EOF
	}

	return len(buf.Data), nil
}

func (reader *Reader) Process64(buf *audio.Buffer64) (int, error) {
	panic("TODO")
	return 0, nil
}
