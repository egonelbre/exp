package ui

import "strings"

type Div []Renderer

func (el Div) Render(w Writer) {
	el.RenderOpen(w)
	el.RenderClose(w)
}

func (el Div) RenderOpen(w Writer) {
	w.Open("div")
	w.RenderAll(el...)
}

func (el Div) RenderClose(w Writer) {
	w.Close("div")
}

type Span []Renderer

func (el Span) Render(w Writer) {
	el.RenderOpen(w)
	el.RenderClose(w)
}

func (el Span) RenderOpen(w Writer) {
	w.Open("div")
	w.RenderAll(el...)
}

func (el Span) RenderClose(w Writer) {
	w.Close("div")
}

type Text struct{ Content string }

func (text Text) Render(w Writer) { w.Text(text.Content) }

type Attrs map[string]string

func (attrs Attrs) Render(w Writer) {
	for name, value := range attrs {
		w.Attr(name, value)
	}
}

type Attr struct{ Name, Value string }

func (attr Attr) Render(w Writer) {
	w.Attr(attr.Name, attr.Value)
}

type Children []Renderer

func (children Children) Render(w Writer) {
	for _, child := range children {
		child.Render(w)
	}
}

type URL struct{ Value string }

func (attr URL) Render(w Writer) { w.Attr("href", attr.Value) }

type Class []string

func (class Class) Render(w Writer) {
	w.Attr("class", strings.Join([]string(class), " "))
}
