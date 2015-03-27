// +build ignore

// Response to Data Compression Challenge
// http://gafferongames.com/2015/03/14/the-networked-physics-data-compression-challenge/
/*
#2166 4016.563kbps ±1033.245kbps

MIN   1268.640 kbps
P05   1911.360 kbps
P10   2352.000 kbps
P25   3403.680 kbps
P50   4261.200 kbps
P75   5011.680 kbps
P90   5044.800 kbps
P95   5064.480 kbps
MAX   5157.600 kbps

TOTAL  144997.928 kb
  AVG      66.943 kb per frame
  AVG      30.906 bits per cube

TIMING:
                  MIN        10%        25%        50%        75%        90%        MAX
    encode  6.37552ms 9.357373ms 10.939417ms 15.087927ms 16.153194ms 16.481931ms 21.978885ms
    decode 6.170953ms 9.100995ms 10.726811ms 14.926909ms 16.004459ms 16.347935ms 22.231691ms
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/egonelbre/exp/physicscompress2/physics"
	"github.com/egonelbre/exp/qpc"

	"github.com/montanaflynn/stats"
)

func main() {
	verbose := flag.Bool("v", false, "verbose output")
	flag.Parse()

	file, err := os.Open("delta_data_round_two.bin")
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
