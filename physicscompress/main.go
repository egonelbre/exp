// Response to Data Compression Challenge
// http://gafferongames.com/2015/03/14/the-networked-physics-data-compression-challenge/
//
// The main insight that this compression uses is
// given some previous state P and current state S
// the ordering by some attribute is very similar
//   order(P, attr) ~= order(S, attr)
//
// so we use the previous state to determine the order
// and delta encode the current attributes in that order
//
// Note that the ordering must give the same result on both computers
// but, it doesn't have to guarantee perfect ordering (e.g. see improveApprox)
//
// Anyways, the percentiles for the snapshot size distribution
//
// MIN      9.600 kbps
// P05     15.840 kbps
// P10     37.920 kbps
// P25    136.800 kbps
// P50    206.880 kbps
// P75    278.400 kbps
// P90    338.400 kbps
// P95    395.520 kbps
// MAX    492.960 kbps
//
// TOTAL    8899.984 kb
//   AVG       3.432 kb per frame
//   AVG       3.809 bits per cube
//
// TIMING:
//                   MIN        10%        25%        50%        75%        90%        MAX
//    improve   58.511µs   72.358µs   78.164µs   82.184µs   84.864µs    87.99µs  195.188µs
//     encode  348.837µs  453.354µs  522.585µs  564.124µs  595.837µs  621.296µs 1.102343ms
//     decode  116.576µs  146.056µs  158.562µs  170.622µs  180.895µs  192.954µs  379.209µs
//   pimprove  548.491µs  620.403µs  665.961µs  713.753µs  762.439µs  806.658µs 1.119316ms

package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"

	"github.com/egonelbre/exp/qpc"

	"github.com/montanaflynn/stats"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Delta struct {
	Largest     int32
	A, B, C     int32
	X, Y, Z     int32
	Interacting int32
}

type Deltas []Delta

func Encode(order *Ordering, baseline, current Deltas) (snapshot []byte) {
	wr := NewWriter()

	wr.Write(order.Largest, current, getLargest)
	wr.Write(order.Interacting, current, getInteracting)

	wr.WriteIndexed(order.ABC, baseline, current)
	wr.WriteIndexed(order.XYZ, baseline, current)

	wr.Close()
	return wr.Bytes()
}

// ignores reading errors
func Decode(order *Ordering, baseline, current Deltas, snapshot []byte) {
	rd := NewReader(snapshot)

	rd.Read(order.Largest, current, setLargest)
	rd.Read(order.Interacting, current, setInteracting)

	rd.ReadIndexed(order.ABC, baseline, current)
	rd.ReadIndexed(order.XYZ, baseline, current)
}

func ReadDeltas(r io.Reader, prev, curr Deltas) error {
	for i := range prev {
		if err := curr[i].ReadFrom(r); err != nil {
			return err
		}
		if err := prev[i].ReadFrom(r); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	verbose := flag.Bool("v", false, "verbose output")
	flag.Parse()

	file, err := os.Open("delta_data.bin")
	check(err)
	defer file.Close()

	buffer := bufio.NewReader(file)

	total := 0.0
	speeds := make([]float64, 0)

	improve := qpc.NewHistory("improve")
	encode := qpc.NewHistory("encode")
	decode := qpc.NewHistory("decode")
	pimprove := qpc.NewHistory("pimprove")

	const N = 901
	baseline := make(Deltas, N)
	order := NewOrdering(baseline)

	current := make(Deltas, N)
	mirror := make(Deltas, N)

	frame := 0
	for {
		err = ReadDeltas(buffer, baseline, current)
		if err == io.EOF {
			break
		}
		check(err)
		frame += 1

		runtime.GC()

		//===
		improve.Start()
		order.Improve(baseline)
		improve.Stop()

		encode.Start()
		snapshot := Encode(order, baseline, current)
		encode.Stop()

		decode.Start()
		Decode(order, baseline, mirror, snapshot)
		decode.Stop()

		pimprove.Start()
		order.PostImprove(baseline, mirror)
		pimprove.Stop()
		//===

		size := float64(len(snapshot)*8) / 1000.0
		total += size

		speed := size * 60.0
		speeds = append(speeds, speed)

		if *verbose {
			if !current.Equals(mirror) {
				fmt.Print("! ")
			}
			fmt.Printf("%04d %8.3fkbps %10s %10s %10s %10s\n", frame, speed, improve.Last(), encode.Last(), decode.Last(), pimprove.Last())
		} else {
			if current.Equals(mirror) {
				fmt.Print(".")
			} else {
				fmt.Print("X")
			}
		}
	}
	fmt.Println()

	fmt.Printf("MIN %10.3f kbps\n", stats.Min(speeds))
	for _, p := range []float64{5, 10, 25, 50, 75, 90, 95} {
		fmt.Printf("P%02.f %10.3f kbps\n", p, stats.Percentile(speeds, p))
	}
	fmt.Printf("MAX %10.3f kbps\n", stats.Max(speeds))

	fmt.Println()

	fmt.Printf("TOTAL  %10.3f kb\n", total)
	fmt.Printf("  AVG  %10.3f kb per frame\n", total/float64(frame))
	fmt.Printf("  AVG  %10.3f bits per cube\n", total*1000/float64(frame*N))

	fmt.Println()
	fmt.Println("TIMING:")
	qpc.PrintSummary(improve, encode, decode, pimprove)
}

// delta utilities
func (d *Delta) ReadFrom(r io.Reader) error {
	return read(r, &d.Largest, &d.A, &d.B, &d.C, &d.X, &d.Y, &d.Z, &d.Interacting)
}

func (d *Delta) WriteTo(w io.Writer) error {
	return write(w, &d.Largest, &d.A, &d.B, &d.C, &d.X, &d.Y, &d.Z, &d.Interacting)
}

// for validating content
func (a Deltas) Equals(b Deltas) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ax := range a {
		if ax != b[i] {
			return false
		}
	}
	return true
}

// getters / setters
func getLargest(a *Delta) int32     { return a.Largest }
func getInteracting(a *Delta) int32 { return a.Interacting }

func setLargest(a *Delta, v int32)     { a.Largest = v }
func setInteracting(a *Delta, v int32) { a.Interacting = v }

// Order management

type Ordering struct {
	Largest     []int
	ABC         []IndexValue
	XYZ         []IndexValue
	Interacting []int
}

var improve = improveSort

func improveSort(order []int, deltas Deltas, get Getter) {
	sort.Sort(&sorter{order, deltas, get})
}

func improveApprox(order []int, deltas Deltas, get Getter) {
	const max_steps = 10
	for k := 0; k < max_steps; k += 1 {
		prev := int32(-(1 << 31))
		swapped := false
		for i, idx := range order {
			curr := get(&deltas[idx])
			if prev > curr {
				order[i-1], order[i] = order[i], order[i-1]
				swapped = true
			} else {
				prev = curr
			}
		}
		if !swapped {
			return
		}
	}
}

func (order *Ordering) Improve(baseline Deltas) {
	improve(order.Largest, baseline, getLargest)
	improve(order.Interacting, baseline, getInteracting)
}

func (order *Ordering) PostImprove(baseline, current Deltas) {
	sort.Sort(&byDelta{order.ABC, baseline, current})
	sort.Sort(&byDelta{order.XYZ, baseline, current})
}

// reading input/output
func read(r io.Reader, vs ...interface{}) error {
	for _, v := range vs {
		err := binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func write(w io.Writer, vs ...interface{}) error {
	for _, v := range vs {
		err := binary.Write(w, binary.LittleEndian, v)
		if err != nil {
			return err
		}
	}
	return nil
}
