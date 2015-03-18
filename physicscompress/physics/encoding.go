package physics

import (
	"bytes"
	"compress/flate"
	"io"

	"github.com/egonelbre/exp/bit"
)

type Writer struct {
	*bit.Writer
	pack   io.WriteCloser
	buffer *bytes.Buffer
}

func NewWriter() *Writer {
	var buf bytes.Buffer
	pack, _ := flate.NewWriter(&buf, flate.DefaultCompression)
	return &Writer{
		bit.NewWriter(pack),
		pack,
		&buf,
	}
}

func (w *Writer) Close() {
	w.Writer.Close()
	w.pack.Close()
}

type Reader struct {
	*bit.Reader
	buffer *bytes.Buffer
}

func NewReader(data []byte) *Reader {
	buf := bytes.NewBuffer(data)
	pack := flate.NewReader(buf)
	return &Reader{
		bit.NewReader(pack),
		buf,
	}
}

func (r *Reader) Close() {

}

func (w *Writer) Prepare(baseline, current *Frame) {}

func (r *Reader) Prepare(baseline, current *Frame) { current.Assign(baseline) }

func (w *Writer) Bools(same []bool) {
	for _, v := range same {
		w.WriteBool(!v)
	}
}

func (r *Reader) Bools(same []bool) {
	for i := range same {
		v, _ := r.ReadBool()
		same[i] = !v
	}
}

func (w *Writer) ValueBits(same []bool, order []int, vals Values, bits uint64) {
	for _, i := range order {
		if same[vals.Index(i)] {
			continue
		}
		v := vals.Get(i)
		w.WriteBits(uint64(v), bits)
	}
}

func (r *Reader) ValueBits(same []bool, order []int, vals Values, bits uint64) {
	for _, i := range order {
		if same[vals.Index(i)] {
			continue
		}
		v, _ := r.ReadBits(bits)
		vals.Set(i, int32(v))
	}
}

func (w *Writer) Values(same []bool, order []int, vals Values) {
	max := uint64(0)
	for _, i := range order {
		if same[vals.Index(i)] {
			continue
		}

		x := bit.AbsEncode(int64(vals.Get(i)))
		if x > max {
			max = x
		}
	}
	if max == 0 {
		w.WriteBits(0, 6)
		return
	}

	bits := bit.ScanRight(max) + 1
	w.WriteBits(bits, 6)

	for _, i := range order {
		if same[vals.Index(i)] {
			continue
		}
		v := vals.Get(i)
		uv := bit.AbsEncode(int64(v))
		w.WriteBits(uv, bits)
	}
}

func (r *Reader) Values(same []bool, order []int, vals Values) {
	bits, _ := r.ReadBits(6)
	if bits == 0 {
		return
	}
	for _, i := range order {
		if same[vals.Index(i)] {
			continue
		}
		uv, _ := r.ReadBits(bits)
		v := bit.AbsDecode(uv)
		vals.Set(i, int32(v))
	}
}

func (w *Writer) Bytes() []byte {
	return w.buffer.Bytes()
}
