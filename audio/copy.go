package audio

func CopyBuffer32(dst, src Node32, format Format, bufferSamples int) error {
	buf := NewBuffer32Samples(format, bufferSamples)
	for {
		partial := *buf
		srcn, srcerr := src.Process32(&partial)
		partial.Data = partial.Data[:srcn]

		for len(partial.Data) > 0 {
			writen, dsterr := dst.Process32(&partial)
			if dsterr != nil {
				return dsterr
			}
			partial.Data = partial.Data[writen:]
		}

		if srcerr != nil {
			return srcerr
		}
	}
}

func CopyBuffer64(dst, src Node64, format Format, bufferSamples int) error {
	buf := NewBuffer64Samples(format, bufferSamples)
	for {
		partial := *buf
		srcn, srcerr := src.Process64(&partial)
		partial.Data = partial.Data[:srcn]

		for len(partial.Data) > 0 {
			writen, dsterr := dst.Process64(&partial)
			if dsterr != nil {
				return dsterr
			}
			partial.Data = partial.Data[writen:]
		}

		if srcerr != nil {
			return srcerr
		}
	}
}
