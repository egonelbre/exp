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

func NewWriter() *Writer {
	var buf bytes.Buffer
	pack, err := flate.NewWriter(&buf, flate.DefaultCompression)
	if err != nil {
		panic(err)
	}
	return &Writer{&buf, pack}
}

func (wr *Writer) WriteDelta(order []int, current Deltas, get Getter) {
	p := int32(0)
	for _, i := range order {
		v := get(&current[i])
		d := v - p
		IntWriter{wr}.WriteInt(int(d))
		p = v
	}
	return
}

func (wr *Writer) WriteIndexed(order []IndexValue, baseline, current Deltas) {
	for _, idx := range order {
		d := idx.Get(current) - idx.Get(baseline)
		IntWriter{wr}.WriteInt(int(d))
	}
	return
}

func (wr *Writer) Close() error  { return wr.WriteCloser.Close() }
func (wr *Writer) Bytes() []byte { return wr.buf.Bytes() }

type Reader struct {
	buf *bytes.Buffer
	io.ReadCloser
}

func NewReader(data []byte) *Reader {
	buf := bytes.NewBuffer(data)
	pack := flate.NewReader(buf)
	return &Reader{buf, pack}
}

func (rd *Reader) ReadDelta(order []int, current Deltas, set Setter) error {
	p := int32(0)
	for _, idx := range order {
		d, _ := IntReader{rd}.ReadInt()
		v := p + int32(d)
		set(&current[idx], v)
		p = v
	}
	return nil
}

func (rd *Reader) ReadIndexed(order []IndexValue, baseline, current Deltas) error {
	for _, idx := range order {
		d, _ := IntReader{rd}.ReadInt()
		idx.Set(current, int32(d)+idx.Get(baseline))
	}
	return nil
}

// IntWriter/IntReader
type IntWriter struct{ io.Writer }

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
