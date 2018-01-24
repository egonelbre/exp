package main

import (
	"fmt"
	"os"
	"text/template"
)

const Header = `package goqueuestest
// DO NOT EDIT, GENERATED CODE
import "testing"

`

const testTemplate = `
// {{.Name}}

func Test{{.Name}}(t *testing.T) {
	{{.Type}}Test(t, {{.Ctor}})
	queuePTest(t, {{.Ctor}})
	queueP2Test(t, {{.Ctor}})
}
`
const benchTemplate = "func Benchmark{{.Name}}_{{.Bench}}x{{.Routines}}(b *testing.B) { runParallelBench(b, {{.Routines}}, {{.Ctor}}, bench{{.Bench}})}\n"

type queue struct {
	Name, Ctor, Type string
}

type bench struct {
	queue
	Bench string
	Routines int
}

var queues = []queue{
	{"CFifo", "NewChanFifo(benchIterations)", "fifo"},
	{"ZLifo", "NewZLifo()", "lifo"},
	{"ZFifo", "NewZFifo()", "fifo"},
	// {"ZcFifo", "NewZFifoFreechan()", "fifo"},
	// {"ZrFifo", "NewZFifoFreering()", "fifo"},
	{"ScLifo", "NewScLifo()", "lifo"},
	{"ScFifo", "NewScFifo()", "fifo"},
	{"SmLifo", "NewSmLifo()", "lifo"},
	{"SmFifo", "NewSmFifo()", "fifo"},
	{"RcLifo", "NewRcLifo()", "lifo"},
	{"RcFifo", "NewRcFifo()", "fifo"},
	{"RmLifo", "NewRmLifo()", "lifo"},
	{"RmFifo", "NewRmFifo()", "fifo"},
	//{"PcLifo", "NewPcLifo()", "lifo"},
	//{"PcFifo", "NewPcFifo()", "fifo"},
	//{"PmLifo", "NewPmLifo()", "lifo"},
	//{"PmFifo", "NewPmFifo()", "fifo"},
	//{"LcLifo", "NewListCLifo()", "lifo"},
	//{"LcFifo", "NewListCFifo()", "fifo"},
	//{"LmLifo", "NewListLifo()", "lifo"},
	//{"LmFifo", "NewListFifo()", "fifo"},
}

var routines = []int { 1, 3, 5, 10, 50, 100 }
var benchs = []string { "ANRN", "A1R1", "A1R2", "A2R1" }

func main() {
	tTest := template.Must(template.New("testset").Parse(testTemplate))
	tBench := template.Must(template.New("benchmark").Parse(benchTemplate))

	fmt.Printf(Header)
	for _, q := range queues {
		tTest.Execute(os.Stdout, q)
		for _, benchName := range benchs {
			for _, routines := range routines {
				var b bench
				b.Name = q.Name
				b.Ctor = q.Ctor
				b.Type = q.Type
				b.Bench = benchName
				b.Routines = routines
				tBench.Execute(os.Stdout, b)
			}
		}
	}
}
