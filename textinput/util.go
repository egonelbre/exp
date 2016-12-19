package main

import (
	"image"
	"image/draw"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

func loadImage(filepath string) (*image.RGBA, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	m, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	if rgba, ok := m.(*image.RGBA); ok {
		return rgba, nil
	}

	rgba := image.NewRGBA(m.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), m, image.Point{0, 0}, draw.Src)

	return rgba, nil
}
