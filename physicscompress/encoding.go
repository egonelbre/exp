package main

import (
	"bytes"
	"compress/flate"
	"encoding/binary"
	"io"
)

type Getter func(a *Delta) int32
type Setter func(a *Delta, v int32)

type Writer struct {
	buf *bytes.Buffer
	io.WriteCloser
}

type noclose struct{ *bytes.Buffer }

func (nc noclose) Close() error { return nil }

func NewWriter() *Writer {
	var buf bytes.Buffer
	if dontpack {
		return &Writer{&buf, noclose{&buf}}
	}
	pack, err := flate.NewWriter(&buf, flate.DefaultCompression)
	if err != nil {
		panic(err)
	}

	return &Writer{&buf, pack}
}

func (wr *Writer) WriteBools(current []bool) {
	wr.WriteBytes(PackBools(current))
}

func (wr *Writer) WriteBytes(data []byte) {
	x := IntWriter{wr}
	for _, v := range data {
		x.WriteByte(v)
	}
}

func (wr *Writer) WriteDelta(nochange []bool, order []int, current Deltas, get Getter) {
	p := int32(0)
	for _, idx := range order {
		if nochange[idx] {
			continue
		}
		v := get(&current[idx])
		d := v - p
		IntWriter{wr}.WriteInt(int(d))
		p = v
	}
}

func (wr *Writer) WriteIndexed(nochange []bool, order []IndexValue, baseline, current Deltas) {
	for _, idx := range order {
		if nochange[idx.Index] {
			continue
		}
		d := idx.Get(current) - idx.Get(baseline)
		IntWriter{wr}.WriteInt(int(d))
	}
}

func (wr *Writer) Close() error  { return wr.WriteCloser.Close() }
func (wr *Writer) Bytes() []byte { return wr.buf.Bytes() }

type Reader struct {
	buf *bytes.Buffer
	io.ReadCloser
}

func NewReader(data []byte) *Reader {
	buf := bytes.NewBuffer(data)
	if dontpack {
		return &Reader{buf, noclose{buf}}
	}
	pack := flate.NewReader(buf)
	return &Reader{buf, pack}
}

func (rd *Reader) ReadBools(current []bool) {
	packed := make([]byte, (len(current)+7)/8)
	rd.ReadBytes(packed)
	UnpackBools(packed, current)
}

func (rd *Reader) ReadBytes(current []byte) {
	for i, _ := range current {
		v, _ := IntReader{rd}.ReadByte()
		current[i] = v
	}
}

func (rd *Reader) ReadDelta(nochange []bool, order []int, current Deltas, set Setter) {
	p := int32(0)
	for _, idx := range order {
		if nochange[idx] {
			continue
		}
		d, _ := IntReader{rd}.ReadInt()
		v := p + int32(d)
		set(&current[idx], v)
		p = v
	}
}

func (rd *Reader) ReadIndexed(nochange []bool, order []IndexValue, baseline, current Deltas) {
	for _, idx := range order {
		if nochange[idx.Index] {
			continue
		}
		d, _ := IntReader{rd}.ReadInt()
		idx.Set(current, int32(d)+idx.Get(baseline))
	}
}

// IntWriter/IntReader
type IntWriter struct{ io.Writer }

func (w IntWriter) WriteByte(v byte) error {
	var x [1]byte
	x[0] = v
	_, err := w.Write(x[:])
	return err
}

func (w IntWriter) WriteInt(x int) (int, error) {
	var buf [binary.MaxVarintLen64]byte
	n := binary.PutVarint(buf[:], int64(x))
	return w.Write(buf[:n])
}

type IntReader struct{ io.Reader }

func (r IntReader) ReadByte() (byte, error) {
	var x [1]byte
	_, err := r.Read(x[:])
	return x[0], err
}

func (r IntReader) ReadInt() (int, error) {
	x, err := binary.ReadVarint(r)
	return int(x), err
}

func PackBools(vs []bool) []byte {
	r := make([]byte, (len(vs)+7)/8)
	for i, v := range vs {
		if !v {
			r[i/8] |= byte(1 << uint(i%8))
		}
	}
	return r
}

func UnpackBools(packed []byte, into []bool) {
	for i := range into {
		into[i] = (packed[i/8] >> uint(i%8) & 1) == 0
	}
}

func Reorder(nochange []bool, order []int, current Deltas, get Getter) []int32 {
	r := make([]int32, 0, len(order))
	for _, i := range order {
		if !nochange[i] {
			r = append(r, get(&current[i]))
		}
	}
	return r
}

func PackBit1(vs []int32) []byte {
	r := make([]byte, (len(vs)+7)/8)
	for i, v := range vs {
		if v == 1 {
			r[i/8] |= byte(1 << uint(i%8))
		}
	}
	return r
}

func PackBit2(vs []int32) []byte {
	r := make([]byte, (len(vs)+3)/4)
	for i, v := range vs {
		r[i/4] |= byte((v & 3) << uint(i%4))
	}
	return r
}
