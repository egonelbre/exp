// Response to Data Compression Challenge
// http://gafferongames.com/2015/03/14/the-networked-physics-data-compression-challenge/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/egonelbre/exp/physicscompress/physics"
	"github.com/egonelbre/exp/qpc"

	"github.com/montanaflynn/stats"
)

var DebugSnap = flag.Bool("snap", true, "take debug frame snapshots")

func DebugSnapshot(frame int, snapshot []byte) {
	if *DebugSnap && (frame == 1000 || frame == 1200 || frame == 1201 || frame == 1770) {
		size := float64(len(snapshot)*8) / 1000.0
		speed := size * 60
		fmt.Printf("\n<%d,%.3fkb,%.3fkbps>\n", frame, size, speed)
		ioutil.WriteFile(fmt.Sprintf("snapshot_%d.bin", frame), snapshot, 0755)
	}
}

func main() {
	verbose := flag.Bool("v", false, "verbose output")
	flag.Parse()

	file, err := os.Open("delta_data.bin")
	check(err)
	defer file.Close()

	buffer := bufio.NewReader(file)

	sizes := make([]float64, 0)
	speeds := make([]float64, 0)

	encode := qpc.NewHistory("encode")
	decode := qpc.NewHistory("decode")

	server := physics.NewState(901)
	client := physics.NewState(901)

	// initialize the base state
	for i := 0; i < 6; i += 1 {
		server.ReadNext(buffer)
		client.IncFrame()
		client.Current().Assign(server.Current())
	}

	frame := 6
	for {
		err = server.ReadNext(buffer)
		if err == io.EOF {
			break
		}
		check(err)
		frame += 1

		runtime.GC()

		// Server side
		encode.Start()
		snapshot := server.Encode()
		encode.Stop()
		// ===

		DebugSnapshot(frame, snapshot)

		runtime.GC()

		// Client side
		decode.Start()
		client.IncFrame()
		client.Decode(snapshot)
		decode.Stop()
		// ===

		size := float64(len(snapshot)*8) / 1000.0
		sizes = append(sizes, size)

		speed := size * 60.0
		speeds = append(speeds, speed)

		equal := server.Current().Equals(client.Current())
		if *verbose {
			if !equal {
				fmt.Print("! ")
			}
			fmt.Printf("%04d %8.3fkbps %10s %10s\n", frame, speed, encode.Last(), decode.Last())
		} else {
			if equal {
				fmt.Print(".")
			} else {
				fmt.Print("X")
			}
		}
	}

	fmt.Println()
	fmt.Printf("#%d %.3fkbps ±%.3fkbps\n", len(sizes), stats.Mean(speeds), stats.StdDevS(speeds))
	fmt.Println()

	fmt.Printf("MIN %10.3f kbps\n", stats.Min(speeds))
	for _, p := range []float64{5, 10, 25, 50, 75, 90, 95} {
		fmt.Printf("P%02.f %10.3f kbps\n", p, stats.Percentile(speeds, p))
	}
	fmt.Printf("MAX %10.3f kbps\n", stats.Max(speeds))

	fmt.Println()

	fmt.Printf("TOTAL  %10.3f kb\n", stats.Sum(sizes))
	fmt.Printf("  AVG  %10.3f kb per frame\n", stats.Mean(sizes))
	fmt.Printf("  AVG  %10.3f bits per cube\n", stats.Mean(sizes)*1000/float64(len(sizes)))

	fmt.Println()
	fmt.Println("TIMING:")
	qpc.PrintSummary(encode, decode)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
