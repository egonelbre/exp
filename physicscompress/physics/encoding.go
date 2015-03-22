package physics

import (
	"flag"
	"fmt"

	"github.com/egonelbre/exp/bit"
	"github.com/egonelbre/exp/coder/arith"
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

// encodes zeros well
func mZeros() arith.Model {
	return &arith.Shift2{P0: 0x500, I0: 0x1, P1: 0x150, I1: 0x5}
}

var sameness []bool

func (s *State) Encode() []byte {
	enc := arith.NewEncoder()

	baseline := s.Baseline().Cubes
	current := s.Current().Cubes

	mzeros := mZeros()
	same := make([]bool, len(current))

	items := make([]int, 0, len(current))
	for i, cube := range current {
		base := baseline[i]
		same[i] = cube == base
		if same[i] {
			mzeros.Encode(enc, 0)
		} else {
			mzeros.Encode(enc, 1)
			items = append(items, i)
		}
	}
	sameness = same
	_ = items

	for _, i := range items {
		cube := &current[i]
		mzeros.Encode(enc, uint(cube.Interacting^1))
	}

	for _, i := range items {
		cube, base := &current[i], &baseline[i]
		v := uint(cube.Largest ^ base.Largest)
		mzeros.Encode(enc, v&1)
		mzeros.Encode(enc, v>>1)
	}

	SortByZ(items, func(i int) int32 {
		return current[i].X ^ baseline[i].X
	})

	fmt.Println(len(items))
	for _, i := range items {
		fmt.Println(current[i].X^baseline[i].X, current[i].Y^baseline[i].Y, current[i].Z^baseline[i].Z)
	}

	//SortBy(items, baseline.Cubes, func(i int) int32 { return })

	/*
		for i, cube := range current {
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
	*/

	enc.Close()
	return enc.Bytes()
}

func (s *State) Decode(snapshot []byte) {
	dec := arith.NewDecoder(snapshot)

	s.Current().Assign(s.Baseline())
	baseline := s.Baseline().Cubes
	current := s.Current().Cubes

	mzeros := mZeros()
	same := make([]bool, len(current))
	items := make([]int, 0, len(current))
	for i := range current {
		same[i] = mzeros.Decode(dec) == 0
		if !same[i] {
			items = append(items, i)
		}
	}

	for _, i := range items {
		cube := &current[i]
		v := mzeros.Decode(dec)
		cube.Interacting = int32(v ^ 1)
	}

	for _, i := range items {
		cube, base := &current[i], &baseline[i]
		v0, v1 := mzeros.Decode(dec), mzeros.Decode(dec)
		cube.Largest = base.Largest ^ int32(v1<<1|v0)
	}

	return

	/*
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
		}*/
}

func SameBools(a, b []bool) {
	if len(a) != len(b) {
		panic("different length")
	}

	as := BoolsStr(a)
	bs := BoolsStr(b)
	if as != bs {
		fmt.Println("---")
		fmt.Println(as)
		fmt.Println(bs)
		fmt.Println("---")
	}
}

func BoolsStr(a []bool) string {
	r := make([]byte, 0, len(a))
	for _, v := range a {
		if v {
			r = append(r, '.')
		} else {
			r = append(r, 'X')
		}
	}
	return string(r)
}
