package physics

import (
	"bytes"
	"compress/flate"

	"github.com/egonelbre/exp/bit"
	"github.com/egonelbre/exp/bit/expgolomb"
)

func encode32(v int32) uint64 { return uint64(bit.AbsEncode(int64(v))) }
func decode32(v uint64) int32 { return int32(bit.AbsDecode(v)) }

func write(w *bit.Writer, v int32, bits uint) {
	w.WriteBits(uint64(v), bits)
}
func read(r *bit.Reader, bits uint) int32 {
	v, _ := r.ReadBits(bits)
	return int32(v)
}

func write32(w *bit.Writer, v int32, bits uint) {
	w.WriteBits(bit.AbsEncode(int64(v)), bits)
}

func read32(r *bit.Reader, bits uint) int32 {
	v, _ := r.ReadBits(bits)
	return int32(bit.AbsDecode(v))
}

func maxbits(vs ...int32) (r uint) {
	r = 0
	for _, x := range vs {
		if x != 0 {
			bits := bit.ScanRight(bit.AbsEncode(int64(x))) + 1
			if r < bits {
				r = bits
			}
		}
	}
	return
}

func deltabits(base, cube *Cube) uint {
	return maxbits(
		cube.Interacting^base.Interacting,
		cube.Largest^base.Largest,
		cube.A^base.A,
		cube.B^base.B,
		cube.C^base.C,
		cube.X^base.X,
		cube.Y^base.Y,
		cube.Z^base.Z,
	)
}

func (s *State) Encode() []byte {
	var buf bytes.Buffer

	pack, _ := flate.NewWriter(&buf, flate.DefaultCompression)
	w := bit.NewWriter(pack)

	baseline := s.Baseline()
	current := s.Current()

	for i, cube := range current.Cubes {
		base := baseline.Cubes[i]

		if cube == base {
			w.WriteBit(0)
			continue
		}
		w.WriteBit(1)

		bits := deltabits(&base, &cube)
		expgolomb.WriteInt(w, int(bits)-9)

		write(w, cube.Interacting^base.Interacting, 1)
		write(w, cube.Largest^base.Largest, 2)

		write32(w, cube.A^base.A, bits)
		write32(w, cube.B^base.B, bits)
		write32(w, cube.C^base.C, bits)
		write32(w, cube.X^base.X, bits)
		write32(w, cube.Y^base.Y, bits)
		write32(w, cube.Z^base.Z, bits)
	}

	w.Close()
	pack.Close()
	return buf.Bytes()
}

func (s *State) Decode(snapshot []byte) {
	buf := bytes.NewBuffer(snapshot)
	pack := flate.NewReader(buf)
	r := bit.NewReader(pack)

	baseline := s.Baseline()
	current := s.Current()
	current.Assign(baseline)

	for i := range current.Cubes {
		base := baseline.Cubes[i]
		cube := &current.Cubes[i]

		isSame, _ := r.ReadBit()
		if isSame == 0 {
			continue
		}

		x, _ := expgolomb.ReadInt(r)
		bits := uint(x + 9)

		cube.Interacting = read(r, 1) ^ base.Interacting
		cube.Largest = read(r, 2) ^ base.Largest

		cube.A = read32(r, bits) ^ base.A
		cube.B = read32(r, bits) ^ base.B
		cube.C = read32(r, bits) ^ base.C

		cube.X = read32(r, bits) ^ base.X
		cube.Y = read32(r, bits) ^ base.Y
		cube.Z = read32(r, bits) ^ base.Z
	}
}
