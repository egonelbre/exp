package sort_test

import (
	"sort"
	"testing"

	"github.com/shawnsmithdev/zermelo/zuint64"

	"github.com/egonelbre/exp/sorts/dpsortint"
	"github.com/egonelbre/exp/sorts/qpsortint"
	"github.com/egonelbre/exp/sorts/rdxsort"
	"github.com/egonelbre/exp/sorts/stdsortint"
	"github.com/zeebo/pcg"
)

func bench(b *testing.B, size int, algo func([]int), name string) {
	b.StopTimer()

	data := make([]int, size)

	for i := 0; i < b.N; i++ {
		for n := size - 3; n <= size+3; n++ {
			rng := pcg.New(uint64(n))
			for i := 0; i < len(data); i++ {
				data[i] = int(rng.Uint64())
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

func benchRadix(b *testing.B, size int, algo func(src, dst []uint64), name string) {
	b.StopTimer()

	data := make([]uint64, size)
	buf := make([]uint64, len(data))

	for i := 0; i < b.N; i++ {
		for n := size - 3; n <= size+3; n++ {
			rng := pcg.New(uint64(n))
			for i := 0; i < len(data); i++ {
				data[i] = rng.Uint64()
			}

			b.StartTimer()
			algo(data, buf)
			b.StopTimer()

			if !sort.SliceIsSorted(data, func(i, k int) bool { return data[i] < data[k] }) {
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

func BenchmarkRdxSort1e4(b *testing.B)        { benchRadix(b, 1e4, rdxsort.Uint64, "Rdx") }
func BenchmarkRdxSort1e6(b *testing.B)        { benchRadix(b, 1e6, rdxsort.Uint64, "Rdx") }
func BenchmarkZermeloRdxSort1e4(b *testing.B) { benchRadix(b, 1e4, zuint64.SortBYOB, "Zrdx") }
func BenchmarkZermeloRdxSort1e6(b *testing.B) { benchRadix(b, 1e6, zuint64.SortBYOB, "Zrdx") }
