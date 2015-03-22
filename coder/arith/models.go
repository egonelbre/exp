package arith

import "github.com/egonelbre/exp/bit"

type Model interface {
	NBits() uint
	Encode(enc *Encoder, value uint)
	Decode(dec *Decoder) (value uint)
}

type Shift struct {
	P P
	I byte
}

func (m *Shift) NBits() uint { return 1 }

func (m *Shift) adapt(bit uint) {
	switch bit {
	case 1:
		m.P += (MaxP - m.P) >> m.I
	case 0:
		m.P -= m.P >> m.I
	}
}

func (m *Shift) Encode(enc *Encoder, bit uint) {
	enc.Encode(bit, m.P)
	m.adapt(bit)
}

func (m *Shift) Decode(dec *Decoder) (bit uint) {
	bit = dec.Decode(m.P)
	m.adapt(bit)
	return bit
}

type Shift2 struct {
	P0 P
	I0 byte

	P1 P
	I1 byte
}

func (m *Shift2) NBits() uint { return 1 }

func (m *Shift2) adapt(bit uint) {
	switch bit {
	case 1:
		m.P0 += (MaxP/2 - m.P0) >> m.I0
		m.P1 += (MaxP/2 - m.P1) >> m.I1
	case 0:
		m.P0 -= m.P0 >> m.I0
		m.P1 -= m.P1 >> m.I1
	}
}

func (m *Shift2) Encode(enc *Encoder, bit uint) {
	enc.Encode(bit, m.P0+m.P1)
	m.adapt(bit)
}

func (m *Shift2) Decode(dec *Decoder) (bit uint) {
	bit = dec.Decode(m.P0 + m.P1)
	m.adapt(bit)
	return bit
}

type Tree []Model

func (tree Tree) NBits() uint { return bit.ScanRight(uint64(tree.syms())) }

func (tree Tree) syms() uint { return uint(len(tree) + 1) }
func (tree Tree) msb() uint  { return tree.syms() / 2 }

func NewTree(nbits uint, model func() Model) Tree {
	syms := 1 << nbits
	tree := make(Tree, syms-1)
	for i := range tree {
		tree[i] = model()
	}
	return tree
}

func NewEmptyTree(nbits uint) Tree {
	return make(Tree, 1<<nbits-1)
}

func (tree Tree) Encode(enc *Encoder, value uint) {
	if value > tree.syms() {
		panic("")
	}

	syms, msb := tree.syms(), tree.msb()
	ctx := uint(1)
	for ctx < syms {
		bit := uint(0)
		if value&msb != 0 {
			bit = 1
		}

		value += value
		tree[ctx-1].Encode(enc, bit)
		ctx += ctx + bit
	}
}

func (tree Tree) Decode(dec *Decoder) (value uint) {
	ctx := uint(1)
	syms := tree.syms()
	for ctx < syms {
		ctx += ctx + tree[ctx-1].Decode(dec)
	}
	return ctx - syms
}
