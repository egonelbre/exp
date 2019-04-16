package main

import (
	"flag"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"math/rand"
	"os"
	"time"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

func pt(p fixed.Point26_6) image.Point {
	return image.Point{
		X: int(p.X+32) >> 6,
		Y: int(p.Y+32) >> 6,
	}
}

func main() {
	flag.Parse()

	const padding = 10
	const fontSize = 64

	text := "The quick brown fox jumps over the lazy dog."

	fontdata, err := truetype.Parse(goregular.TTF)
	check(err)
	face := truetype.NewFace(fontdata, &truetype.Options{
		Size: fontSize,
	})

	drawer := &font.Drawer{Src: image.Black, Face: face}
	drawer.Dot = fixed.P(0, 0)

	bounds, advance := drawer.BoundString(text)

	drawer.Dot = fixed.P(padding, padding-bounds.Min.Y.Ceil())
	height := (bounds.Max.Y - bounds.Min.Y).Ceil()

	dst := image.NewRGBA(image.Rect(0, 0, advance.Ceil()+2*padding, height+2*padding))
	drawer.Dst = dst

	drawer.DrawString(text)

	out, err := os.Create("out.png")
	check(err)
	defer out.Close()

	err = png.Encode(out, dst)
	check(err)
}

func main2() {
	const size = 64
	const frames = 5
	const delay = 50

	animation := &gif.GIF{}
	for i := 0; i < frames; i++ {
		animation.Image = append(animation.Image, frame(image.Point{size, size}))
		animation.Delay = append(animation.Delay, delay)
	}

	file, err := os.Create("sphere.gif")
	check(err)
	defer file.Close()

	check(gif.EncodeAll(file, animation))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func frame(size image.Point) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, size.X, size.Y), palette.Plan9)
	draw.Draw(
		img,
		img.Rect,
		&image.Uniform{randomColor()},
		image.ZP,
		draw.Src,
	)
	return img
}

func randomColor() color.RGBA {
	rand.Seed(time.Now().UnixNano())
	r, g, b := rand.Intn(255), rand.Intn(255), rand.Intn(255)
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}
