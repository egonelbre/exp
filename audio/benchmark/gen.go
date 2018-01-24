package benchmark_audio

type Sampler func(buf *Buffer, generate func())

// Generete fills audio left, right channels and rest with 0-s
func Generate_Dynamic(buf *Buffer, generate func() float32) {
	nchan := int(buf.ChannelCount)
	maxg := 2
	if maxg > nchan {
		maxg = nchan
	}
	for i := 0; i < len(buf.Data); i += nchan {
		sample := generate()
		for k := 0; k < maxg; k++ {
			buf.Data[i+k] = sample
		}
		for k := maxg; k < nchan; k++ {
			buf.Data[i+k] = 0
		}
	}
}

// Generete fills audio channels
func Generate_Switch(buf *Buffer, generate func() float32) {
	switch buf.ChannelCount {
	case 1:
		for i := 0; i < len(buf.Data); i++ {
			sample := generate()
			buf.Data[i] = sample
		}
	case 2:
		for i := 0; i < len(buf.Data)-1; i += 2 {
			sample := generate()
			buf.Data[i] = sample
			buf.Data[i+1] = sample
		}
	case 5:
		for i := 0; i < len(buf.Data)-5; i += 5 {
			sample := generate()
			buf.Data[i] = sample
			buf.Data[i+1] = sample
			buf.Data[i+2] = 0
			buf.Data[i+3] = 0
			buf.Data[i+4] = 0
		}
	case 7:
		for i := 0; i < len(buf.Data)-6; i += 7 {
			sample := generate()
			buf.Data[i] = sample
			buf.Data[i+1] = sample
			buf.Data[i+2] = 0
			buf.Data[i+3] = 0
			buf.Data[i+4] = 0
			buf.Data[i+5] = 0
			buf.Data[i+6] = 0
		}
	}

}

// Sawtooth
func Sawtooth_Dynamic(buf *Buffer, sampleRate uint32, hz float32) {
	phase := float32(0.0)
	phaseSpeed := 2 * hz / float32(sampleRate)
	Generate_Dynamic(buf, func() float32 {
		phase += phaseSpeed
		if phase > 1 {
			phase -= 2.0
		}
		return phase
	})
}

// Sawtooth
func Sawtooth_Switch(buf *Buffer, sampleRate uint32, hz float32) {
	phase := float32(0.0)
	phaseSpeed := 2 * hz / float32(sampleRate)
	Generate_Switch(buf, func() float32 {
		phase += phaseSpeed
		if phase > 1 {
			phase -= 2.0
		}
		return phase
	})
}

// Sawtooth inlined
func Sawtooth_Baseline_Ch1(buf Buffer_Ch1, sampleRate uint32, hz float32) {
	phase := float32(0.0)
	phaseSpeed := 2 * hz / float32(sampleRate)
	for i := range buf {
		buf[i][0] = phase
		phase += phaseSpeed
		if phase > 1 {
			phase -= 2.0
		}
	}
}

// Sawtooth inlined
func Sawtooth_Baseline_Ch2(buf Buffer_Ch2, sampleRate uint32, hz float32) {
	phase := float32(0.0)
	phaseSpeed := 2 * hz / float32(sampleRate)
	for i := range buf {
		buf[i][0] = phase
		buf[i][1] = phase

		phase += phaseSpeed
		if phase > 1 {
			phase -= 2.0
		}
	}
}

// Sawtooth inlined
func Sawtooth_Baseline_Ch5(buf Buffer_Ch5, sampleRate uint32, hz float32) {
	phase := float32(0.0)
	phaseSpeed := 2 * hz / float32(sampleRate)
	for i := range buf {
		buf[i][0] = phase
		buf[i][1] = phase

		buf[i][2] = 0
		buf[i][3] = 0
		buf[i][4] = 0

		phase += phaseSpeed
		if phase > 1 {
			phase -= 2.0
		}
	}
}

// Sawtooth inlined
func Sawtooth_Baseline_Ch7(buf Buffer_Ch7, sampleRate uint32, hz float32) {
	phase := float32(0.0)
	phaseSpeed := 2 * hz / float32(sampleRate)
	for i := range buf {
		buf[i][0] = phase
		buf[i][1] = phase

		buf[i][2] = 0
		buf[i][3] = 0
		buf[i][4] = 0
		buf[i][5] = 0
		buf[i][6] = 0

		phase += phaseSpeed
		if phase > 1 {
			phase -= 2.0
		}
	}
}
