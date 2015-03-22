// Package bit implements bit level encoding wrappers for io.Reader / io.Writer
// All operations are Little Endian where applicable
//
// Note that Write/Read operations do not return error for convenience,
// you must check bit.Reader.Error() or bit.Writer.Error() manually
//
package bit

import "io"

type Writer struct {
	w     io.Writer
	bits  uint64
	nbits uint
	err   error
}

func NewWriter(w io.Writer) *Writer { return &Writer{w, 0, 0, nil} }

// flush writes as much full bytes as possible to the underlying writer
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

// flushAll writes all the remaining half bytes
func (w *Writer) flushAll() {
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

func (w *Writer) Error() error { return w.err }

// Align aligns the writer to the next byte
func (w *Writer) Align() { w.flushAll() }

// WriteBits writes width lowest bits to the underlying writer
func (w *Writer) WriteBits(x uint64, width uint) {
	if width > 32 {
		w.WriteBits(uint64(uint32(x)), 32)
		x >>= 32
		width -= 32
	}

	x &= 1<<width - 1
	w.bits |= x << w.nbits
	w.nbits += width
	if w.nbits > 16 {
		w.flush()
	}
}

// WriteBit writes the lowest bit in x to the underlying writer
func (w *Writer) WriteBit(x int) {
	w.WriteBits(uint64(x&1), 1)
}

// WriteBool writes a bool the underlying writer depending on x
func (w *Writer) WriteBool(x bool) {
	if x {
		w.WriteBits(1, 1)
	}
	w.WriteBits(0, 1)
}

// WriteByte writes a byte to the underlying writer
func (w *Writer) WriteByte(v byte) {
	w.WriteBits(uint64(v), 8)
}

func (w *Writer) Close() error {
	w.Align()
	return w.err
}

type Reader struct {
	r     io.Reader
	bits  uint64
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
	r.bits = uint64(temp[0])
}

func (r *Reader) Error() error { return r.err }

// Align aligns the reader to the next byte so that the next ReadBits will start
// reading a new byte from the underlying reader
func (r *Reader) Align() {
	r.nbits = 8
}

// ReadBits reads width bits from the underlying reader
func (r *Reader) ReadBits(width uint) uint64 {
	if r.err != nil {
		return 0
	}

	left := 8 - int(r.nbits)
	if left > int(width) {
		mask := uint64((1 << width) - 1)
		x := r.bits >> r.nbits
		r.nbits += width
		return x & mask
	}

	n := 8 - r.nbits
	x := r.bits >> r.nbits
	for int(width)-int(n) > 0 {
		r.read()
		r.nbits -= 8
		if r.err != nil {
			return 0
		}
		x |= r.bits << n
		n += 8
	}
	r.nbits += width
	mask := uint64(1<<width - 1)
	return x & mask
}

// ReadBit reads a single bit from the underlying reader
func (r *Reader) ReadBit() int { return int(r.ReadBits(1)) }

// ReadBool reads a single bool from the underlying reader
func (r *Reader) ReadBool() bool { return r.ReadBits(1) == 1 }

// ReadByte reads a single bit from the underlying reader
func (r *Reader) ReadByte() byte { return byte(r.ReadBits(8)) }
