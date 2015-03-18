package physics

import (
	"bytes"
	"compress/flate"

	"github.com/egonelbre/exp/bit"
	"github.com/egonelbre/exp/bit/expgolomb"
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

type bits struct {
	ABC uint64
	XYZ uint64
}

func (b *bits) WriteTo(w *bit.Writer) {
	expgolomb.WriteInt(w, int(b.ABC)-0xA)
	expgolomb.WriteInt(w, int(b.XYZ)-0xA)
}

func (b *bits) ReadFrom(r *bit.Reader) {
	abc, _ := expgolomb.ReadInt(r)
	xyz, _ := expgolomb.ReadInt(r)

	b.ABC = uint64(abc + 0xA)
	b.XYZ = uint64(xyz + 0xA)
}

func (b *bits) Update(baseline, current *Frame) {
	for i, cube := range current.Cubes {
		base := baseline.Cubes[i]
		b.ABC = maxbits(b.ABC, cube.A^base.A)
		b.ABC = maxbits(b.ABC, cube.B^base.B)
		b.ABC = maxbits(b.ABC, cube.C^base.C)
		b.XYZ = maxbits(b.XYZ, cube.X^base.X)
		b.XYZ = maxbits(b.XYZ, cube.Y^base.Y)
		b.XYZ = maxbits(b.XYZ, cube.Z^base.Z)
	}
}

func maxbits(a uint64, b int32) uint64 {
	x := bit.AbsEncode(int64(b))
	w := bit.ScanRight(x) + 1
	if a < w {
		return w
	}
	return a
}

func (s *State) Encode() []byte {
	var buf bytes.Buffer

	pack, _ := flate.NewWriter(&buf, flate.DefaultCompression)
	w := bit.NewWriter(pack)

	baseline := s.Baseline()
	current := s.Current()

	var bits bits
	bits.Update(baseline, current)
	bits.WriteTo(w)

	for i, cube := range current.Cubes {
		base := baseline.Cubes[i]

		write(w, cube.Interacting^base.Interacting, 1)
		write(w, cube.Largest^base.Largest, 2)

		write32(w, cube.A^base.A, bits.ABC)
		write32(w, cube.B^base.B, bits.ABC)
		write32(w, cube.C^base.C, bits.ABC)

		write32(w, cube.X^base.X, bits.XYZ)
		write32(w, cube.Y^base.Y, bits.XYZ)
		write32(w, cube.Z^base.Z, bits.XYZ)
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

	var bits bits
	bits.ReadFrom(r)

	for i := range current.Cubes {
		base := baseline.Cubes[i]
		cube := &current.Cubes[i]

		cube.Interacting = read(r, 1) ^ base.Interacting
		cube.Largest = read(r, 2) ^ base.Largest

		cube.A = read32(r, bits.ABC) ^ base.A
		cube.B = read32(r, bits.ABC) ^ base.B
		cube.C = read32(r, bits.ABC) ^ base.C

		cube.X = read32(r, bits.XYZ) ^ base.X
		cube.Y = read32(r, bits.XYZ) ^ base.Y
		cube.Z = read32(r, bits.XYZ) ^ base.Z
	}
}
