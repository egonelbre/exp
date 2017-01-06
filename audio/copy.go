package audio

func CopyBuffer32(dst StreamWriter, src StreamReader, format Format, bufferSamples int) error {
	in := src.NewFrame()
	buf := NewBuffer32Samples(format, bufferSamples)
	out := dst.NewFrame()

	for {
		srcerr := src.Read(in)
		for !in.Done() {
			partial := *buf

			// read as much of input frame into a temporary buffer
			readn, inerr := in.Process32(&partial)
			partial.Data = partial.Data[:readn]

			// write the read data to output
			for len(partial.Data) > 0 {
				// write as much to the frame as possible
				writen, outerr := out.Process32(&partial)
				partial.Data = partial.Data[writen:]
				if outerr != nil {
					return outerr
				}

				// is the output frame complete?
				if out.Done() {
					// write the frame
					dsterr := dst.Write(out)
					if dsterr != nil {
						return dsterr
					}
				}
			}

			if inerr != nil {
				dst.Write(out)
				return inerr
			}
		}

		if srcerr != nil {
			dst.Write(out)
			return srcerr
		}
	}
}

func CopyBuffer64(dst StreamWriter, src StreamReader, format Format, bufferSamples int) error {
	in := src.NewFrame()
	buf := NewBuffer64Samples(format, bufferSamples)
	out := dst.NewFrame()

	for {
		srcerr := src.Read(in)
		for !in.Done() {
			partial := *buf

			// read as much of input frame into a temporary buffer
			readn, inerr := in.Process64(&partial)
			partial.Data = partial.Data[:readn]

			// write the read data to output
			for len(partial.Data) > 0 {
				// write as much to the frame as possible
				writen, outerr := out.Process64(&partial)
				partial.Data = partial.Data[writen:]
				if outerr != nil {
					return outerr
				}

				// is the output frame complete?
				if out.Done() {
					// write the frame
					dsterr := dst.Write(out)
					if dsterr != nil {
						return dsterr
					}
				}
			}

			if inerr != nil {
				dst.Write(out)
				return inerr
			}
		}

		if srcerr != nil {
			dst.Write(out)
			return srcerr
		}
	}
}
