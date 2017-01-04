package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"

	"github.com/egonelbre/exp/audio"
	"github.com/egonelbre/exp/audio/example/internal/effect"
	"github.com/egonelbre/exp/audio/example/internal/osc"
	"github.com/gordonklaus/portaudio"
)

func main() {
	sine := &osc.Sine{}
	sine.Frequency.Set(440.0)

	gain := &effect.Gain{}
	gain.Value.Set(0.5)

	go func() {
		runtime.LockOSThread()

		work := audio.NewBuffer32(audio.Format{
			Channels:   1,
			SampleRate: 44100,
		}, 512)

		portaudio.Initialize()
		defer portaudio.Terminate()

		stream, err := portaudio.OpenDefaultStream(
			0, int(work.Channels),
			float64(work.SampleRate),
			len(work.Data), &work.Data,
		)
		if err != nil {
			log.Fatal(err)
		}

		if err := stream.Start(); err != nil {
			log.Fatal(err)
		}

		defer stream.Stop()

		for {
			sine.Process32(work)
			gain.Process32(work)

			if err := stream.Write(); err != nil {
				log.Fatal(err)
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			k := scanner.Text()[0]
			switch k {
			case 'q':
				return
			case '+':
				gain.Value.Set(gain.Value.Get() + 0.10)
			case '-':
				gain.Value.Set(gain.Value.Get() - 0.10)
			default:
				v := float64(math.Abs(float64(int(k - 100))))
				note := 440.0 * math.Pow(2, (v)/12.0)
				if note > 22000 {
					note = 440.0
				}

				fmt.Printf("switching oscillator to %.2f Hz\n", note)
				sine.Frequency.Set(note)
			}
		}
	}
}
