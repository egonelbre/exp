package main

import (
	"log"

	"github.com/egonelbre/exp/audio"
	"github.com/egonelbre/exp/audio/example/internal/effect"
	"github.com/egonelbre/exp/audio/native"
)

func main() {
	reader := NewReader(440, 460, 512, 531, 713, 734)
	output, err := native.NewOutputDevice(audio.DeviceInfo{
		ChannelCount:      2,
		SampleRate:        44100,
		SamplesPerChannel: 128,
	})
	if err != nil {
		log.Fatal(err)
	}

	gain := effect.NewGain(0.5)

	buf := audio.NewBufferF32Frames(audio.Format{44100, 2}, 128)
	pipe := audio.Pipe{reader, gain, output}
	err = pipe.Run(buf)
	if err != nil {
		log.Fatal(err)
	}
}
