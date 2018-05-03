package combiner

import "testing"

func Test(t *testing.T)        { All.TestDefault(t) }
func TestLatency(t *testing.T) { All.LatencyDefault(t) }
func Benchmark(b *testing.B)   { All.BenchmarkDefault(b) }
