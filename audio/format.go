package audio

import (
	"errors"
	"math"
	"time"
)

type Format struct {
	SampleRate   int
	ChannelCount int
}

var (
	ErrIncompatibleFormat  = errors.New("incompatible formats")
	ErrDifferentBufferSize = errors.New("different buffer sizes")
)

func (format Format) BufferDuration(n int) time.Duration {
	samplesPerChannel := n / format.ChannelCount
	seconds := float64(samplesPerChannel) / float64(format.SampleRate)
	return time.Duration(float64(time.Second) * seconds)
}

func (format Format) SamplesPerChannel(duration time.Duration) int {
	seconds := duration.Seconds()
	return int(math.Ceil(seconds * float64(format.SampleRate)))
}
