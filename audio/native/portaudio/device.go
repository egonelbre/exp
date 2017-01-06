package portaudio

import (
	"github.com/egonelbre/exp/audio"
	"github.com/gordonklaus/portaudio"
)

func init() {
	portaudio.Initialize()
}

type OutputDevice struct {
	stream *portaudio.Stream
	info   audio.DeviceInfo

	head   int
	buffer audio.Buffer32
}

func (device *OutputDevice) OutputInfo() audio.DeviceInfo { return device.info }

func (device *OutputDevice) Process32(buf *audio.Buffer32) (int, error) {
	// TODO: support all conversions
	if buf.Format != device.buffer.Format {
		return 0, audio.ErrInvalidFrame
	}

	partial := buf.Data
	total := 0
	for len(partial) > 0 {
		n := copy(device.buffer.Data[device.head:], partial)
		partial = partial[n:]

		total += n
		device.head += n

		if device.head >= len(device.buffer.Data) {
			device.head = 0
			if err := device.stream.Write(); err != nil {
				return total, err
			}
		}
	}

	return total, nil
}

func (device *OutputDevice) Process64(buf *audio.Buffer64) (int, error) {
	panic("TODO")
	return 0, audio.ErrIncompatibleFormat
}

func (device *OutputDevice) Close() error {
	return device.stream.Stop()
}

func NewOutputDevice(pref audio.DeviceInfo) (audio.OutputDevice, error) {
	var err error
	device := &OutputDevice{}

	device.info = pref

	device.buffer = *audio.NewBuffer32Samples(audio.Format{
		ChannelCount: pref.ChannelCount,
		SampleRate:   pref.SampleRate,
	}, pref.SamplesPerChannel)

	device.stream, err = portaudio.OpenDefaultStream(
		0, int(pref.ChannelCount),
		float64(pref.SampleRate),
		len(device.buffer.Data),
		&device.buffer.Data,
	)
	if err != nil {
		return nil, err
	}

	return device, device.stream.Start()
}
