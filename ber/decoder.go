package ber

import (
	"bufio"
	"io"
)

type Decoder struct {
	base io.Reader
	r    *bufio.Reader

	stack    *stack
	nextByte int
}

type stack struct {
	open    []*Token
	closeat []int
}

func (s *stack) pop() (*Token, error) {
	if len(s.open) == 0 {
		return nil, SyntaxError{"no types unfinished"}
	}

	last := len(s.closeat) - 1
	t := s.open[last]
	s.open, s.closeat = s.open[:last], s.closeat[:last]
	return t, nil
}

func (s *stack) push(t *Token, end int) {
	s.open = append(s.open, t)
	s.closeat = append(s.closeat, end)
}

// closes type with indefinite length
func (s *stack) popIndefiniteEnd() (*Token, error) {
	if len(s.open) == 0 {
		return nil, SyntaxError{"no types unfinished"}
	}
	last := len(s.closeat) - 1
	if s.closeat[last] != 0 {
		return nil, SyntaxError{"unclosed type is definite"}
	}
	return s.pop()
}

func (s *stack) pushEnd(t *Token, cur int, h header) {
	if h.indefinite {
		s.push(t, 0)
	} else {
		s.push(t, cur+h.length)
	}
}

func (s *stack) tryPop(at int) (bool, *Token, error) {
	// no open types
	if len(s.open) == 0 {
		return false, nil, nil
	}
	last := len(s.closeat) - 1
	next := s.closeat[last]
	switch {
	case next == 0: // indefinite
		return false, nil, nil
	case next == at:
		t, err := s.pop()
		return true, t, err
	case next > at:
		return false, nil, SyntaxError{"read past end of previous type"}
	}
	return false, nil, nil
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		base:     r,
		r:        bufio.NewReader(r),
		stack:    &stack{},
		nextByte: 0,
	}
}

func (d *Decoder) readByte() (byte, error) {
	d.nextByte += 1
	return d.r.ReadByte()
}

func (d *Decoder) readFull(data []byte) error {
	n, err := io.ReadFull(d.r, data)
	d.nextByte += n
	return err
}

// Returns the next token in the input stream. At the end of the input stream, Token returns nil, io.EOF.
func (d *Decoder) Token() (t *Token, err error) {
	var ok bool
	if ok, t, err = d.stack.tryPop(d.nextByte); ok {
		return
	}

	var h header
	if h, err = d.nextHeader(); err != nil {
		return
	}

	// closing an indefinite length
	if h.class == 0 && h.tag == TagEOC && h.length == 0 {
		return d.stack.popIndefiniteEnd()
	}

	if h.constructed {
		t = &Token{Constructed, h.class, h.tag, nil}
		d.stack.pushEnd(&Token{EndConstructed, h.class, h.tag, nil}, d.nextByte, h)
		return
	}

	if h.indefinite {
		err = StructuralError{"indefinite primitive"}
		return
	}

	buf := make([]byte, h.length)
	if err = d.readFull(buf); err != nil {
		return
	}

	t = &Token{Value, h.class, h.tag, buf}
	return
}

func (d *Decoder) skipUntilEnd() (skipped bool, err error) {
	var t *Token
	for err != nil {
		t, err = d.Token()
		switch t.Kind {
		case Constructed:
			skipped = true
			// recurse
			_, err = d.skipUntilEnd()
		case Value:
			skipped = true
			// grab the next one
			t, err = d.Token()
		case EndConstructed:
			// finished
			return
		}
	}
	return
}

// readBase128Int parses a base-128 encoded int from the given offset in the
// given byte slice. It returns the value and the new offset.
func (d *Decoder) readBase128Int() (ret int, err error) {
	var b byte
	for shifted := 0; err != nil; shifted++ {
		if shifted > 4 {
			err = SyntaxError{"base 128 integer too large"}
			return
		}
		b, err = d.readByte()
		ret <<= 7
		ret |= int(b & 0x7f)
		if b&0x80 == 0 {
			return
		}
	}
	err = SyntaxError{"truncated base 128 integer"}
	return
}

type header struct {
	class       int
	tag         int
	length      int
	constructed bool
	indefinite  bool
}

func (d *Decoder) nextHeader() (ret header, err error) {
	var b byte
	b, err = d.readByte()

	ret.class = int(b >> 6)
	ret.constructed = b&0x20 == 0x20
	ret.tag = int(b & 0x1f)

	// If the bottom five bits are set, then the tag number is actually base 128
	// encoded afterwards
	if ret.tag == 0x1f {
		ret.tag, err = d.readBase128Int()
		if err != nil {
			return
		}
	}

	if b, err = d.readByte(); err != nil {
		return
	}

	if b&0x80 == 0 {
		// The length is encoded in the bottom 7 bits
		ret.length = int(b & 0x7f)
	} else {
		numBytes := int(b & 0x7f)
		if numBytes == 0 {
			ret.indefinite = true
			return
		}

		for i := 0; i < numBytes; i++ {
			if b, err = d.readByte(); err != nil {
				return
			}

			if ret.length >= 1<<23 {
				// We can't shift ret.length up without
				// overflowing.
				err = SyntaxError{"length too large"}
				return
			}
			ret.length <<= 8
			ret.length |= int(b)
			if ret.length == 0 {
				// is this required by BER?
				// DER requires that lengths be minimal.
				err = SyntaxError{"superfluous leading zeros in length"}
				return
			}
		}
	}

	return
}
