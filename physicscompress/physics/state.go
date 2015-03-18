package physics

import (
	"io"

	"github.com/egonelbre/exp/bit"
)

const HistorySize = 8

type State struct {
	History    [HistorySize]*Frame
	FrameIndex int
}

type Frame struct {
	Cubes []Cube
}

type Cube struct {
	Largest     int32
	A, B, C     int32
	X, Y, Z     int32
	Interacting int32
}

func NewFrame(size int) *Frame { return &Frame{make([]Cube, size)} }

func (frame *Frame) Assign(other *Frame) {
	frame.Cubes = append([]Cube{}, other.Cubes...)
}

func (frame *Frame) ReadFrom(r io.Reader) error {
	for i := range frame.Cubes {
		if err := frame.Cubes[i].ReadFrom(r); err != nil {
			return err
		}
	}
	return nil
}

func (frame *Frame) WriteTo(w io.Writer) error {
	for i := range frame.Cubes {
		if err := frame.Cubes[i].WriteTo(w); err != nil {
			return err
		}
	}
	return nil
}

func (a *Frame) Equals(b *Frame) bool {
	if len(a.Cubes) != len(b.Cubes) {
		return false
	}
	for i, ax := range a.Cubes {
		if ax != b.Cubes[i] {
			return false
		}
	}
	return true
}

// Cube utilities
func (cube *Cube) ReadFrom(r io.Reader) error {
	return bit.Read(r,
		&cube.Largest,
		&cube.A, &cube.B, &cube.C,
		&cube.X, &cube.Y, &cube.Z,
		&cube.Interacting,
	)
}

func (cube *Cube) WriteTo(w io.Writer) error {
	return bit.Write(w,
		&cube.Largest,
		&cube.A, &cube.B, &cube.C,
		&cube.X, &cube.Y, &cube.Z,
		&cube.Interacting,
	)
}

func NewState(size int) *State {
	s := &State{}
	for i := range s.History {
		s.History[i] = NewFrame(size)
	}
	s.FrameIndex = -1
	return s
}

func (s *State) IncFrame() {
	s.FrameIndex += 1
}

func (s *State) ReadNext(r io.Reader) error {
	s.FrameIndex += 1
	return s.Current().ReadFrom(r)
}

func (s *State) Current() *Frame  { return s.Prev(0) }
func (s *State) Baseline() *Frame { return s.Prev(6) }

func (s *State) Prev(i int) *Frame {
	return s.History[(s.FrameIndex-i+HistorySize)%HistorySize]
}
