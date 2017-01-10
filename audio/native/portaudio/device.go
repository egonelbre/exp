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
	buffer audio.BufferF32
}

func (device *OutputDevice) OutputInfo() audio.DeviceInfo { return device.info }

func (device *OutputDevice) Process(buf audio.Buffer) error {
	_, err := device.Write(buf)
	return err
}

func (device *OutputDevice) Write(buf audio.Buffer) (int, error) {
	// TODO: support all conversions
	if buf.ChannelCount() != device.buffer.ChannelCount() ||
		buf.FrameCount() != device.buffer.FrameCount() {
		return 0, audio.ErrUnknownBuffer
	}

	switch buf := buf.(type) {
	case *audio.BufferF32:
		head := 0
		for head < len(buf.Data) {
			n := copy(device.buffer.Data[device.head:], buf.Data[head:])
			head += n
			device.head += n

			if device.head >= len(device.buffer.Data) {
				device.head = 0
				if err := device.stream.Write(); err != nil {
					return head / buf.ChannelCount(), err
				}
			}
		}
		return head / buf.ChannelCount(), nil

	case *audio.BufferF64:
		head := 0
		for head < len(buf.Data) {
			n := len(device.buffer.Data[device.head:])
			if n < len(buf.Data[head:]) {
				n = len(buf.Data[head:])
			}

			for i := 0; i < n; i++ {
				device.buffer.Data[device.head+i] = float32(buf.Data[head+i])
			}

			head += n
			device.head += n

			if device.head >= len(device.buffer.Data) {
				device.head = 0
				if err := device.stream.Write(); err != nil {
					return head / buf.ChannelCount(), err
				}
			}
		}
		return head / buf.ChannelCount(), nil
	default:
		return 0, audio.ErrUnknownBuffer
	}
}

func (device *OutputDevice) Close() error {
	return device.stream.Stop()
}

func NewOutputDevice(pref audio.DeviceInfo) (audio.OutputDevice, error) {
	var err error
	device := &OutputDevice{}

	device.info = pref

	device.buffer = *audio.NewBufferF32Frames(audio.Format{
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
