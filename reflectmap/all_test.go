package reflectmap

import (
	"testing"
)

func TestReflect(t *testing.T) {
	m := NewMap(0, 0)
	m.Add(56, 12)
}
func TestUnsafe(t *testing.T) {
	m := NewUnsafeMap(0, 0)
	m.Add(56, 12)
}

func BenchmarkReflect(b *testing.B) {
	m := NewMap(0, 0)
	for i := 0; i < b.N; i++ {
		m.Add(i, i)
	}
}

func BenchmarkUnsafe(b *testing.B) {
	m := NewUnsafeMap(0, 0)
	for i := 0; i < b.N; i++ {
		m.Add(i, i)
	}
}
func BenchmarkSpecial(b *testing.B) {
	m := NewSpecialMap()
	for i := 0; i < b.N; i++ {
		m.Add(i, i)
	}
}
