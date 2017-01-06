package audio

import (
	"errors"
	"time"
)

var (
	ErrInvalidFrame = errors.New("invalid frame")
)

type Seeker interface {
	Position() time.Duration
	// Seek seeks to the nearest possible position and returns the final position
	Seek(offset time.Duration, whence int) (time.Duration, error)
}

// Anything implemeting Durationer has a known finite length
type Durationer interface {
	Duration() time.Duration
}

type Closer interface {
	Close() error
}
