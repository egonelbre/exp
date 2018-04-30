package queue

import (
	"strconv"
	"testing"

	"github.com/loov/queue/testsuite"
)

func Test(t *testing.T) {
	t.Helper()
	for _, desc := range Descs {
		batchSizes := testsuite.BatchSizes
		if !desc.HasBatchSize() {
			batchSizes = []int{0}
		}

		testSizes := testsuite.TestSizes
		if desc.IsUnbouned() {
			testSizes = []int{0}
		}

		t.Run(desc.Name, func(t *testing.T) {
			t.Helper()
			for _, batchSize := range batchSizes {
				for _, size := range testSizes {
					name := "b" + strconv.Itoa(batchSize) + "s" + strconv.Itoa(size)
					t.Run(name, func(t *testing.T) {
						t.Helper()
						testsuite.Tests(t, func() testsuite.Queue {
							return desc.Create(batchSize, size)
						})
					})
				}
			}
		})
	}
}

func Benchmark(b *testing.B) {
	b.Helper()
	for _, desc := range Descs {
		batchSizes := testsuite.BatchSizes
		if !desc.HasBatchSize() {
			batchSizes = []int{0}
		}

		benchSizes := testsuite.BenchSizes
		if desc.IsUnbouned() {
			benchSizes = []int{0}
		}

		b.Run(desc.Name, func(b *testing.B) {
			b.Helper()
			for _, batchSize := range batchSizes {
				for _, size := range benchSizes {
					name := "b" + strconv.Itoa(batchSize) + "s" + strconv.Itoa(size)
					b.Run(name, func(b *testing.B) {
						b.Helper()
						testsuite.Benchmarks(b, func() testsuite.Queue {
							return desc.Create(batchSize, size)
						})
					})
				}
			}
		})
	}
}
