package ui

type Label struct {
	Caption string
}

func (label *Label) Focusable() bool                  { return false }
func (label *Label) Handle(form *Form, action Action) {}

func (label *Label) Unserialize(data []byte) error {
	label.Caption = string(data)
	return nil
}
func (label *Label) Serialize() ([]byte, error) {
	return []byte(label.Caption), nil
}

func (label *Label) Render(form *Form, component *Component) {
	form.DrawText(component.Rect, label.Caption, false)
}
