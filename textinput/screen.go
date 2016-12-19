package main

import (
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Screen struct {
	Tick float64
	Font *Font
	Text string
}

func NewScreen() *Screen {
	screen := &Screen{}
	screen.Font = ArcadeFont.Load()
	return screen
}

func (screen *Screen) KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Release {
		return
	}

	if key == glfw.KeyEnter {
		screen.Text += "\n"
		return
	}

	if key == glfw.KeyBackspace {
		if len(screen.Text) > 0 {
			screen.Text = screen.Text[:len(screen.Text)-1]
		}
		return
	}

	if glfw.KeyA <= key && key <= glfw.KeyZ {
		screen.Text += string(key - glfw.KeyA + 'a')
		return
	}

	if glfw.Key0 <= key && key <= glfw.Key9 {
		screen.Text += string(key - glfw.Key0 + '0')
		return
	}

	switch key {
	case glfw.KeySpace:
		screen.Text += " "
	case glfw.KeyPeriod:
		screen.Text += "."
	case glfw.KeyComma:
		screen.Text += ","
	case glfw.KeySemicolon:
		screen.Text += ";"
	case glfw.KeySlash:
		screen.Text += "/"
	case glfw.KeyEqual:
		screen.Text += "="
	}
}

func (screen *Screen) Update(window *glfw.Window) {
	now := time.Now()

	width, height := window.GetSize()
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Ortho(0, float64(width), float64(height), 0, 30, -30)

	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.ALPHA_TEST)

	text := screen.Text
	if now.Unix()&1 == 1 {
		text += "|"
	}
	screen.Font.DrawText(Point{1, 1}, 74, text)
}
