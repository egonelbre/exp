package audio

func Copy(dst Writer, src Reader, buffer Buffer) error {
	for {
		partial := buffer.ShallowCopy()
		srcn, srcerr := src.Read(partial)
		partial.Slice(0, srcn)

		for !partial.Empty() {
			writen, dsterr := dst.Write(partial)
			if dsterr != nil {
				return dsterr
			}
			partial.Slice(writen, partial.FrameCount())
		}

		if srcerr != nil {
			return srcerr
		}
	}
}
