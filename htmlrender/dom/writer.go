package dom

type Writer interface {
	Open(tag string)
	Attr(name, value string)
	Text(text string)
	Close(tag string)

	Wrap(r RendererExtended) func()
	Render(r Renderer)
	RenderAll(rs ...Renderer)

	Bytes() []byte
	String() string
}

type UnsafeWriter interface {
	Writer
	CloseAttributes()
	UnsafeWrite(text string)
	UnsafeContent(text string)
}

type Renderer interface {
	Render(w Writer)
}

type RendererExtended interface {
	Renderer
	RenderOpen(w Writer)
	RenderClose(w Writer)
}

type RenderFunc func(w Writer)

func (fn RenderFunc) Render(w Writer) { fn(w) }

func (w *writer) Wrap(r RendererExtended) func() {
	r.RenderOpen(w)
	return func() {
		r.RenderClose(w)
	}
}

func (w *writer) Render(r Renderer) {
	r.Render(w)
}

func (w *writer) RenderAll(rs ...Renderer) {
	for _, r := range rs {
		r.Render(w)
	}
}
