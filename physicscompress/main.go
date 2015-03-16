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
// then delta encoded values will be compressed
//   at the moment it's using flate, but even simple RLE gives good results
//
// Anyways, the percentiles for the snapshot size distribution
//
//         size      speed      improve     encode     decode
// MIN    7.240kb  434.400kbps  370.277us  110.324us   82.184us
// P05    8.728kb  523.680kbps  484.620us  131.763us  106.304us
// P10    8.784kb  527.040kbps  602.090us  136.230us  113.004us
// P25   10.560kb  633.600kbps  862.043us  145.609us  126.850us
// P50   12.184kb  731.040kbps 1006.759us  154.096us  142.036us
// P75   13.656kb  819.360kbps 1119.316us  160.796us  153.649us
// P90   15.360kb  921.600kbps 1177.828us  167.495us  166.602us
// P95   16.512kb  990.720kbps 1191.228us  171.962us  174.195us
// MAX   18.664kb 1119.840kbps 1348.897us  247.893us  239.407us
//
// TOTAL 31669.424kb

package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"syscall"
	"unsafe"

	"github.com/egonelbre/deltagolomb"

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

func Encode(order *Ordering, curr Deltas) (snapshot []byte) {
	wr := NewWriter()

	wr.Write(order.Largest, curr, getLargest)
	wr.Write(order.Interacting, curr, getInteracting)

	wr.WriteIndexed(order.ABC, curr)
	wr.WriteIndexed(order.XYZ, curr)

	wr.Close()
	return wr.Bytes()
}

// ignores reading errors
func Decode(order *Ordering, curr Deltas, snapshot []byte) {
	rd := NewReader(snapshot)

	rd.Read(order.Largest, curr, setLargest)
	rd.Read(order.Interacting, curr, setInteracting)

	rd.ReadIndexed(order.ABC, curr)
	rd.ReadIndexed(order.XYZ, curr)
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

	output, err := os.Create("snaps_data.out")
	check(err)

	defer output.Close()
	outbuf := bufio.NewWriter(output)
	defer outbuf.Flush()

	buffer := bufio.NewReader(file)

	var start, stop, freq int64
	QueryPerformanceCounter(&start)
	QueryPerformanceCounter(&stop)
	QueryPerformanceFrequency(&freq)

	sizes := make([]float64, 0)
	speeds := make([]float64, 0)

	improves := make([]float64, 0)
	encodes := make([]float64, 0)
	decodes := make([]float64, 0)

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

		// measuring of code
		//
		// order.Improve(baseline)
		// snapshot := Encode(order, current)
		// Decode(order, mirror, snapshot)
		//
		const scale = 1000000.0

		runtime.GC()
		QueryPerformanceCounter(&start)
		order.Improve(baseline)
		QueryPerformanceCounter(&stop)
		improving := (float64(stop-start) * scale) / float64(freq)
		improves = append(improves, improving)

		runtime.GC()
		QueryPerformanceCounter(&start)
		snapshot := Encode(order, current)
		QueryPerformanceCounter(&stop)
		encoding := (float64(stop-start) * scale) / float64(freq)
		encodes = append(encodes, encoding)

		runtime.GC()
		QueryPerformanceCounter(&start)
		Decode(order, mirror, snapshot)
		QueryPerformanceCounter(&stop)
		decoding := (float64(stop-start) * scale) / float64(freq)
		decodes = append(decodes, decoding)
		// ---

		outbuf.Write(snapshot)

		// statistics, validation
		fps := 60.0

		// uint64_t bps = sizeof( packet ) * 60; float kbps = bps * 8.0 / 1000.0;

		size := float64(len(snapshot)*8) / 1000.0
		sizes = append(sizes, size)

		speed := size * fps
		speeds = append(speeds, speed)

		if *verbose {
			if !current.Equals(mirror) {
				fmt.Print("! ")
			}
			fmt.Printf("%04d %8.3fkb %8.3fkbps %8.3fus %8.3fus %8.3fus\n", frame, size, speed, improving, encoding, decoding)
		} else {
			if current.Equals(mirror) {
				fmt.Print(".")
			} else {
				fmt.Print("X")
			}
		}
	}
	fmt.Println()

	fmt.Printf("    %8s   %8s     %8s   %8s   %8s  \n", "size", "speed", "improve", "encode", "decode")

	fmt.Printf("MIN %8.3fkb %8.3fkbps %8.3fus %8.3fus %8.3fus\n", stats.Min(sizes), stats.Min(speeds), stats.Min(improves), stats.Min(encodes), stats.Min(decodes))
	for _, p := range []float64{5, 10, 25, 50, 75, 90, 95} {
		fmt.Printf("P%02.f %8.3fkb %8.3fkbps %8.3fus %8.3fus %8.3fus\n", p, stats.Percentile(sizes, p), stats.Percentile(speeds, p), stats.Percentile(improves, p), stats.Percentile(encodes, p), stats.Percentile(decodes, p))
	}
	fmt.Printf("MAX %8.3fkb %8.3fkbps %8.3fus %8.3fus %8.3fus\n", stats.Max(sizes), stats.Max(speeds), stats.Max(improves), stats.Max(encodes), stats.Max(decodes))

	fmt.Println()
	total := stats.Sum(sizes)
	fmt.Printf("TOTAL %.3fkb\n", total)
	fmt.Printf("AVG %.3fkb per frame\n", total/float64(frame))
	fmt.Printf("AVG %.3fbits per cube\n", total*1000/float64(frame*N))
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

type IndexValue struct {
	Value byte
	Index int
}

func (v IndexValue) Get(deltas Deltas) int32 {
	switch v.Value {
	case 0:
		return deltas[v.Index].A
	case 1:
		return deltas[v.Index].B
	case 2:
		return deltas[v.Index].C
	case 3:
		return deltas[v.Index].X
	case 4:
		return deltas[v.Index].Y
	case 5:
		return deltas[v.Index].Z
	}
	return 0
}

func (v IndexValue) Set(deltas Deltas, x int32) {
	switch v.Value {
	case 0:
		deltas[v.Index].A = x
	case 1:
		deltas[v.Index].B = x
	case 2:
		deltas[v.Index].C = x
	case 3:
		deltas[v.Index].X = x
	case 4:
		deltas[v.Index].Y = x
	case 5:
		deltas[v.Index].Z = x
	}
}

type byValue struct {
	order  []IndexValue
	deltas Deltas
}

func (s *byValue) Len() int           { return len(s.order) }
func (s *byValue) Swap(i, j int)      { s.order[i], s.order[j] = s.order[j], s.order[i] }
func (s *byValue) Less(i, j int) bool { return s.order[i].Get(s.deltas) < s.order[j].Get(s.deltas) }

func sequence(n int) []int {
	r := make([]int, n)
	for i := range r {
		r[i] = i
	}
	return r
}

func NewOrdering(deltas Deltas) *Ordering {
	n := len(deltas)
	order := &Ordering{
		Largest:     sequence(n),
		Interacting: sequence(n),

		ABC: make([]IndexValue, n*3),
		XYZ: make([]IndexValue, n*3),
	}

	for i := 0; i < n*3; i += 1 {
		order.ABC[i].Index = i % n
		order.ABC[i].Value = byte(i / n)

		order.XYZ[i].Index = i % n
		order.XYZ[i].Value = byte(i/n + 3)
	}

	return order
}

type sorter struct {
	order  []int
	deltas Deltas
	get    Getter
}

func (s *sorter) Len() int      { return len(s.order) }
func (s *sorter) Swap(i, j int) { s.order[i], s.order[j] = s.order[j], s.order[i] }
func (s *sorter) Less(i, j int) bool {
	return s.get(&s.deltas[s.order[i]]) < s.get(&s.deltas[s.order[j]])
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

func (order *Ordering) Improve(deltas Deltas) {
	sort.Sort(&byValue{order.ABC, deltas})
	sort.Sort(&byValue{order.XYZ, deltas})
	improve(order.Largest, deltas, getLargest)
	improve(order.Interacting, deltas, getInteracting)
}

// Delta Writer/Reader
type Getter func(a *Delta) int32
type Setter func(a *Delta, v int32)

type Writer struct {
	enc *deltagolomb.ExpGolombEncoder
	buf *bytes.Buffer
}

func NewWriter() *Writer {
	var buf bytes.Buffer
	return &Writer{deltagolomb.NewExpGolombEncoder(&buf), &buf}
}

func (wr *Writer) Write(order []int, values Deltas, get Getter) {
	p := int32(0)
	for _, i := range order {
		v := get(&values[i])
		d := v - p
		wr.enc.WriteInt(int(d))
		p = v
	}
	return
}

func (wr *Writer) WriteIndexed(order []IndexValue, values Deltas) {
	p := int32(0)
	for _, i := range order {
		v := i.Get(values)
		d := v - p
		wr.enc.WriteInt(int(d))
		p = v
	}
	return
}

func (wr *Writer) Close()        { wr.enc.Close() }
func (wr *Writer) Bytes() []byte { return wr.buf.Bytes() }

type Reader struct {
	dec *deltagolomb.ExpGolombDecoder
	buf *bytes.Buffer
}

func NewReader(data []byte) *Reader {
	buf := bytes.NewBuffer(data)
	dec := deltagolomb.NewExpGolombDecoder(buf)
	return &Reader{dec, buf}
}

func (rd *Reader) Read(order []int, values Deltas, set Setter) error {
	deltas := make([]int, len(order))
	rd.dec.Read(deltas)

	p := int32(0)
	for i, idx := range order {
		d := int32(deltas[i])
		v := p + d
		set(&values[idx], v)
		p = v
	}

	return nil
}

func (rd *Reader) ReadIndexed(order []IndexValue, values Deltas) error {
	deltas := make([]int, len(order))
	rd.dec.Read(deltas)

	p := int32(0)
	for i, idx := range order {
		d := int32(deltas[i])
		v := p + d
		idx.Set(values, v)
		p = v
	}

	return nil
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

// precision timing
var (
	modkernel32                   = syscall.NewLazyDLL("kernel32.dll")
	procQueryPerformanceFrequency = modkernel32.NewProc("QueryPerformanceFrequency")
	procQueryPerformanceCounter   = modkernel32.NewProc("QueryPerformanceCounter")
)

func QueryPerformanceFrequency(frequency *int64) (err error) {
	r1, _, e1 := syscall.Syscall(procQueryPerformanceFrequency.Addr(), 1, uintptr(unsafe.Pointer(frequency)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func QueryPerformanceCounter(counter *int64) (err error) {
	r1, _, e1 := syscall.Syscall(procQueryPerformanceCounter.Addr(), 1, uintptr(unsafe.Pointer(counter)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
