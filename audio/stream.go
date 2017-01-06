package audio

import (
	"errors"
	"time"
)

var (
	ErrInvalidFrame = errors.New("invalid frame")
)

// Frame is a chunk of seekable audio data in an unknown format
// that knows how to convert to/from Buffer32/Buffer64
type Frame interface {
	// Offset returns the stream start time
	Offset() time.Duration
	// Position returns the current read/write position inside this frame
	Position() time.Duration
	// Duration returns the total time in seconds
	Duration() time.Duration
	// Done should return true, when further ProcessX calls do not work
	Done() bool
	// Process32/Process64 either read/write data from buffer
	// depending on whether it was returned by StreamReader or StreamWriter
	Node
	// Seek seeks to the nearest sample possible in the frame
	Seeker
}

type Seeker interface {
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

// StreamReader implements reading Frame-s
type StreamReader interface {
	// Position returns the position of the next Frame to be read
	Position() time.Duration
	// NewFrame returns an empty frame compatible with StreamReader
	NewFrame() Frame
	// Read reads the next frame, it may change Frame
	Read(frame Frame) error
}

// StreamReadSeeker implements a seekable StreamReader
type StreamReadSeeker interface {
	StreamReader
	Seeker
}

// StreamWriter implements writing Frames
type StreamWriter interface {
	// Position returns the position of the next Frame to be written
	Position() time.Duration
	// NewFrame returns an empty frame compatible with StreamWriter
	NewFrame() Frame
	// Write writes the next Frame, it may change Frame
	//       it must reset/replace frame with an empty frame
	Write(frame Frame) error
}

//TODO: is there any format that supports this? I would guess, yes
// type StreamWriteSeeker interface {
// 	StreamWriter
// 	Seeker
// }
