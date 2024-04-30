package main

import (
	"fmt"
	"strings"
	"testing"
)

func Cut0(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

func CutByte0(s string, sep byte) (before, after string, found bool) {
	if i := strings.IndexByte(s, sep); i >= 0 {
		return s[:i], s[i+1:], true
	}
	return s, "", false
}

func Cut1(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

func CutByte1(s string, sep byte) (before, after string, found bool) {
	if i := strings.IndexByte(s, sep); i >= 0 {
		return s[:i], s[i+1:], true
	}
	return s, "", false
}

func Cut2(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

func CutByte2(s string, sep byte) (before, after string, found bool) {
	if i := strings.IndexByte(s, sep); i >= 0 {
		return s[:i], s[i+1:], true
	}
	return s, "", false
}

func Cut3(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

func CutByte3(s string, sep byte) (before, after string, found bool) {
	if i := strings.IndexByte(s, sep); i >= 0 {
		return s[:i], s[i+1:], true
	}
	return s, "", false
}

var before, after string
var found bool

func BenchmarkCut0(b *testing.B) {
	b.ReportAllocs()

	for _, skip := range [...]int{2, 4, 8, 16, 32, 64} {
		s := strings.Repeat(strings.Repeat(" ", skip)+"a"+strings.Repeat(" ", skip), 1<<16/skip)
		b.Run(fmt.Sprintf("func=Cut/skip=%d", skip), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				before, after, found = Cut0(s, "a")
			}
		})

		b.Run(fmt.Sprintf("func=CutByte/skip=%d", skip), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				before, after, found = CutByte0(s, 'a')
			}
		})
	}
}
func BenchmarkCut1(b *testing.B) {
	b.ReportAllocs()

	for _, skip := range [...]int{2, 4, 8, 16, 32, 64} {
		s := strings.Repeat(strings.Repeat(" ", skip)+"a"+strings.Repeat(" ", skip), 1<<16/skip)
		b.Run(fmt.Sprintf("func=Cut/skip=%d", skip), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				before, after, found = Cut1(s, "a")
			}
		})

		b.Run(fmt.Sprintf("func=CutByte/skip=%d", skip), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				before, after, found = CutByte1(s, 'a')
			}
		})
	}
}

func BenchmarkCut2(b *testing.B) {
	b.ReportAllocs()

	for _, skip := range [...]int{2, 4, 8, 16, 32, 64} {
		s := strings.Repeat(strings.Repeat(" ", skip)+"a"+strings.Repeat(" ", skip), 1<<16/skip)
		b.Run(fmt.Sprintf("func=Cut/skip=%d", skip), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				before, after, found = Cut2(s, "a")
			}
		})

		b.Run(fmt.Sprintf("func=CutByte/skip=%d", skip), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				before, after, found = CutByte2(s, 'a')
			}
		})
	}
}

func BenchmarkCut3(b *testing.B) {
	b.ReportAllocs()

	for _, skip := range [...]int{2, 4, 8, 16, 32, 64} {
		s := strings.Repeat(strings.Repeat(" ", skip)+"a"+strings.Repeat(" ", skip), 1<<16/skip)
		b.Run(fmt.Sprintf("func=Cut/skip=%d", skip), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				before, after, found = Cut3(s, "a")
			}
		})

		b.Run(fmt.Sprintf("func=CutByte/skip=%d", skip), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				before, after, found = CutByte3(s, 'a')
			}
		})
	}
}
