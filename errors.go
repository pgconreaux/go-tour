package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	s := fmt.Sprint(float64(e))
	return fmt.Sprint("cannot Sqrt negative number: ", s)
}

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	y := 0.0
	i := 0
	for math.Abs(y-z) > 0.000000001 {
		y = z
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
		i++
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
