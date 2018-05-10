package combiner

import (
	"flag"
	"testing"
)

func Test(t *testing.T) { All.TestDefault(t) }

var latency = flag.Bool("latency", false, "run latency measurements")

func TestLatency(t *testing.T) {
	if !*latency {
		t.Skip("latency measurements disabled")
	}
	All.LatencyDefault(t)
}

func Benchmark(b *testing.B) { All.BenchmarkDefault(b) }
