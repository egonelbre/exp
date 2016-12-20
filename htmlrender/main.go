package main

import (
	"fmt"

	"github.com/egonelbre/exp/uirender/ui"
)

type Input struct {
	ID          string
	Placeholder string
}

func (input Input) Render(w ui.Writer) {
	w.Open("input")
	if input.ID != "" {
		w.Attr("id", input.ID)
		w.Attr("name", input.ID)
	}

	ui.Class{"mdl-input"}.Render(w)

	if input.Placeholder != "" {
		w.Attr("placeholder", input.Placeholder)
	}

	w.Close("input")
}

func main() {
	writer := ui.NewWriter()

	writer.Render(ui.Form{
		ui.Class{"example", "test"},
		ui.Method{"POST"},

		Input{"first", "First Name"},
		Input{"last", "Last Name"},
	})

	fmt.Println(writer.String())
}
