package ui

import termbox "github.com/nsf/termbox-go"

type Button struct {
	Caption string
	Click   func(form *Form)
}

func (button *Button) Focusable() bool { return true }

func (button *Button) Handle(form *Form, action Action) {
	switch action := action.(type) {
	case KeyPress:
		switch action.Key {
		case termbox.KeyEnter:
			button.Click(form)
		}
	}
}

func (button *Button) Unserialize(data []byte) error {
	button.Caption = string(data)
	return nil
}
func (button *Button) Serialize() ([]byte, error) {
	return []byte(button.Caption), nil
}

func (button *Button) Render(form *Form, component *Component) {
	form.DrawText(component.Rect, button.Caption, component.Focused)
}
