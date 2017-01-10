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

func (format Format) BufferDuration(frameCount int) time.Duration {
	seconds := float64(frameCount) / float64(format.SampleRate)
	return time.Duration(float64(time.Second) * seconds)
}

func (format Format) FrameCount(duration time.Duration) int {
	seconds := duration.Seconds()
	return int(math.Ceil(seconds * float64(format.SampleRate)))
}
