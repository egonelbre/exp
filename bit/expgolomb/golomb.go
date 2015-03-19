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

	nbits := bit.ScanRight(uint64(v)) + 1

	w.WriteBits(0, nbits)
	rbits := bit.Reverse(uint64(v), nbits)
	w.WriteBits(rbits, nbits)
	w.WriteBit(sign)
}

func ReadInt(r *bit.Reader) int {
	nbits := uint(0)

	zero, err := r.ReadBit(), r.Error()
	for err == nil && zero == 0 {
		nbits += 1
		zero, err = r.ReadBit(), r.Error()
	}

	if nbits == 0 || err != nil {
		return 0
	}

	x := r.ReadBits(nbits - 1)
	x = x<<1 | 1
	x = bit.Reverse(x, nbits)

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

	nbits := bit.ScanRight(uint64(v)) + 1

	w.WriteBits(0, nbits)
	rbits := bit.Reverse(uint64(v), nbits)
	w.WriteBits(rbits, nbits)
}

func ReadUint(r *bit.Reader) uint {
	nbits := uint(0)
	zero, err := r.ReadBit(), r.Error()
	for err == nil && zero == 0 {
		nbits += 1
		zero, err = r.ReadBit(), r.Error()
	}

	if nbits == 0 || err != nil {
		return 0
	}

	x := r.ReadBits(nbits - 1)
	x = x<<1 | 1
	x = bit.Reverse(x, nbits)
	return uint(x)
}
