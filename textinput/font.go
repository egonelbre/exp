package main

import (
	"image"
	"path/filepath"

	"github.com/go-gl/gl/v2.1/gl"
)

type Point struct{ X, Y float32 }

func (a Point) Add(b Point) Point {
	return Point{
		a.X + b.X,
		a.Y + b.Y,
	}
}

type GlyphRect struct {
	Advance  float32
	Min, Max Point
}

type Font struct {
	Glyph map[rune]GlyphRect
	RGBA  *image.RGBA
	Width float32
	GL    uint32
}

type FontInfo struct {
	Path   string
	Glyphs string
	Size   Point
	Widths map[float32]string
}

var ArcadeFont = FontInfo{
	Path:   "arcade_43x74.png",
	Glyphs: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789.,;:?!-_~#\"'&()[]|`\\/@°+=*$£€<>%",
	Size:   Point{43, 74},
	Widths: map[float32]string{
		12: "|",
		15: "il;:'",
		16: ".",
		20: "[]",
		21: ",`*",
		25: "j()$<>",
		27: "°",
		30: "frt-\"&=",
		34: "ITYkn01?/+%",
		36: "_",
		39: "ABCDEFGHJKLMNOPQRSUVWXZabcdeghmopqsuvwxyz23456789!#\\@£€",
		41: "~",
	},
}

func (info *FontInfo) Load() *Font {
	m, err := loadImage(filepath.FromSlash(info.Path))
	if err != nil {
		panic(err)
	}

	widths := map[rune]float32{}
	for size, runes := range info.Widths {
		for _, r := range runes {
			widths[r] = size
		}
	}

	font := &Font{}
	font.Glyph = make(map[rune]GlyphRect)
	font.RGBA = m
	font.Width = info.Size.X / info.Size.Y

	sz := m.Bounds().Size()
	u := Point{info.Size.X / float32(sz.X), info.Size.Y / float32(sz.Y)}
	dot := Point{0, 0}
	for _, r := range info.Glyphs {
		advance := widths[r] / float32(info.Size.Y)
		if advance == 0 {
			advance = u.X / u.Y
		}
		font.Glyph[r] = GlyphRect{
			Advance: advance,
			Min:     dot,
			Max:     dot.Add(u),
		}

		dot.X += u.X
		if dot.X+u.X >= 1 {
			dot.X = 0
			dot.Y += u.Y
		}
	}

	font.upload()
	return font
}

func (font *Font) upload() {
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &font.GL)
	gl.BindTexture(gl.TEXTURE_2D, font.GL)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(font.RGBA.Rect.Size().X),
		int32(font.RGBA.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(font.RGBA.Pix))

	gl.Disable(gl.TEXTURE_2D)
}

func (font *Font) DrawText(p Point, textHeight float32, text string) {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.Enable(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, font.GL)
	{
		gl.Color4ub(0xff, 0xff, 0xff, 0xff)
		gl.Begin(gl.QUADS)

		dot := p
		glyphWidth := font.Width * textHeight
		for _, r := range text {
			if r == '\n' {
				dot.X = p.X
				dot.Y += textHeight
				continue
			}

			g, ok := font.Glyph[r]
			if !ok {
				dot.X += glyphWidth
			}

			gl.TexCoord2f(g.Min.X, g.Min.Y)
			gl.Vertex2f(dot.X, dot.Y)

			gl.TexCoord2f(g.Max.X, g.Min.Y)
			gl.Vertex2f(dot.X+glyphWidth, dot.Y)

			gl.TexCoord2f(g.Max.X, g.Max.Y)
			gl.Vertex2f(dot.X+glyphWidth, dot.Y+textHeight)

			gl.TexCoord2f(g.Min.X, g.Max.Y)
			gl.Vertex2f(dot.X, dot.Y+textHeight)

			dot.X += textHeight * g.Advance
		}
		gl.End()
	}
	gl.Disable(gl.TEXTURE_2D)
}
