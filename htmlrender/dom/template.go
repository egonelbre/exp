package dom

import (
	"bytes"
	"errors"
	"html/template"
	"sync"
)

var errUnsafeWriterNotSupported = errors.New("Unsafe writing unsupported.")

type Template struct {
	Content string

	once sync.Once
	tmpl *template.Template
	err  error
}

func NewTemplate(content string) (*Template, error) {
	t := &Template{Content: content}
	t.compile()
	return t, t.err
}

func MustTemplate(content string) *Template {
	t, err := NewTemplate(content)
	if err != nil {
		panic(err)
	}
	return t
}

func (t *Template) Renderer(data interface{}) Renderer {
	return RenderFunc(func(w Writer) {
		t.RenderData(data, w)
	})
}

func (t *Template) compile() {
	t.once.Do(func() {
		tmpl, err := template.New("").Funcs(template.FuncMap{
			"Render": RenderHTML,
		}).Parse(t.Content)

		if err != nil {
			t.err = err
			return
		}

		t.tmpl = tmpl
	})
}

func (t *Template) RenderData(data interface{}, w Writer) {
	t.compile()

	if t.err != nil {
		Error{t.err}.Render(w)
		return
	}

	uw, ok := w.(UnsafeWriter)
	if !ok {
		Error{errUnsafeWriterNotSupported}.Render(w)
		return
	}

	var buf bytes.Buffer
	err := t.tmpl.Execute(&buf, data)
	if err != nil {
		Error{err}.Render(w)
		return
	}

	uw.UnsafeContent(buf.String())
}

func RenderHTML(r Renderer) template.HTML {
	w := NewWriter()
	r.Render(w)
	return template.HTML(w.String())
}

// type TemplateFile
