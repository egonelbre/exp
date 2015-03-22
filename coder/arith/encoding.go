// This package implements simple arithmetic coders
// based on code by https://github.com/rygorous/mini_arith
package arith

type Encoder struct {
	lo, hi uint32
	data   []byte
}

func NewEncoder() *Encoder { return &Encoder{lo: 0, hi: ^uint32(0)} }

func (enc *Encoder) Encode(bit uint, prob P) {
	x := prob.midpoint32(enc.lo, enc.hi)

	if bit == 1 {
		enc.hi = x
	} else {
		enc.lo = x + 1
	}

	for enc.lo^enc.hi < 1<<24 {
		enc.data = append(enc.data, byte(enc.lo>>24))
		enc.lo <<= 8
		enc.hi = enc.hi<<8 | 0xFF
	}
}

func (enc *Encoder) Close() {
	for i := 0; i < 4; i++ {
		enc.data = append(enc.data, byte(enc.lo>>24))
		enc.lo <<= 8
	}
}

func (enc *Encoder) Bytes() []byte { return enc.data }

type Decoder struct {
	lo, hi uint32
	code   uint32
	data   []byte
	read   int
}

func NewDecoder(data []byte) *Decoder {
	dec := &Decoder{lo: 0, hi: ^uint32(0)}
	dec.data = data

	for i := 0; i < 4; i++ {
		dec.code = dec.code<<8 | uint32(dec.data[dec.read])
		dec.read++
	}

	return dec
}

func (dec *Decoder) Decode(prob P) (bit uint) {
	x := prob.midpoint32(dec.lo, dec.hi)

	if dec.code <= x {
		dec.hi = x
		bit = 1
	} else {
		dec.lo = x + 1
		bit = 0
	}

	for dec.lo^dec.hi < 1<<24 {
		dec.code = dec.code<<8 | uint32(dec.data[dec.read])
		dec.read++
		dec.lo <<= 8
		dec.hi = dec.hi<<8 | 0xff
	}

	return bit
}
