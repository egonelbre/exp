package syncpool

import (
	"sync"
	"testing"
)

const numElems = 10000

func BenchmarkPoolPointer(b *testing.B) {
	pool := sync.Pool{
		New: func() any {
			s := make([]byte, 0, numElems)
			return &s
		},
	}

	s := pool.Get().(*[]byte)
	if len(*s) != 0 {
		b.Error("returned slice doesn't have the expected number of elements")
	}
	s2 := *s
	s2 = s2[:0]
	pool.Put(&s2)

	for b.Loop() {
		s := pool.Get().(*[]byte)
		s2 := *s
		s2 = append(s2, 10)
		if len(s2) != 1 {
			b.Error("returned slice doesn't have the expected number of elements")
		}
		s2 = s2[:0]
		*s = s2
		pool.Put(s)
	}
}

func BenchmarkPoolValue(b *testing.B) {
	pool := sync.Pool{
		New: func() any {
			return make([]byte, 0, numElems)
		},
	}

	s := pool.Get().([]byte)
	if len(s) != 0 {
		b.Error("returned slice doesn't have the expected number of elements")
	}
	s = s[:0]
	pool.Put(s)

	for b.Loop() {
		s := pool.Get().([]byte)
		s = append(s, 10)
		if len(s) != 1 {
			b.Error("returned slice doesn't have the expected number of elements")
		}
		s = s[:0]
		pool.Put(s)
	}
}

func BenchmarkPoolWrapper(b *testing.B) {
	pool := sync.Pool{
		New: func() any {
			s := make([]byte, 0, numElems)
			return &s
		},
	}

	get := func() []byte {
		s := pool.Get().(*[]byte)
		return *s
	}

	put := func(s []byte) {
		pool.Put(&s)
	}

	s := get()
	if len(s) != 0 {
		b.Error("returned slice doesn't have the expected number of elements")
	}
	s = s[:0]
	put(s)

	for b.Loop() {
		s := get()
		s = append(s, 10)
		if len(s) != 1 {
			b.Error("returned slice doesn't have the expected number of elements")
		}
		s = s[:0]
		put(s)
	}
}

func BenchmarkPoolWrapperPointers(b *testing.B) {
	pool := sync.Pool{
		New: func() any {
			s := make([]byte, 0, numElems)
			return &s
		},
	}

	appendByte := func(buffer *[]byte, val byte) *[]byte {
		b := *buffer
		b = append(b, val)
		return &b
	}

	reset := func(buffer *[]byte) *[]byte {
		b := *buffer
		b = b[:0]
		return &b
	}

	s := pool.Get().(*[]byte)
	if len(*s) != 0 {
		b.Error("returned slice doesn't have the expected number of elements")
	}
	s = reset(s)
	pool.Put(s)

	for b.Loop() {
		s := pool.Get().(*[]byte)
		s = appendByte(s, 10)
		if len(*s) != 1 {
			b.Error("returned slice doesn't have the expected number of elements")
		}
		s = reset(s)
		pool.Put(s)
	}
}
