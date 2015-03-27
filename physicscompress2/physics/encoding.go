package physics

import (
	"fmt"

	"github.com/egonelbre/exp/bit"
	"github.com/egonelbre/exp/coder/arith"
)

func maxbits(vs ...int) (r uint) {
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
	_ = historic
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
		v := uint(1)
		if cube.Interacting {
			v = 0
		}
		mzeros.Encode(enc, v)
	}

	props := IndexProps(items, len(baseline))
	SortByZ(props, Props(baseline))
	getProp := Props(current)

	max := uint64(0)
	for _, i := range props {
		ext := uint64(bit.ZEncode(int64(getProp(i))))
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
	for _, i := range props {
		val := uint(bit.ZEncode(int64(getProp(i))))
		mvalues.Encode(enc, val)
	}

	enc.Close()
	return enc.Bytes()
}

func (s *State) Decode(snapshot []byte) {
	dec := arith.NewDecoder(snapshot)

	s.Current().Assign(s.Baseline())
	historic := s.Historic().Cubes
	_ = historic
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
		cube.Interacting = mzeros.Decode(dec) == 0
	}

	props := IndexProps(items, len(baseline))
	SortByZ(props, Props(baseline))
	setProp := SetProps(current)

	nbits := uint(0)
	for mzeros.Decode(dec) == 0 {
		nbits += 1
	}

	if nbits == 0 {
		for _, i := range props {
			setProp(i, 0)
		}
		return
	}

	mvalues := mValues(nbits)
	for _, i := range props {
		z := mvalues.Decode(dec)
		v := bit.ZDecode(uint64(z))
		setProp(i, int(v))
	}
}

func IndexProps(index []int, N int) []int {
	r := make([]int, 0, len(index)*7)
	for _, v := range index {
		r = append(r, v, v+N, v+N*2, v+N*3, v+N*4, v+N*5, v+N*6)
	}
	return r
}

func Props(base []Cube) func(i int) int {
	N := len(base)
	return func(i int) int {
		k := i % N
		switch i / N {
		case 0:
			return int(base[k].Pos.X * UnitsPerMeter)
		case 1:
			return int(base[k].Pos.Y * UnitsPerMeter)
		case 2:
			return int(base[k].Pos.Z * UnitsPerMeter)
		case 3:
			return int(base[k].Rot.X * UnitsPerQuat)
		case 4:
			return int(base[k].Rot.Y * UnitsPerQuat)
		case 5:
			return int(base[k].Rot.Z * UnitsPerQuat)
		case 6:
			return int(base[k].Rot.W * UnitsPerQuat)
		default:
			panic("invalid")
		}
	}
}

func SetProps(cur []Cube) func(i int, v int) {
	N := len(cur)
	return func(i int, v int) {
		k := i % N
		switch i / N {
		case 0:
			cur[k].Pos.X = float32(v) / UnitsPerMeter
		case 1:
			cur[k].Pos.Y = float32(v) / UnitsPerMeter
		case 2:
			cur[k].Pos.Z = float32(v) / UnitsPerMeter
		case 3:
			cur[k].Rot.X = float32(v) / UnitsPerQuat
		case 4:
			cur[k].Rot.Y = float32(v) / UnitsPerQuat
		case 5:
			cur[k].Rot.Z = float32(v) / UnitsPerQuat
		case 6:
			cur[k].Rot.W = float32(v) / UnitsPerQuat
		default:
			panic("invalid")
		}
	}
}

func DeltaProps(hist, base []Cube) func(i int) int {
	N := len(base)
	return func(i int) int {
		k := i % N
		switch i / N {
		case 0:
			return int((hist[k].Pos.X - base[k].Pos.X) * UnitsPerMeter)
		case 1:
			return int((hist[k].Pos.Y - base[k].Pos.Y) * UnitsPerMeter)
		case 2:
			return int((hist[k].Pos.Z - base[k].Pos.Z) * UnitsPerMeter)
		case 3:
			return int((hist[k].Rot.X - base[k].Rot.X) * UnitsPerQuat)
		case 4:
			return int((hist[k].Rot.Y - base[k].Rot.Y) * UnitsPerQuat)
		case 5:
			return int((hist[k].Rot.Z - base[k].Rot.Z) * UnitsPerQuat)
		case 6:
			return int((hist[k].Rot.W - base[k].Rot.W) * UnitsPerQuat)
		default:
			panic("invalid")
		}
	}
}

func SetDeltaProps(base, cur []Cube) func(i int, v int) {
	N := len(base)
	return func(i int, v int) {
		k := i % N
		switch i / N {
		case 0:
			cur[k].Pos.X = float32(v)/UnitsPerMeter + base[k].Pos.X
		case 1:
			cur[k].Pos.Y = float32(v)/UnitsPerMeter + base[k].Pos.Y
		case 2:
			cur[k].Pos.Z = float32(v)/UnitsPerMeter + base[k].Pos.Z
		case 3:
			cur[k].Rot.X = float32(v)/UnitsPerQuat + base[k].Rot.X
		case 4:
			cur[k].Rot.Y = float32(v)/UnitsPerQuat + base[k].Rot.Y
		case 5:
			cur[k].Rot.Z = float32(v)/UnitsPerQuat + base[k].Rot.Z
		case 6:
			cur[k].Rot.W = float32(v)/UnitsPerQuat + base[k].Rot.W
		default:
			panic("invalid")
		}
	}
}

func ExtraProps(hist, base, cur []Cube) func(i int) int {
	N := len(base)
	return func(i int) int {
		k := i % N
		switch i / N {
		case 0:
			return int((cur[k].Pos.X - 2*base[k].Pos.X + hist[k].Pos.X) * UnitsPerMeter)
		case 1:
			return int((cur[k].Pos.Y - 2*base[k].Pos.Y + hist[k].Pos.Y) * UnitsPerMeter)
		case 2:
			return int((cur[k].Pos.Z - 2*base[k].Pos.Z + hist[k].Pos.Z) * UnitsPerMeter)
		case 3:
			return int((cur[k].Rot.X - 2*base[k].Rot.X + hist[k].Rot.X) * UnitsPerQuat)
		case 4:
			return int((cur[k].Rot.Y - 2*base[k].Rot.Y + hist[k].Rot.Y) * UnitsPerQuat)
		case 5:
			return int((cur[k].Rot.Z - 2*base[k].Rot.Z + hist[k].Rot.Z) * UnitsPerQuat)
		case 6:
			return int((cur[k].Rot.W - 2*base[k].Rot.W + hist[k].Rot.W) * UnitsPerQuat)
		default:
			panic("invalid")
		}
	}
}

func SetExtraProps(hist, base, cur []Cube) func(i int, v int) {
	N := len(base)
	return func(i int, v int) {
		k := i % N
		switch i / N {
		case 0:
			cur[k].Pos.X = float32(v)/UnitsPerMeter + 2*base[k].Pos.X - hist[k].Pos.X
		case 1:
			cur[k].Pos.Y = float32(v)/UnitsPerMeter + 2*base[k].Pos.Y - hist[k].Pos.Y
		case 2:
			cur[k].Pos.Z = float32(v)/UnitsPerMeter + 2*base[k].Pos.Z - hist[k].Pos.Z
		case 3:
			cur[k].Rot.X = float32(v)/UnitsPerQuat + 2*base[k].Rot.X - hist[k].Rot.X
		case 4:
			cur[k].Rot.Y = float32(v)/UnitsPerQuat + 2*base[k].Rot.Y - hist[k].Rot.Y
		case 5:
			cur[k].Rot.Z = float32(v)/UnitsPerQuat + 2*base[k].Rot.Z - hist[k].Rot.Z
		case 6:
			cur[k].Rot.W = float32(v)/UnitsPerQuat + 2*base[k].Rot.W - hist[k].Rot.W
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
