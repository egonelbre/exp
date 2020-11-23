package mispredict

import (
	"math"
	"math/rand"
	"sort"
	"testing"
	"unsafe"
)

type Shape interface {
	Area() float32
}

type Circle struct{ Radius float32 }
type Square struct{ Side float32 }

func (s Circle) Area() float32 {
	return math.Pi * s.Radius * s.Radius
}
func (s Square) Area() float32 {
	return s.Side * s.Side
}

func TotalArea(shapes []Shape) (total float32) {
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// random shapes
var shapes = func() []Shape {
	shapes := make([]Shape, 1e4)
	for i := range shapes {
		if rand.Intn(2) == 0 {
			shapes[i] = Circle{rand.Float32()}
		} else {
			shapes[i] = Square{rand.Float32()}
		}
	}
	return shapes
}()

func BenchmarkRandomShapes(b *testing.B) {
	total := float32(0)
	for k := 0; k < b.N; k++ {
		total += TotalArea(shapes)
	}
}

var sortedShapes = func() []Shape {
	sorted := make([]Shape, len(shapes))
	copy(sorted, shapes)

	type iface struct {
		itab uintptr
		data unsafe.Pointer
	}
	sort.Slice(sorted, func(i, k int) bool {
		// don't do this in production
		a := (*iface)(unsafe.Pointer(&sorted[i])).itab
		b := (*iface)(unsafe.Pointer(&sorted[k])).itab
		return a < b
	})
	return sorted
}()

func BenchmarkSortedShapes(b *testing.B) {
	total := float32(0)
	for k := 0; k < b.N; k++ {
		total += TotalArea(sortedShapes)
	}
}
