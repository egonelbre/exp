package ui

import termbox "github.com/nsf/termbox-go"

type KeyPress struct {
	Char rune
	Key  termbox.Key
}
