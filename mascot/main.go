package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	Scale         = 10
	TaskbarHeight = 40
)

type Vector struct {
	X, Y float64
}

type Gopher struct {
	Sheet *ebiten.Image

	Pos Vector
	Vel Vector

	Jumping bool
	Right   bool

	Frame float64
}

func NewGopher() *Gopher {
	png, _, err := image.Decode(bytes.NewReader(SpriteSheetPNG[:]))
	if err != nil {
		panic(err)
	}

	m, err := ebiten.NewImageFromImage(png, ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

	g := &Gopher{Sheet: m}
	g.Pos.X += 400

	return g
}

func (g *Gopher) FrameRect(i int) image.Rectangle {
	return image.Rectangle{
		Min: image.Pt(SpriteWidth*i, 0),
		Max: image.Pt(SpriteWidth*i+SpriteWidth, SpriteHeight),
	}
}

func (g *Gopher) Update(screen *ebiten.Image) error {
	screenWidth, screenHeight := ebiten.ScreenSizeInFullscreen()

	dt := 1.0 / ebiten.FPS

	var dir Vector
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !g.Jumping {
			g.Vel.Y = 80
		}
		g.Jumping = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dir.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dir.X += 1
	}

	g.Vel.X += dir.X * dt * 400
	g.Vel.X *= math.Pow(0.01, dt)

	g.Vel.Y -= dt * 120
	if g.Pos.Y <= 0 && g.Vel.Y <= 0 {
		g.Vel.Y = 0
		g.Pos.Y = 0
		g.Jumping = false
	}

	g.Pos.X += g.Vel.X * dt
	g.Pos.Y += g.Vel.Y * dt

	if g.Pos.X < 0 {
		g.Pos.X = 0
	} else if g.Pos.X > float64(screenWidth/Scale-SpriteWidth) {
		g.Pos.X = float64(screenWidth/Scale - SpriteWidth)
	}

	if g.Vel.X > 0.01 {
		g.Right = true
	} else if g.Vel.X < -0.01 {
		g.Right = false
	}

	ebiten.SetWindowPosition(
		int(g.Pos.X*Scale),
		-int(g.Pos.Y*Scale)-SpriteHeight*Scale+screenHeight-TaskbarHeight,
	)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if g.Pos.Y > 0 {
		if math.Abs(g.Vel.X) < 10 || math.Abs(g.Vel.Y) < 10 {
			g.DrawFrame(screen, SpriteFrames_Jump_Float)
		} else if g.Vel.Y < 0 {
			g.DrawFrame(screen, SpriteFrames_Jump_Fall)
		} else {
			g.DrawFrame(screen, SpriteFrames_Jump_Rise)
		}
	} else {
		if math.Abs(g.Vel.X) < 1 {
			g.Frame += 8 * dt
			g.DrawFrame(screen, SpriteFrames_Front)
		} else {
			if math.Abs(g.Vel.X) < 52 {
				g.DrawFrame(screen, SpriteFrames_Walk)
			} else {
				g.DrawFrame(screen, SpriteFrames_Run)
			}
		}
	}

	g.Frame += math.Abs(g.Vel.X*dt) * 0.2

	return nil
}

func (g *Gopher) DrawFrame(screen *ebiten.Image, frames []int) {
	op := &ebiten.DrawImageOptions{}
	if g.Right {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(SpriteWidth, 0)
	}
	op.GeoM.Scale(Scale, Scale)

	frameIndex := int(g.Frame) % len(frames)
	frameRect := g.FrameRect(frames[frameIndex])
	sprite := g.Sheet.SubImage(frameRect).(*ebiten.Image)
	screen.DrawImage(sprite, op)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetScreenTransparent(true)
	ebiten.SetWindowDecorated(false)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowFloating(true)

	gopher := NewGopher()
	if err := ebiten.Run(gopher.Update, SpriteWidth*Scale, SpriteHeight*Scale, 1, "Running Gopher"); err != nil {
		log.Fatal(err)
	}
}
