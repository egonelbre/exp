package ui

import termbox "github.com/nsf/termbox-go"

func (form *Form) Show() {
	err := form.Record.Load()
	if err != nil {
		//TODO: handle error properly
		panic(err)
	}

	form.Init()

	defer func() {
		if form.Save {
			if err := form.Record.Save(); err != nil {
				//TODO: handle error properly
				panic(err)
			}
		}
	}()

	if len(form.Screen.Actions) == 0 {
		form.Render()
	}

	for action := range form.Screen.Actions {
		form.Handle(action)

		if len(form.Screen.Actions) == 0 {
			form.Render()
		}

		if form.Close {
			break
		}
	}
}

func (form *Form) TabFocus(delta int) {
	base := form.FocusedIndex
	if base < 0 {
		base = 0
	}

	for _, component := range form.Components {
		component.Focused = false
	}

	for offset := 0; offset < len(form.Components); offset++ {
		k := (base + delta*offset + len(form.Components)) % len(form.Components)
		if k == form.FocusedIndex {
			continue
		}

		component := form.Components[k]
		if component.Widget.Focusable() {
			form.FocusedIndex = k
			component.Focused = true
			return
		}
	}

	form.FocusedIndex = -1
}

func (form *Form) Handle(action Action) {
	switch action := action.(type) {
	case KeyPress:
		switch action.Key {
		// focus handling
		case termbox.KeyArrowUp:
			form.TabFocus(-1)
			return
		case termbox.KeyTab, termbox.KeyArrowDown: // TODO: add ShiftTab
			form.TabFocus(1)
			return

		// closing/saving
		case termbox.KeyEsc:
			form.Close = true
			return
		case termbox.KeyCtrlS:
			form.Close = true
			form.Save = true
			return
		}
	}

	if 0 <= form.FocusedIndex && form.FocusedIndex < len(form.Components) {
		focused := form.Components[form.FocusedIndex]
		focused.Widget.Handle(form, action)
	}
}
