package arith

const (
	PBits = 12
	MaxP  = 1 << PBits
)

type P uint32

func Prob(p float64) P { return P(float64(p) * float64(MaxP)) }

func (p P) midpoint32(lo, hi uint32) uint32 {
	return lo + uint32(uint64(hi-lo)*uint64(p)>>PBits)
}

type BinEncoder struct {
	lo, hi uint32
	data   []byte
}

func NewBinEncoder() *BinEncoder { return &BinEncoder{lo: 0, hi: ^uint32(0)} }

func (enc *BinEncoder) Encode(bit int, prob P) {
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

func (enc *BinEncoder) Close() {
	for i := 0; i < 4; i++ {
		enc.data = append(enc.data, byte(enc.lo>>24))
		enc.lo <<= 8
	}
}

func (enc *BinEncoder) Bytes() []byte { return enc.data }

type BinDecoder struct {
	lo, hi uint32
	code   uint32
	data   []byte
	read   int
}

func NewBinDecoder(data []byte) *BinDecoder {
	dec := &BinDecoder{lo: 0, hi: ^uint32(0)}
	dec.data = data

	for i := 0; i < 4; i++ {
		dec.code = dec.code<<8 | uint32(dec.data[dec.read])
		dec.read++
	}

	return dec
}

func (dec *BinDecoder) Decode(prob P) (bit int) {
	x := prob.midpoint32(dec.lo, dec.hi)

	if dec.code < x {
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
