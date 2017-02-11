package portaudio

import (
	"github.com/egonelbre/exp/audio"
	"github.com/egonelbre/exp/audio/slice"
	"github.com/gordonklaus/portaudio"
)

func init() {
	portaudio.Initialize()
}

type OutputDevice struct {
	stream *portaudio.Stream
	info   audio.DeviceInfo

	head   int
	buffer []float32
}

func (device *OutputDevice) OutputInfo() audio.DeviceInfo { return device.info }

func (device *OutputDevice) Process(buf audio.Buffer) error {
	_, err := device.Write(buf)
	return err
}

func (device *OutputDevice) Write(buf audio.Buffer) (int, error) {
	nchan := device.info.ChannelCount

	// TODO: support all conversions
	if buf.ChannelCount() != nchan {
		return 0, audio.ErrUnknownBuffer
	}

	src := buf.ShallowCopy()
	frameCount := buf.FrameCount()

	frames := 0
	for frames < frameCount {
		n := slice.Interleave32(src, nchan, device.buffer[device.head:])
		frames += n
		src.CutLeading(n)
		device.head += n * nchan

		if device.head >= len(device.buffer) {
			device.head = 0
			if err := device.stream.Write(); err != nil {
				return frames, err
			}
		}
	}

	return frames, nil
}

func (device *OutputDevice) Close() error {
	return device.stream.Stop()
}

func NewOutputDevice(pref audio.DeviceInfo) (audio.OutputDevice, error) {
	var err error
	device := &OutputDevice{}

	device.info = pref
	device.buffer = make([]float32, pref.ChannelCount*pref.SamplesPerChannel)

	device.stream, err = portaudio.OpenDefaultStream(
		0, int(pref.ChannelCount),
		float64(pref.SampleRate),
		len(device.buffer),
		&device.buffer,
	)
	if err != nil {
		return nil, err
	}

	return device, device.stream.Start()
}
