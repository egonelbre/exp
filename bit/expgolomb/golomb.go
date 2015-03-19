package expgolomb

import "github.com/egonelbre/exp/bit"

func WriteInt(w *bit.Writer, v int) {
	if v == 0 {
		w.WriteBit(1)
		return
	}

	sign := 0
	if v < 0 {
		sign = 1
		v = -v
	}

	bits := bit.ScanRight(uint64(v)) + 1

	w.WriteBits(0, bits)
	w.WriteBitsReverse(uint64(v), bits)
	w.WriteBit(sign)
}

func ReadInt(r *bit.Reader) int {
	bits := uint(0)

	zero, err := r.ReadBit(), r.Error()
	for err == nil && zero == 0 {
		bits += 1
		zero, err = r.ReadBit(), r.Error()
	}

	if bits == 0 || err != nil {
		return 0
	}

	x := r.ReadBits(bits - 1)
	x = x<<1 | 1
	x = bit.Reverse(x, bits)

	sign := r.ReadBit()
	v := int(x)
	if sign == 1 {
		v = -v
	}

	return v
}

func WriteUint(w *bit.Writer, v uint) {
	if v == 0 {
		w.WriteBit(1)
		return
	}

	bits := bit.ScanRight(uint64(v)) + 1

	w.WriteBits(0, bits)
	w.WriteBitsReverse(uint64(v), bits)
}

func ReadUint(r *bit.Reader) uint {
	bits := uint(0)
	zero, err := r.ReadBit(), r.Error()
	for err == nil && zero == 0 {
		bits += 1
		zero, err = r.ReadBit(), r.Error()
	}

	if bits == 0 || err != nil {
		return 0
	}

	x := r.ReadBits(bits - 1)
	x = x<<1 | 1
	x = bit.Reverse(x, bits)
	return uint(x)
}
