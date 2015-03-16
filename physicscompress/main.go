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
// Anyways, the all the interesting stats
//
// MIN     11.520 kbps
// P05     18.240 kbps
// P10    122.880 kbps
// P25    250.080 kbps
// P50    367.200 kbps
// P75    471.840 kbps
// P90    564.000 kbps
// P95    607.680 kbps
// MAX    705.120 kbps
//
// TOTAL   16804.168 kb
//   AVG       5.923 kb per frame
//   AVG       6.574 bits per cube
//
// TIMING:
//                   MIN        10%        25%        50%        75%        90%        MAX
//    improve   58.065µs   75.037µs   80.844µs    86.65µs    92.01µs   95.584µs   208.14µs
//     encode   340.35µs  516.332µs  580.204µs  627.549µs  670.428µs  709.287µs 1.327903ms
//     decode  118.363µs  167.942µs  187.148µs  205.461µs  242.086µs  259.059µs  428.788µs
//   pimprove  535.092µs  664.175µs  748.592µs  835.243µs  920.108µs 1.001399ms  1.57133ms
//
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

func ReadDelta(r io.Reader, current Deltas) error {
	for i := range current {
		if err := current[i].ReadFrom(r); err != nil {
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

	var history [7]Deltas
	for i := range history {
		history[i] = make(Deltas, N)
	}

	frame := 0
	for i := 0; i < 6; i += 1 {
		ReadDelta(buffer, history[frame])
		frame += 1
	}

	mirror := make(Deltas, N)

	for {
		baseline := history[(frame-6+len(history))%len(history)]
		current := history[frame%len(history)]

		err = ReadDelta(buffer, current)
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
