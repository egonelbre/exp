package audio

type DeviceInfo struct {
	Name              string
	ChannelCount      int
	SampleRate        int
	SamplesPerChannel int // TODO: rename to BufferFrameCount
}

type OutputDevice interface {
	OutputInfo() DeviceInfo
	Close() error
	Writer
	Processor
}

type InputDevice interface {
	InputInfo() DeviceInfo
	Close() error
	Reader
}
