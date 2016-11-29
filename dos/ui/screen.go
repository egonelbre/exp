package ui

import termbox "github.com/nsf/termbox-go"

func (screen *Screen) Start() {
	if screen.Actions == nil {
		screen.Actions = make(chan Action)
	}

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)

	go screen.Listen()
}

func (screen *Screen) Stop() {
	termbox.Close()
}

func (screen *Screen) Listen() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC:
				close(screen.Actions)
				return
			default:
				screen.Actions <- KeyPress{ev.Ch, ev.Key}
			}
		case termbox.EventError:
			//TODO: proper error handling
			panic(ev.Err)
		}
	}
}
