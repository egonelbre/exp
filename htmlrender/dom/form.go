package dom

type Name struct{ Value string }

func (attr Name) Render(w Writer) { w.Attr("name", attr.Value) }

type Method struct{ Value string }

func (attr Method) Render(w Writer) { w.Attr("method", attr.Value) }

type Form []Renderer

func (el Form) Render(w Writer) {
	el.RenderOpen(w)
	el.RenderClose(w)
}

func (el Form) RenderOpen(w Writer) {
	w.Open("form")
	w.RenderAll(el...)
}

func (el Form) RenderClose(w Writer) {
	w.Close("form")
}
