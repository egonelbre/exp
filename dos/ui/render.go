package ui

import termbox "github.com/nsf/termbox-go"

func (form *Form) Init() {
	widgets := ReflectLoadWidgets(form.Record)

	form.Components = nil
	form.FocusedIndex = -1

	form.Components = nil
	for _, w := range widgets {
		form.Components = append(form.Components,
			&Component{Widget: w})
	}
	form.UpdateLayout()

	form.TabFocus(1)
}

func (form *Form) UpdateLayout() {
	form.ClientRect = form.BoundsRect

	form.ClientRect.Left += 1
	form.ClientRect.Top += 3
	form.ClientRect.Width -= 2
	form.ClientRect.Height -= 4

	box := form.ClientRect
	box.Height = 1
	for _, component := range form.Components {
		component.Rect = box
		box.Top += 1
	}
}

func (form *Form) Render() {
	r := Rect{0, 0, form.BoundsRect.Width, form.BoundsRect.Height}
	form.DrawBlock(r, termbox.Cell{Ch: ' '})
	form.DrawBorder(r)

	r = Rect{1, 1, form.BoundsRect.Width - 2, 1}
	form.DrawText(r, form.Record.Caption(), false)

	for _, component := range form.Components {
		component.Widget.Render(form, component)
	}

	form.DrawFlush()
}
