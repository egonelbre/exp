package benchmark_audio

type (
	// Buffer represents an interleaved block of samples
	Buffer struct {
		ChannelCount byte
		Data         []float32
	}

	// Buffer_Ch1 represents mono sound samples
	Buffer_Ch1 [][1]float32

	// Buffer_Ch2 represents stereo sound samples
	Buffer_Ch2 [][2]float32

	// Buffer_Ch5 represents surround sound samples
	Buffer_Ch5 [][5]float32

	// Buffer_Ch7 represents surround sound samples
	Buffer_Ch7 [][7]float32
)
