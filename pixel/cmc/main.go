package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const Detail = 1 << 16

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "CMC!",
		Bounds:    pixel.R(0, 0, 480, 320),
		Resizable: true,
		VSync:     true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	r := rand.New(rand.NewSource(0))

	canvas := NewImage()
	backbuffer := pixel.MakePictureData(pixel.R(0, 0, 64, 64))
	for i := range backbuffer.Pix {
		backbuffer.Pix[i] = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	}

	now := time.Now()
	tt := float32(0.0)
	brightness := float32(1.0)
	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape))

		next := time.Now()
		dt := next.Sub(now).Seconds()
		now = next
		if dt > 1.0/15.0 {
			dt = 1.0 / 15.0
		}
		tt += float32(dt)

		winsize := win.Bounds().Size()
		if canvas.SetSize(int(winsize.X), int(winsize.Y)) {
			backbuffer = pixel.MakePictureData(
				pixel.R(0, 0, float64(canvas.Width), float64(canvas.Height)),
			)
		}
		canvas.Clear()

		{
			os := Sin(float32(tt))

			size := V{float32(winsize.X) * 0.5, float32(winsize.Y) * 0.5, 0}
			for i := 0; i < Detail; i++ {
				p := V{r.Float32() - 0.5, r.Float32() - 0.5, r.Float32() - 0.5}
				p = N(p, 17)
				h := Atan2(p.Y, p.X)
				p = T0(tt, os, p)
				p = Rx(Ry(p, tt), os)
				canvas.Put(P(p, size), HL{h, 0.1})
			}
		}
		canvas.Smear()

		if backbuffer.Stride != canvas.Width {
			panic("invalid stride")
		}

		max, total, nz := float32(0.0), float32(0.0), 0
		for _, v := range canvas.Pix {
			if v.L == 0 {
				continue
			}
			nz++
			total += v.L
			if v.L > max {
				max = v.L
			}
		}
		average := total / float32(nz)
		brightness = Lerp(brightness, Lerp(average, max, 0.1), 0.01)
		scale := 1 / brightness

		for i, v := range canvas.Pix {
			backbuffer.Pix[i] = v.RGBA(scale)
		}

		sprite := pixel.NewSprite(backbuffer, backbuffer.Bounds())
		sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		win.Update()
	}
}

func F0(tt float32, v V) float32 { return Sin(tt+v.X+v.Z) * 3 }
func F1(tt float32, v V) float32 { return Sin(tt+v.X/3+v.Y/3) + 1.5 }
func T0(tt, os float32, v V) V {
	return S(v, F1(tt, v)+os*F0(tt, v)/10+os)
}
