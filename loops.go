package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	y := 0.0
	i := 0
	for math.Abs(y-z) > 0.000000001 {
		y = z
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
	    i++
	}
	fmt.Println(i)
	return z
}

func main() {
	v := 49.0
	fmt.Println(Sqrt(v))
	fmt.Println(math.Sqrt(v))
}
