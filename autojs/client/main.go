package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func main() {
	vecty.SetTitle("Auto-reload demo")
	vecty.RenderBody(&App{})
}

type App struct {
	vecty.Core
}

func (app *App) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(vecty.Text("CHANGE ME")),
	)
}
