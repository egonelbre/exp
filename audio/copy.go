package audio

type Pipe struct {
	Reader    Reader
	Processor Processor
	Writer    Writer
}

func (p *Pipe) Step(buf Buffer) error {
	var readn, writen int
	var readerr, processerr, writeerr error

	partial := buf.ShallowCopy()

	if p.Reader != nil {
		readn, readerr = p.Reader.Read(partial)
		partial.Slice(0, readn)
	}

	if p.Processor != nil {
		processerr = p.Processor.Process(partial)
		if processerr != nil {
			return processerr
		}
	}

	if p.Writer != nil {
		for !partial.Empty() {
			writen, writeerr = p.Writer.Write(partial)
			if writeerr != nil {
				return writeerr
			}
			partial.Slice(writen, partial.FrameCount())
		}
	}

	return readerr
}

func (p *Pipe) Run(buf Buffer) error {
	var err error
	for err == nil {
		err = p.Step(buf)
	}
	return err
}

func Copy(dst Writer, src Reader, buffer Buffer) error {
	p := Pipe{src, nil, dst}
	return p.Run(buffer)
}
