package native

import (
	"github.com/egonelbre/exp/audio"
	"github.com/egonelbre/exp/audio/native/portaudio"
)

func NewOutputDevice(pref audio.DeviceInfo) (audio.OutputDevice, error) {
	return portaudio.NewOutputDevice(pref)
}
