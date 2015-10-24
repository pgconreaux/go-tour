package main

import (
	"code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	var ys [][]uint8
	for y := 0; y < dy; y++ {
		var xs []uint8
		for x := 0; x < dx; x++ {
			xs = append(xs, uint8(x*y))
		}
		ys = append(ys, xs)
	}
	return ys
}

func main() {
	pic.Show(Pic)
}