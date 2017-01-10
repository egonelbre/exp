package main

import (
	"log"

	"github.com/egonelbre/exp/audio"
	"github.com/egonelbre/exp/audio/native"
)

func main() {
	dec := NewReader(440, 460, 512, 531, 713, 734)
	out, err := native.NewOutputDevice(audio.DeviceInfo{
		ChannelCount:      2,
		SampleRate:        44100,
		SamplesPerChannel: 128,
	})
	if err != nil {
		log.Fatal(err)
	}

	buf := audio.NewBufferF32Frames(audio.Format{44100, 2}, 128)
	err = audio.Copy(out, dec, buf)
	if err != nil {
		log.Fatal(err)
	}
}
