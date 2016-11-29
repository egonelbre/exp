package ui

type Screen struct {
	Forms   []*Form
	Actions chan Action
}

func (screen *Screen) NewForm() *Form {
	return &Form{Screen: screen}
}

func (screen *Screen) Push(form *Form) { screen.Forms = append(screen.Forms, form) }
func (screen *Screen) Pop()            { screen.Forms = screen.Forms[:len(screen.Forms)-1] }

type Form struct {
	BoundsRect Rect
	ClientRect Rect

	Screen *Screen

	Close bool
	Save  bool

	Record Record

	Components   []*Component
	FocusedIndex int
}

type Action interface{}

type Record interface {
	Caption() string

	Load() error
	Save() error
}

type Component struct {
	Rect
	Focused bool
	Widget  Widget
}

type Widget interface {
	Focusable() bool
	Handle(screen *Form, action Action)

	Unserialize(data []byte) error
	Serialize() ([]byte, error)

	Render(screen *Form, component *Component)
}
