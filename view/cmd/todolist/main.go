package main

import ui "github.com/egonelbre/exp/view/cli"

type Item struct {
	Done  bool
	Value string
}

type App struct {
	Items []Item
}

func (app *App) Delete(i int) {
	app.Items = append(app.Items[:i], app.Items[i+1:])
}

func (app *App) RenderItem(i int) ui.Element {
	item := app.Items[i]
	return ui.Node(
		"item",
		ui.Text(item.Value),
		ui.Button(
			"Delete",
			ui.On("click", func(s *ui.State) {
				app.Delete(i)
			}),
		),
	)
}

func (app *App) Render() ui.Element {
	return ui.Node(
		"application",
		ui.Class("application"),
		ui.Input(
			"",
			ui.Ref("new-item"),
			ui.Attr{"placeholder": "Todo item..."},
		),
		ui.Button(
			"Add",
			ui.On("click", func(s *ui.State) {
				input := s.ByRef("new-item")
				text := input.Attr("value")
				app.Items = append(app.Items, Item{Value: text})
			}),
		),
		ui.Each(len(app.Items), app.RenderItem),
	)
}

func main() {

}
