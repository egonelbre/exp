package portaudio

import (
	"time"

	"github.com/egonelbre/exp/audio"
	"github.com/gordonklaus/portaudio"
)

func init() {
	portaudio.Initialize()
}

type OutputDevice struct {
	stream     *portaudio.Stream
	backbuffer audio.Buffer32
	info       audio.DeviceInfo
}

func (device *OutputDevice) Position() time.Duration { return 0 }

func (device *OutputDevice) OutputInfo() audio.DeviceInfo {
	return device.info
}

func (device *OutputDevice) Write(frame audio.Frame) error {
	f, ok := frame.(*Frame)
	if !ok ||
		f.buffer.Format != device.backbuffer.Format ||
		len(f.buffer.Data) != len(device.backbuffer.Data) {
		return audio.ErrInvalidFrame
	}

	for i := range f.buffer.Data[f.head:] {
		f.buffer.Data[f.head+i] = 0
	}

	copy(device.backbuffer.Data, f.buffer.Data)
	f.head = 0

	return device.stream.Write()
}

func (device *OutputDevice) Close() error {
	return device.stream.Stop()
}

type Frame struct {
	head   int
	buffer audio.Buffer32
}

func (device *OutputDevice) NewFrame() audio.Frame {
	return &Frame{0, *device.backbuffer.Clone()}
}

func (frame *Frame) Done() bool {
	return frame.head >= len(frame.buffer.Data)
}

func (frame *Frame) Offset() time.Duration { return 0 }

func (frame *Frame) Seek(offset time.Duration, whence int) (time.Duration, error) {
	panic("TODO")
	return 0, nil
}

func (frame *Frame) Position() time.Duration {
	return frame.buffer.Format.BufferDuration(frame.head)
}

func (frame *Frame) Duration() time.Duration {
	return frame.buffer.Format.BufferDuration(len(frame.buffer.Data))
}

func (frame *Frame) Process32(buf *audio.Buffer32) (int, error) {
	// TODO: support all conversions
	if buf.Format != frame.buffer.Format {
		return 0, audio.ErrInvalidFrame
	}

	n := copy(frame.buffer.Data[frame.head:], buf.Data)
	frame.head += n

	return n, nil
}

func (frame *Frame) Process64(buf *audio.Buffer64) (int, error) {
	return 0, audio.ErrIncompatibleFormat
}

func NewOutputDevice(pref audio.DeviceInfo) (audio.OutputDevice, error) {
	var err error
	device := &OutputDevice{}

	device.info = pref

	device.backbuffer = *audio.NewBuffer32Samples(audio.Format{
		ChannelCount: pref.ChannelCount,
		SampleRate:   pref.SampleRate,
	}, pref.SamplesPerChannel)

	device.stream, err = portaudio.OpenDefaultStream(
		0, int(pref.ChannelCount),
		float64(pref.SampleRate),
		len(device.backbuffer.Data),
		&device.backbuffer.Data,
	)
	if err != nil {
		return nil, err
	}

	return device, device.stream.Start()
}
