package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x + 20), uint8(y + 2), uint8(x + 1), uint8(y)}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
