package physics

import (
	"bytes"
	"compress/flate"
	"flag"
	"io"

	"github.com/egonelbre/exp/bit"
	"github.com/egonelbre/exp/bit/expgolomb"
)

var flagFlate = flag.Bool("flate", true, "use flate compression")

func encode32(v int32) uint64 { return uint64(bit.ZEncode(int64(v))) }
func decode32(v uint64) int32 { return int32(bit.ZDecode(v)) }

func write(w *bit.Writer, v int32, bits uint) { w.WriteBits(uint64(v), bits) }
func read(r *bit.Reader, bits uint) int32     { return int32(r.ReadBits(bits)) }

func write32(w *bit.Writer, v int32, bits uint) { w.WriteBits(bit.ZEncode(int64(v)), bits) }
func read32(r *bit.Reader, bits uint) int32     { return int32(bit.ZDecode(r.ReadBits(bits))) }

func maxbits(vs ...int32) (r uint) {
	r = 0
	for _, x := range vs {
		if x != 0 {
			bits := bit.ScanRight(bit.ZEncode(int64(x))) + 1
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
	var next io.Writer

	var buf bytes.Buffer
	next = &buf

	var pack *flate.Writer
	if *flagFlate {
		pack, _ = flate.NewWriter(next, flate.BestCompression)
		next = pack
	}

	w := bit.NewWriter(next)

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

	if err := w.Close(); err != nil {
		panic(err)
	}

	if pack != nil {
		pack.Close()
	}
	return buf.Bytes()
}

func (s *State) Decode(snapshot []byte) {
	var next io.Reader
	buf := bytes.NewBuffer(snapshot)
	next = buf
	if *flagFlate {
		next = flate.NewReader(next)
	}
	r := bit.NewReader(next)

	baseline := s.Baseline()
	current := s.Current()
	current.Assign(baseline)

	for i := range current.Cubes {
		base := baseline.Cubes[i]
		cube := &current.Cubes[i]

		isSame := r.ReadBit()
		if isSame == 0 {
			continue
		}

		bits := uint(expgolomb.ReadInt(r) + 9)

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
