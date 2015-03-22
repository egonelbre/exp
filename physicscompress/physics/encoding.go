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
	return &arith.Shift2{P0: 0x34a, I0: 0x1, P1: 0xd0, I1: 0x5}
}

func mValues(nbits uint) arith.Model {
	return arith.NewTree(nbits, func() arith.Model {
		return &arith.Shift2{P0: 0x438, I0: 0x3, P1: 0x61a, I1: 0x1}
	})
}

var totalbits = 0

func (s *State) Encode() []byte {
	enc := arith.NewEncoder()

	historic := s.Historic().Cubes
	baseline := s.Baseline().Cubes
	current := s.Current().Cubes

	mzeros := mZeros()
	items := []int{}
	for i := range current {
		cube, base := &current[i], &baseline[i]
		if *cube == *base {
			mzeros.Encode(enc, 0)
		} else {
			mzeros.Encode(enc, 1)
			items = append(items, i)
		}
	}

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

	items6 := Index6(items, len(baseline))
	SortByZ(items6, Delta6(historic, baseline))
	ext6 := Extra6(historic, baseline, current)

	max := uint64(0)
	for _, i := range items6 {
		ext := uint64(bit.ZEncode(int64(ext6(i))))
		if max < ext {
			max = ext
		}
	}

	if max == 0 {
		mzeros.Encode(enc, 1)
		enc.Close()
		return enc.Bytes()
	}

	nbits := bit.ScanRight(max) + 1
	for i := 0; i < int(nbits); i += 1 {
		mzeros.Encode(enc, 0)
	}
	mzeros.Encode(enc, 1)

	mvalues := mValues(nbits)
	for _, i := range items6 {
		val := uint(bit.ZEncode(int64(ext6(i))))
		mvalues.Encode(enc, val)
	}

	enc.Close()
	return enc.Bytes()
}

func (s *State) Decode(snapshot []byte) {
	dec := arith.NewDecoder(snapshot)

	s.Current().Assign(s.Baseline())
	historic := s.Historic().Cubes
	baseline := s.Baseline().Cubes
	current := s.Current().Cubes

	mzeros := mZeros()

	items := []int{}
	for i := range current {
		if mzeros.Decode(dec) == 1 {
			items = append(items, i)
		}
	}

	for _, i := range items {
		cube := &current[i]
		cube.Interacting = int32(mzeros.Decode(dec) ^ 1)
	}

	for _, i := range items {
		cube, base := &current[i], &baseline[i]
		v0, v1 := mzeros.Decode(dec), mzeros.Decode(dec)
		cube.Largest = base.Largest ^ int32(v1<<1|v0)
	}

	items6 := Index6(items, len(baseline))
	SortByZ(items6, Delta6(historic, baseline))
	set6 := SetExtra6(historic, baseline, current)

	nbits := uint(0)
	for mzeros.Decode(dec) == 0 {
		nbits += 1
	}

	if nbits == 0 {
		for _, i := range items6 {
			set6(i, 0)
		}
		return
	}

	mvalues := mValues(nbits)
	for _, i := range items6 {
		z := mvalues.Decode(dec)
		v := bit.ZDecode(uint64(z))
		set6(i, int32(v))
	}
}

func Index6(index []int, N int) []int {
	r := make([]int, 0, len(index)*6)
	for _, v := range index {
		r = append(r, v, v+N, v+N*2, v+N*3, v+N*4, v+N*5)
	}
	return r
}

func Value6(base []Cube) func(i int) int32 {
	N := len(base)
	return func(i int) int32 {
		k := i % N
		switch i / N {
		case 0:
			return base[k].A
		case 1:
			return base[k].B
		case 2:
			return base[k].C
		case 3:
			return base[k].X
		case 4:
			return base[k].Y
		case 5:
			return base[k].Z
		default:
			panic("invalid")
		}
	}
}

func Delta6(hist, base []Cube) func(i int) int32 {
	N := len(base)
	return func(i int) int32 {
		k := i % N
		switch i / N {
		case 0:
			return hist[k].B - base[k].A
		case 1:
			return hist[k].B - base[k].B
		case 2:
			return hist[k].C - base[k].C
		case 3:
			return hist[k].X - base[k].X
		case 4:
			return hist[k].Y - base[k].Y
		case 5:
			return hist[k].Z - base[k].Z
		default:
			panic("invalid")
		}
	}
}

func Extra6(hist, base, cur []Cube) func(i int) int32 {
	N := len(base)
	return func(i int) int32 {
		k := i % N
		switch i / N {
		case 0:
			return cur[k].A - base[k].A + (hist[k].B - base[k].A)
		case 1:
			return cur[k].B - base[k].B + (hist[k].B - base[k].B)
		case 2:
			return cur[k].C - base[k].C + (hist[k].C - base[k].C)
		case 3:
			return cur[k].X - base[k].X + (hist[k].X - base[k].X)
		case 4:
			return cur[k].Y - base[k].Y + (hist[k].Y - base[k].Y)
		case 5:
			return cur[k].Z - base[k].Z + (hist[k].Z - base[k].Z)
		default:
			panic("invalid")
		}
	}
}

func SetExtra6(hist, base, cur []Cube) func(i int, v int32) {
	N := len(base)
	return func(i int, v int32) {
		k := i % N
		switch i / N {
		case 0:
			cur[k].A = v + base[k].A - (hist[k].B - base[k].A)
		case 1:
			cur[k].B = v + base[k].B - (hist[k].B - base[k].B)
		case 2:
			cur[k].C = v + base[k].C - (hist[k].C - base[k].C)
		case 3:
			cur[k].X = v + base[k].X - (hist[k].X - base[k].X)
		case 4:
			cur[k].Y = v + base[k].Y - (hist[k].Y - base[k].Y)
		case 5:
			cur[k].Z = v + base[k].Z - (hist[k].Z - base[k].Z)
		default:
			panic("invalid")
		}
	}
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
