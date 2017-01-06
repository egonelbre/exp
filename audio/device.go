package audio

type DeviceInfo struct {
	Name              string
	ChannelCount      int
	SampleRate        int
	SamplesPerChannel int
}

type OutputDevice interface {
	OutputInfo() DeviceInfo
	Close() error
	Node
}

type InputDevice interface {
	InputInfo() DeviceInfo
	Close() error
	Node
}
