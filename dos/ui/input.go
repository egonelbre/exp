package ui

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

type Input struct {
	Format string
	Text   string
}

func (input *Input) Focusable() bool { return true }

func (input *Input) Handle(form *Form, action Action) {
	switch action := action.(type) {
	case KeyPress:
		switch action.Key {
		case termbox.KeyBackspace:
			if len(input.Text) > 0 {
				input.Text = input.Text[:len(input.Text)-1]
			}
		default:
			// TODO: check whether it's a character
			input.Text = input.Text + string([]rune{action.Char})
		}
	default:
		// TODO: log ignored action
	}
}

func (input *Input) Unserialize(data []byte) error {
	input.Text = string(data)
	return nil
}
func (input *Input) Serialize() ([]byte, error) {
	return []byte(input.Text), nil
}

func (input *Input) Render(form *Form, component *Component) {
	text := input.Text
	if input.Format != "" {
		text = fmt.Sprintf(input.Format, input.Text)
	}
	form.DrawText(component.Rect, text, component.Focused)
}
