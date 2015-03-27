package physics

import (
	"encoding/gob"
	"io"
	"math"

	"github.com/egonelbre/exp/bit"
)

const (
	HistorySize = 16

	UnitsPerMeter = 512
	UnitsPerQuat  = 1024

	PositionUnit = 1.0 / UnitsPerMeter
	QuatUnit     = 1.0 / UnitsPerQuat
)

type State struct {
	History    [HistorySize]*Frame
	FrameIndex int
}

type Frame struct {
	Cubes []Cube
}

type Rotation struct{ X, Y, Z, W float32 }

func same32(delta, precision float32) bool {
	if delta < 0 {
		delta = -delta
	}
	return delta <= precision
}

func (a Rotation) Equals(b Rotation) bool {
	return same32(a.X-b.X, QuatUnit) &&
		same32(a.Y-b.Y, QuatUnit) &&
		same32(a.Z-b.Z, QuatUnit) &&
		same32(a.W-b.W, QuatUnit)
}

type Position struct{ X, Y, Z float32 }

func (a Position) Sub(b Position) Position {
	return Position{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func (p Position) Len() float32 {
	return float32(math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z)))
}

func (a Position) Equals(b Position) bool {
	return same32(a.X-b.X, PositionUnit) &&
		same32(a.Y-b.Y, PositionUnit) &&
		same32(a.Z-b.Z, PositionUnit)
}

type Cube struct {
	Pos Position
	Rot Rotation

	Interacting bool
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
		bx := b.Cubes[i]
		switch {
		case ax.Interacting != bx.Interacting:
			return false
		case !ax.Pos.Equals(bx.Pos):
			return false
		case !ax.Rot.Equals(bx.Rot):
			return false
		}
	}
	return true
}

// Cube utilities
func (cube *Cube) ReadFrom(r io.Reader) error {
	interacting := int32(0)
	err := bit.Read(r,
		&cube.Rot.X, &cube.Rot.Y, &cube.Rot.Z, &cube.Rot.W,
		&cube.Pos.X, &cube.Pos.Y, &cube.Pos.Z,
		&interacting,
	)
	cube.Interacting = interacting == 1
	return err
}

func (cube *Cube) WriteTo(w io.Writer) error {
	interacting := int32(0)
	if cube.Interacting {
		interacting = 1
	}
	return bit.Write(w,
		&cube.Rot.X, &cube.Rot.Y, &cube.Rot.Z, &cube.Rot.W,
		&cube.Pos.X, &cube.Pos.Y, &cube.Pos.Z,
		&interacting,
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
func (s *State) Historic() *Frame { return s.Prev(8) }

func (s *State) Prev(i int) *Frame {
	return s.History[(s.FrameIndex-i+HistorySize)%HistorySize]
}

func init() {
	gob.Register([]Cube{})
}
