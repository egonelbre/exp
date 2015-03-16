package main

import (
	"bytes"
	"compress/flate"
	"io"

	"github.com/egonelbre/deltagolomb"
)

// Delta Writer/Reader
type Getter func(a *Delta) int32
type Setter func(a *Delta, v int32)

type Writer struct {
	enc *deltagolomb.ExpGolombEncoder
	com *flate.Writer
	buf *bytes.Buffer
}

func NewWriter() *Writer {
	var buf bytes.Buffer
	com, err := flate.NewWriter(&buf, flate.DefaultCompression)
	if err != nil {
		panic(err)
	}
	enc := deltagolomb.NewExpGolombEncoder(com)
	return &Writer{enc, com, &buf}
}

func (wr *Writer) Write(order []int, current Deltas, get Getter) {
	p := int32(0)
	for _, i := range order {
		v := get(&current[i])
		d := v - p
		wr.enc.WriteInt(int(d))
		p = v
	}
	return
}

func (wr *Writer) WriteIndexed(order []IndexValue, baseline, current Deltas) {
	for _, idx := range order {
		v := idx.Get(current) - idx.Get(baseline)
		wr.enc.WriteInt(int(v))
	}
	return
}

func (wr *Writer) Close() {
	wr.enc.Close()
	wr.com.Close()
}
func (wr *Writer) Bytes() []byte { return wr.buf.Bytes() }

type Reader struct {
	dec *deltagolomb.ExpGolombDecoder
	com io.ReadCloser
	buf *bytes.Buffer
}

func NewReader(data []byte) *Reader {
	buf := bytes.NewBuffer(data)
	com := flate.NewReader(buf)
	dec := deltagolomb.NewExpGolombDecoder(com)
	return &Reader{dec, com, buf}
}

func (rd *Reader) Read(order []int, current Deltas, set Setter) error {
	deltas := make([]int, len(order))
	rd.dec.Read(deltas)

	p := int32(0)
	for i, idx := range order {
		d := int32(deltas[i])
		v := p + d
		set(&current[idx], v)
		p = v
	}

	return nil
}

func (rd *Reader) ReadIndexed(order []IndexValue, baseline, current Deltas) error {
	deltas := make([]int, len(order))
	rd.dec.Read(deltas)

	for i, idx := range order {
		d := int32(deltas[i])
		idx.Set(current, d+idx.Get(baseline))
	}

	return nil
}
