package main

import (
	"log"
	"time"

	"honnef.co/go/js/dom/v2"
)

func main() {
	viewElem := dom.GetWindow().Document().QuerySelector("#view")

	view, ok := viewElem.(*dom.HTMLCanvasElement)
	if !ok {
		log.Printf("not a canvas, was %T", viewElem)
		return
	}

	imageElem := dom.GetWindow().Document().QuerySelector("#image")
	image, ok := imageElem.(*dom.HTMLImageElement)
	if !ok {
		log.Printf("not an image, was %T", imageElem)
		return
	}

	context := view.GetContext2d()
	start := time.Now()
	for y := 0; y < 25; y++ {
		for x := 0; x < 25; x++ {
			context.DrawImage(image, float64(x)*16, float64(y)*16)
		}
	}
	delta := time.Since(start)
	context.FillText(delta.String(), 10, 10, 10000000)
}
