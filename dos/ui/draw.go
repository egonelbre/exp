package ui

import termbox "github.com/nsf/termbox-go"

func (screen *Screen) DrawFlush() {
	err := termbox.Flush()
	if err != nil {
		// TODO: handle error properly
		panic(err)
	}
}

func (screen *Screen) DrawCell(x, y int, cell termbox.Cell) {
	termbox.SetCell(x, y, cell.Ch, cell.Fg, cell.Bg)
}

func (form *Form) DrawFlush() {
	form.Screen.DrawFlush()
}

func (form *Form) DrawCell(x, y int, cell termbox.Cell) {
	if x < 0 || form.BoundsRect.Height <= x ||
		y < 0 || form.BoundsRect.Width <= y {
		return
	}

	form.Screen.DrawCell(x+form.BoundsRect.Left, y+form.BoundsRect.Top, cell)
}

func (form *Form) DrawBlock(r Rect, cell termbox.Cell) {
	for x := r.Left; x < r.Left+r.Width; x += 1 {
		for y := r.Top; y < r.Top+r.Height; y += 1 {
			form.DrawCell(x, y, cell)
		}
	}
}

func (form *Form) DrawBorder(r Rect) {
	form.DrawBlock(Rect{r.Left, r.Top, r.Width, 1}, termbox.Cell{'─', termbox.ColorDefault, termbox.ColorDefault})
	form.DrawBlock(Rect{r.Left, r.Top + r.Height - 1, r.Width, 1}, termbox.Cell{'─', termbox.ColorDefault, termbox.ColorDefault})
	form.DrawBlock(Rect{r.Left, r.Top, 1, r.Height}, termbox.Cell{'│', termbox.ColorDefault, termbox.ColorDefault})
	form.DrawBlock(Rect{r.Left + r.Width - 1, r.Top, 1, r.Height}, termbox.Cell{'│', termbox.ColorDefault, termbox.ColorDefault})

	form.DrawCell(r.Left, r.Top, termbox.Cell{'┌', termbox.ColorDefault, termbox.ColorDefault})
	form.DrawCell(r.Left+r.Width, r.Top, termbox.Cell{'┐', termbox.ColorDefault, termbox.ColorDefault})

	form.DrawCell(r.Left, r.Top+r.Height, termbox.Cell{'└', termbox.ColorDefault, termbox.ColorDefault})
	form.DrawCell(r.Left+r.Width, r.Top+r.Height, termbox.Cell{'┘', termbox.ColorDefault, termbox.ColorDefault})
}

func (form *Form) DrawText(r Rect, t string, active bool) {
	fg, bg := termbox.ColorDefault, termbox.ColorDefault
	if active {
		fg, bg = termbox.ColorBlack, termbox.ColorWhite
	}
	form.DrawBlock(r, termbox.Cell{' ', fg, bg})

	x0 := r.Left
	x1 := r.Left + r.Width
	for _, ch := range t {
		//TODO: implement wrap
		if x0 > x1 {
			break
		}

		form.DrawCell(x0, r.Top, termbox.Cell{ch, fg, bg})
		x0 += 1
	}
}
