package main

import (
	"fmt"

	"github.com/egonelbre/exp/htmlrender/dom"
)

var CustomTemplate = dom.MustTemplate(`
	<div id="{{.id}}">
		<span>{{.glyph}}</span>
		<span>{{.title}}</span>
		<span>{{ Render .child }}</span>
	</div>
`)

type Input struct {
	ID          string
	Placeholder string
}

func (input Input) Render(w dom.Writer) {
	w.Open("input")
	if input.ID != "" {
		w.Attr("id", input.ID)
		w.Attr("name", input.ID)
	}

	dom.Class{"mdl-input"}.Render(w)

	if input.Placeholder != "" {
		w.Attr("placeholder", input.Placeholder)
	}

	w.Close("input")
}

func main() {
	writer := dom.NewWriter()

	writer.Render(dom.Form{
		dom.Class{"example", "test"},
		dom.Method{"POST"},

		Input{"first", "First Name"},
		Input{"last", "Last Name"},

		CustomTemplate.Renderer(map[string]interface{}{
			"id":    "clicky",
			"glyph": "pen",
			"title": "clicky",
			"child": Input{"middle", "Middle Name"},
		}),
	})

	fmt.Println(writer.String())
}
