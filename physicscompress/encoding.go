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

func (wr *Writer) WriteBools(order []int, current []bool) {
	for _, idx := range order {
		if current[idx] {
			IntWriter{wr}.WriteInt(1)
		} else {
			IntWriter{wr}.WriteInt(0)
		}
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
	pack := flate.NewReader(buf)
	return &Reader{buf, pack}
}

func (rd *Reader) ReadBools(order []int, current []bool) {
	for _, i := range order {
		v, _ := IntReader{rd}.ReadInt()
		current[i] = v == 1
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
