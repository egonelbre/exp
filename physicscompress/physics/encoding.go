package physics

import (
	"bytes"

	"github.com/egonelbre/exp/bit"
)

func encode32(v int32) uint64 { return uint64(bit.AbsEncode(int64(v))) }
func decode32(v uint64) int32 { return int32(bit.AbsDecode(v)) }

func write(w *bit.Writer, v int32, bits uint64) {
	w.WriteBits(uint64(v), bits)
}
func read(r *bit.Reader, bits uint64) int32 {
	v, _ := r.ReadBits(bits)
	return int32(v)
}

func write32(w *bit.Writer, v int32, bits uint64) {
	w.WriteBits(bit.AbsEncode(int64(v)), bits)
}

func read32(r *bit.Reader, bits uint64) int32 {
	v, _ := r.ReadBits(bits)
	return int32(bit.AbsDecode(v))
}

func (s *State) Encode() []byte {
	var buf bytes.Buffer
	w := bit.NewWriter(&buf)

	current := s.Current()

	for _, cube := range current.Cubes {
		write(w, cube.Interacting, 1)
		write(w, cube.Largest, 2)

		write32(w, cube.A, 32)
		write32(w, cube.B, 32)
		write32(w, cube.C, 32)

		write32(w, cube.X, 32)
		write32(w, cube.Y, 32)
		write32(w, cube.Z, 32)
	}

	w.Close()
	return buf.Bytes()
}

func (s *State) Decode(snapshot []byte) {
	buf := bytes.NewBuffer(snapshot)
	r := bit.NewReader(buf)

	current := s.Current()

	for i := range current.Cubes {
		cube := &current.Cubes[i]

		cube.Interacting = read(r, 1)
		cube.Largest = read(r, 2)

		cube.A = read32(r, 32)
		cube.B = read32(r, 32)
		cube.C = read32(r, 32)

		cube.X = read32(r, 32)
		cube.Y = read32(r, 32)
		cube.Z = read32(r, 32)
	}
}
