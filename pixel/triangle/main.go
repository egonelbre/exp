package main

import (
	"image"
	"image/color"
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "Draw a Triangle",
		Bounds:    pixel.R(0, 0, 480, 520),
		Resizable: false,
		VSync:     true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// create a test image
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	x, y := 0, 0
	for i := 0; i < len(m.Pix); i += 4 {
		if x >= 100 {
			y++
			x = 0
		}
		x++

		m.Pix[i] = byte(y + x + (i * 1 % 32))
		m.Pix[i+1] = byte(y + (i * 2 % 32))
		m.Pix[i+2] = byte(y + (i * 3 % 32))
		m.Pix[i+3] = 0xFF
	}

	pd := pixel.PictureDataFromImage(m)

	start := time.Now()
	for !win.Closed() {
		t := time.Since(start).Seconds()

		win.SetClosed(win.JustPressed(pixelgl.KeyEscape))
		win.Clear(color.RGBA{0x80, 0x80, 0x80, 0xFF})

		sprite := pixel.NewSprite(pd, pd.Bounds())
		sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		s := math.Sin(t)

		tridata := pixel.MakeTrianglesData(3)
		tris := *tridata
		tris[0].Picture = pixel.Vec{0.5, 0}.Scaled(100)
		tris[1].Picture = pixel.Vec{0, s}.Scaled(100)
		tris[2].Picture = pixel.Vec{1, 1}.Scaled(100)

		tris[0].Intensity = 1
		tris[1].Intensity = 1
		tris[2].Intensity = 1

		tris[0].Position = pixel.Vec{0.5, 0}.Scaled(100)
		tris[1].Position = pixel.Vec{0, 1}.Scaled(100)
		tris[2].Position = pixel.Vec{1, 1}.Scaled(100)

		// using pixel.Drawer
		drawer := pixel.Drawer{}
		drawer.Triangles = tridata
		drawer.Picture = pd
		drawer.Draw(win)

		// using pixel.Batch
		// batch := pixel.NewBatch(tridata, pd)
		// batch.Draw(win)

		win.Update()
	}
}
