package expgolomb

import "github.com/egonelbre/exp/bit"

func WriteInt(w *bit.Writer, v int) error {
	if v == 0 {
		return w.WriteBit(1)
	}

	sign := 0
	if v < 0 {
		sign = 1
		v = -v
	}

	bits := bit.ScanRight(uint64(v)) + 1

	if err := w.WriteBits(0, bits); err != nil {
		return err
	}
	if err := w.WriteBitsReverse(uint64(v), bits); err != nil {
		return err
	}

	if err := w.WriteBit(sign); err != nil {
		return err
	}

	return nil
}

func ReadInt(r *bit.Reader) (int, error) {
	bits := uint64(0)

	zero, err := r.ReadBit()
	for err == nil && zero == 0 {
		bits += 1
		zero, err = r.ReadBit()
	}

	if bits == 0 || err != nil {
		return 0, err
	}

	x, err := r.ReadBits(bits - 1)
	if err != nil {
		return 0, err
	}
	x = x<<1 | 1

	x = bit.Reverse(x, bits)

	sign, err := r.ReadBit()
	if err != nil {
		return 0, err
	}

	v := int(x)
	if sign == 1 {
		v = -v
	}
	return v, nil
}

func WriteUint(w *bit.Writer, v uint) error {
	if v == 0 {
		return w.WriteBit(1)
	}

	bits := bit.ScanRight(uint64(v)) + 1

	if err := w.WriteBits(0, bits); err != nil {
		return err
	}
	if err := w.WriteBitsReverse(uint64(v), bits); err != nil {
		return err
	}

	return nil
}

func ReadUint(r *bit.Reader) (uint, error) {
	bits := uint64(0)

	zero, err := r.ReadBit()
	for err == nil && zero == 0 {
		bits += 1
		zero, err = r.ReadBit()
	}

	if bits == 0 || err != nil {
		return 0, err
	}

	x, err := r.ReadBits(bits - 1)
	if err != nil {
		return 0, err
	}
	x = x<<1 | 1

	x = bit.Reverse(x, bits)

	return uint(x), nil
}
