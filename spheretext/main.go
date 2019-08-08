package main

import (
	"flag"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"math"
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

	const size = 256
	const frames = 5
	const delay = 50

	sphere := &Sphere{
		Size:      image.Point{size, size},
		Reference: dst,
		Radius:    size / 2,
	}
	animation := &gif.GIF{}
	for i := 0; i < frames; i++ {
		frame := sphere.Frame(float64(i) * delay * 0.001)
		animation.Image = append(animation.Image, frame)
		animation.Delay = append(animation.Delay, delay)
	}

	file, err := os.Create("sphere.gif")
	check(err)
	defer file.Close()

	check(gif.EncodeAll(file, animation))
}

type Sphere struct {
	Size image.Point

	Reference *image.RGBA
	Radius    float64
}

const Tau = math.Pi

func (sphere *Sphere) Frame(time float64) *image.Paletted {
	r := image.Rect(0, 0, sphere.Size.X, sphere.Size.Y)
	img := image.NewPaletted(r, palette.Plan9)

	for y := -sphere.Radius; y <= sphere.Radius; y++ {
		v := (y + sphere.Radius) / 2 * sphere.Radius
		for rad := 0.0; rad <= Tau; rad += 1 / Tau {
			u := rad / Tau
			rgba := sphere.Reference.RGBAAt(int64(u*10), int64(v*sphere.Reference.Size().Y))

			tx := sphere.Radius * math.Sin(rad)
			tz := sphere.Radius * math.Cos(rad)
			sy := sphere.Size.X/2 + int64(y)
		}
	}

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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
