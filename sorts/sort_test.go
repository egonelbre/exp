package sort_test

import (
	"testing"

	"github.com/egonelbre/exp/sorts/dpsortint"
	"github.com/egonelbre/exp/sorts/qpsortint"
	"github.com/egonelbre/exp/sorts/stdsortint"
)

func bench(b *testing.B, size int, algo func([]int), name string) {
	b.StopTimer()

	data := make([]int, size)

	x := ^uint32(0)
	for i := 0; i < b.N; i++ {
		for n := size - 3; n <= size+3; n++ {
			for i := 0; i < len(data); i++ {
				x += x
				x ^= 1
				if int32(x) < 0 {
					x ^= 0x88888eef
				}
				data[i] = int(x % uint32(n/5))
			}

			b.StartTimer()
			algo(data)
			b.StopTimer()

			if !stdsortint.IsSorted(data) {
				b.Errorf("%s did not sort %d ints", name, n)
			}
		}
	}
}

func BenchmarkStdSort1e2(b *testing.B) { bench(b, 1e2, stdsortint.Sort, "Std") }
func BenchmarkStdSort1e4(b *testing.B) { bench(b, 1e4, stdsortint.Sort, "Std") }
func BenchmarkStdSort1e6(b *testing.B) { bench(b, 1e6, stdsortint.Sort, "Std") }

func BenchmarkQpSort1e2(b *testing.B) { bench(b, 1e2, qpsortint.Sort, "Qp") }
func BenchmarkQpSort1e4(b *testing.B) { bench(b, 1e4, qpsortint.Sort, "Qp") }
func BenchmarkQpSort1e6(b *testing.B) { bench(b, 1e6, qpsortint.Sort, "Qp") }

func BenchmarkQpSortBasic1e2(b *testing.B) { bench(b, 1e2, qpsortint.SortBasic, "Qp") }
func BenchmarkQpSortBasic1e4(b *testing.B) { bench(b, 1e4, qpsortint.SortBasic, "Qp") }
func BenchmarkQpSortBasic1e6(b *testing.B) { bench(b, 1e6, qpsortint.SortBasic, "Qp") }

func BenchmarkDpSort1e2(b *testing.B) { bench(b, 1e2, dpsortint.Sort, "Dp") }
func BenchmarkDpSort1e4(b *testing.B) { bench(b, 1e4, dpsortint.Sort, "Dp") }
func BenchmarkDpSort1e6(b *testing.B) { bench(b, 1e6, dpsortint.Sort, "Dp") }
