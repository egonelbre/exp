package ber

import (
	"bytes"
	"io"
)

type forkableWriter struct {
	*bytes.Buffer
	pre, post *forkableWriter
}

func newForkableWriter() *forkableWriter {
	return &forkableWriter{new(bytes.Buffer), nil, nil}
}
func (f *forkableWriter) fork() (pre, post *forkableWriter) {
	if f.pre != nil || f.post != nil {
		panic("have already forked")
	}
	f.pre = newForkableWriter()
	f.post = newForkableWriter()
	return f.pre, f.post
}
func (f *forkableWriter) Len() (l int) {
	l += f.Buffer.Len()
	if f.pre != nil {
		l += f.pre.Len()
	}
	if f.post != nil {
		l += f.post.Len()
	}
	return
}

func (f *forkableWriter) writeTo(out io.Writer) (n int, err error) {
	n, err = out.Write(f.Bytes())
	if err != nil {
		return
	}

	var nn int

	if f.pre != nil {
		nn, err = f.pre.writeTo(out)
		n += nn
		if err != nil {
			return
		}
	}

	if f.post != nil {
		nn, err = f.post.writeTo(out)
		n += nn
	}
	return
}

func marshalHeader(out *forkableWriter, h header) (err error) {
	b := uint8(h.class) << 6
	if h.constructed {
		b |= 0x20
	}
	if h.tag >= 31 {
		b |= 0x1f
		err = out.WriteByte(b)
		if err != nil {
			return
		}
		err = marshalBase128Int(out, int64(h.tag))
		if err != nil {
			return
		}
	} else {
		b |= uint8(h.tag)
		err = out.WriteByte(b)
		if err != nil {
			return
		}
	}
	if h.length >= 128 {
		l := lengthLength(h.length)
		err = out.WriteByte(0x80 | byte(l))
		if err != nil {
			return
		}
		err = marshalLength(out, h.length)
		if err != nil {
			return
		}
	} else {
		err = out.WriteByte(byte(h.length))
		if err != nil {
			return
		}
	}
	return nil
}

func marshalBase128Int(out *forkableWriter, n int64) (err error) {
	if n == 0 {
		err = out.WriteByte(0)
		return
	}

	l := 0
	for i := n; i > 0; i >>= 7 {
		l++
	}

	for i := l - 1; i >= 0; i-- {
		o := byte(n >> uint(i*7))
		o &= 0x7f
		if i != 0 {
			o |= 0x80
		}
		err = out.WriteByte(o)
		if err != nil {
			return
		}
	}

	return nil
}

func marshalLength(out *forkableWriter, i int) (err error) {
	n := lengthLength(i)

	for ; n > 0; n-- {
		err = out.WriteByte(byte(i >> uint((n-1)*8)))
		if err != nil {
			return
		}
	}

	return nil
}

func lengthLength(i int) (numBytes int) {
	numBytes = 1
	for i > 255 {
		numBytes++
		i >>= 8
	}
	return
}
