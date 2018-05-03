package combiner

import (
	"strconv"
	"testing"
)

type Combiner interface {
	Do(op Argument)
}

type Batcher interface {
	Start()
	Include(op Argument)
	Finish()
}

type Argument interface{}

// Other possibilities
// 1. Include(v interface)
// 2. type Op func()
// 2. type Op interface { Execute() }
// 3. specialized

var (
	TestProcs  = []int{1, 2, 3, 4, 8, 16, 32, 64, 128, 256}
	TestBounds = []int{1, 2, 3, 4, 5, 6, 7, 8}

	BenchProcs  = []int{1, 2, 4, 8, 16, 32, 128, 512}
	BenchBounds = []int{4, 8, 16, 32, 64}
)

type Desc struct {
	Name    string
	Bounded bool
	Create  func(exe Batcher, bound int) Combiner
}

type Descs []Desc

type TestCase func(t *testing.T, procs int, create func(Batcher) Combiner)
type BenchCase func(b *testing.B, procs int, create func(Batcher) Combiner)

func (descs Descs) TestDefault(t *testing.T) {
	t.Helper()
	descs.Test(t, RunTests)
}
func (descs Descs) LatencyDefault(t *testing.T) {
	t.Helper()
	descs.Latency(t, RunLatency)
}
func (descs Descs) BenchmarkDefault(b *testing.B) {
	b.Helper()
	descs.Benchmark(b, RunBenchmarks)
}

func (descs Descs) Test(t *testing.T, test TestCase) {
	t.Helper()
	for _, desc := range descs {
		t.Run(desc.Name, func(t *testing.T) {
			t.Helper()
			bounds := TestBounds
			if !desc.Bounded {
				bounds = []int{0}
			}
			for _, procs := range TestProcs {
				for _, bound := range bounds {
					name := "p" + strconv.Itoa(procs) + "b" + strconv.Itoa(bound)
					t.Run(name, func(t *testing.T) {
						t.Helper()
						test(t, procs, func(exe Batcher) Combiner {
							return desc.Create(exe, bound)
						})
					})
				}
			}
		})
	}
}

func (descs Descs) Latency(t *testing.T, test TestCase) {
	t.Helper()
	for _, desc := range descs {
		t.Run(desc.Name, func(t *testing.T) {
			t.Helper()
			bounds := BenchBounds
			if !desc.Bounded {
				bounds = []int{0}
			}
			for _, procs := range BenchProcs {
				for _, bound := range bounds {
					name := "p" + strconv.Itoa(procs) + "b" + strconv.Itoa(bound)
					t.Run(name, func(t *testing.T) {
						t.Helper()
						test(t, procs, func(exe Batcher) Combiner {
							return desc.Create(exe, bound)
						})
					})
				}
			}
		})
	}
}

func (descs Descs) Benchmark(b *testing.B, bench BenchCase) {
	b.Helper()
	for _, desc := range descs {
		b.Run(desc.Name, func(b *testing.B) {
			b.Helper()
			bounds := BenchBounds
			if !desc.Bounded {
				bounds = []int{0}
			}
			for _, procs := range BenchProcs {
				for _, bound := range bounds {
					name := "p" + strconv.Itoa(procs) + "b" + strconv.Itoa(bound)
					b.Run(name, func(b *testing.B) {
						b.Helper()
						bench(b, procs, func(exe Batcher) Combiner {
							return desc.Create(exe, bound)
						})
					})
				}
			}
		})
	}
}
