package qpc

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/montanaflynn/stats"
)

type Count int64

func Since(start Count) time.Duration {
	return Now().Sub(start).Duration()
}

type checkpoint struct {
	name     string
	duration time.Duration
}

type History struct {
	name   string
	values []time.Duration
	start  Count
	last   time.Duration
}

func NewHistory(name string) *History { return &History{name: name} }
func (hist *History) Start()          { hist.start = Now() }

func (hist *History) Stop() {
	now := Now()
	hist.last = now.Sub(hist.start).Duration()
	hist.values = append(hist.values, hist.last)
}

func (hist *History) Last() time.Duration { return hist.last }

func (hist *History) PrintSummary() {
	nanos := []float64{}
	for _, duration := range hist.values {
		nanos = append(nanos, float64(duration))
	}

	fmt.Printf("%10s", hist.name)
	fmt.Printf(" %10s", time.Duration(stats.Min(nanos)))
	for _, p := range percentiles {
		nano := time.Duration(stats.Percentile(nanos, p))
		fmt.Printf(" %10s", nano)
	}
	fmt.Printf(" %10s", time.Duration(stats.Max(nanos)))
	fmt.Println()
}

var percentiles = []float64{10, 25, 50, 75, 90}

func FprintSummary(out io.Writer, hists ...*History) {
	fmt.Fprintf(out, "%10s", "")
	fmt.Fprintf(out, " %10s", "MIN")
	for _, p := range percentiles {
		fmt.Fprintf(out, " %9d%%", int(p))
	}
	fmt.Fprintf(out, " %10s", "MAX")
	fmt.Fprintln(out)

	for _, hist := range hists {
		nanos := []float64{}
		for _, duration := range hist.values {
			nanos = append(nanos, float64(duration))
		}

		fmt.Fprintf(out, "%10s", hist.name)
		fmt.Fprintf(out, " %10s", time.Duration(stats.Min(nanos)))
		for _, p := range percentiles {
			fmt.Fprintf(out, " %10s", time.Duration(stats.Percentile(nanos, p)))
		}
		fmt.Fprintf(out, " %10s", time.Duration(stats.Max(nanos)))
		fmt.Fprintln(out)
	}
}

func PrintSummary(hists ...*History) {
	FprintSummary(os.Stdout, hists...)
}
