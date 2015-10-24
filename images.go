package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"

)

type Image struct{
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}
func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
