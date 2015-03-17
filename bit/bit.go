package bit

import "io"

type Writer struct {
	w     io.Writer
	bits  uint64
	nbits uint
	err   error
}

func NewWriter(w io.Writer) *Writer { return &Writer{w, 0, 0, nil} }

// flush writes as much as it can to the underlying writer
func (w *Writer) flush() {
	if w.err != nil {
		w.nbits = 0
		return
	}

	var buf [16]byte
	n := 0
	for w.nbits > 8 {
		buf[n] = byte(w.bits)
		w.bits >>= 8
		w.nbits -= 8
		n++
	}

	_, w.err = w.w.Write(buf[0:n])
}

// flushpartial flushes all the remaining half bytes
func (w *Writer) flushpartial() {
	w.flush()
	if w.err != nil {
		w.nbits = 0
		return
	}

	if w.nbits > 0 {
		_, w.err = w.w.Write([]byte{byte(w.bits)})
		w.bits = 0
		w.nbits = 0
	}
}

// Align aligns the writer to the next byte
func (w *Writer) Align() error {
	w.flushpartial()
	return w.err
}

// WriteBits writes width lowest bits to the underlying writer
func (w *Writer) WriteBits(x, width uint) error {
	w.bits |= uint64(x) << w.nbits
	w.nbits += width
	if w.nbits > 16 {
		w.flush()
	}
	return w.err
}

// WriteBitsReverse writes width lowest bits in reverse order to the underlying writer
func (w *Writer) WriteBitsReverse(x, width uint) error {
	return w.WriteBits(ReverseBits(x, width), width)
}

// WriteBit writes the lowest bit in x to the underlying writer
func (w *Writer) WriteBit(x int) error {
	return w.WriteBits(uint(x&1), 1)
}

// WriteBool writes the lowest bit in x to the underlying writer
func (w *Writer) WriteBool(x bool) error {
	if x {
		return w.WriteBits(1, 1)
	}
	return w.WriteBits(0, 1)
}

// WriteByte writes a byte to the underlying writer
func (w *Writer) WriteByte(v byte) error {
	return w.WriteBits(uint(v), 8)
}

func (w *Writer) Close() error { return w.Align() }

type Reader struct {
	r     io.Reader
	bits  uint
	nbits uint
	err   error
}

func NewReader(r io.Reader) *Reader {
	return &Reader{r, 0, 8, nil}
}

// read reads a single byte from the underlying reader
func (r *Reader) read() {
	if r.err != nil {
		r.nbits = 8
		return
	}

	var temp [1]byte
	_, r.err = r.r.Read(temp[:])
	r.bits = uint(temp[0])
}

// Align aligns the reader to the next byte so that the next ReadBits will start
// reading a new byte from the underlying reader
func (r *Reader) Align() {
	r.nbits = 8
}

// ReadBits reads width bits from the underlying reader
func (r *Reader) ReadBits(width uint) (uint, error) {
	if r.err != nil {
		return 0, r.err
	}

	left := 8 - int(r.nbits)
	if left > int(width) {
		mask := uint((1 << width) - 1)
		x := r.bits >> r.nbits
		r.nbits += width
		return x & mask, nil
	}

	n := 8 - r.nbits
	x := r.bits >> r.nbits
	for int(width)-int(n) > 0 {
		r.read()
		r.nbits -= 8
		if r.err != nil {
			return 0, r.err
		}
		x |= r.bits << n
		n += 8
	}
	r.nbits += width
	mask := uint(1<<width - 1)
	return x & mask, nil
}

// ReadBitsReverse reads width bits from the underlying reader in reverse order
func (r *Reader) ReadBitsReverse(width uint) (uint, error) {
	rx, err := r.ReadBits(width)
	return ReverseBits(rx, width), err
}

// ReadBit reads a single bit from the underlying reader
func (r *Reader) ReadBit() (int, error) {
	x, err := r.ReadBits(1)
	return int(x), err
}

// ReadBool reads a single bit from the underlying reader
func (r *Reader) ReadBool() (bool, error) {
	x, err := r.ReadBits(1)
	return x == 1, err
}

// ReadByte reads a single bit from the underlying reader
func (r *Reader) ReadByte() (byte, error) {
	x, err := r.ReadBits(8)
	return byte(x), err
}
