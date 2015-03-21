package arith

type Writer struct {
	encoder Encoder
	model   Model
	shift   uint
}

func NewWriter(model Model) *Writer {
	return &Writer{
		encoder: *NewEncoder(),
		model:   model,
		shift:   model.NBits(),
	}
}

func (w *Writer) Write(data []byte) (int, error) {
	shift := w.shift
	model := w.model
	encoder := &w.encoder

	switch shift {
	case 1:
		const mask = 1<<1 - 1
		for _, b := range data {
			model.Encode(encoder, uint(b>>7&mask))
			model.Encode(encoder, uint(b>>6&mask))
			model.Encode(encoder, uint(b>>5&mask))
			model.Encode(encoder, uint(b>>4&mask))
			model.Encode(encoder, uint(b>>3&mask))
			model.Encode(encoder, uint(b>>2&mask))
			model.Encode(encoder, uint(b>>1&mask))
			model.Encode(encoder, uint(b>>0&mask))
		}
	case 2:
		const mask = 1<<2 - 1
		for _, b := range data {
			model.Encode(encoder, uint(b>>6&mask))
			model.Encode(encoder, uint(b>>4&mask))
			model.Encode(encoder, uint(b>>2&mask))
			model.Encode(encoder, uint(b>>0&mask))
		}
	case 4:
		const mask = 1<<4 - 1
		for _, b := range data {
			model.Encode(encoder, uint(b>>4&mask))
			model.Encode(encoder, uint(b>>0&mask))
		}
	case 8:
		for _, b := range data {
			model.Encode(encoder, uint(b))
		}
	default:
		panic("unimplemented")
	}

	return len(data), nil
}

func (w *Writer) Close() error {
	w.encoder.Close()
	return nil
}

func (w *Writer) Bytes() []byte {
	return w.encoder.Bytes()
}

type Reader struct {
	decoder Decoder
	model   Model
	shift   uint

	bits  uint
	nbits uint
}

func NewReader(data []byte, model Model) *Reader {
	return &Reader{
		decoder: *NewDecoder(data),
		model:   model,
		shift:   model.NBits(),

		bits:  0,
		nbits: 0,
	}
}

func (r *Reader) Read(data []byte) (int, error) {
	model := r.model
	decoder := &r.decoder
	shift := model.NBits()

	next := 0
	for next < len(data) {
		value := model.Decode(decoder)
		r.bits <<= shift
		r.bits |= value
		r.nbits += shift

		if r.nbits >= 8 {
			r.nbits -= 8
			data[next] = byte(r.bits >> r.nbits)
			next += 1
			r.bits &= 1<<r.nbits - 1
		}
	}
	return next, nil
}
