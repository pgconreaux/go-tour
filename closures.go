package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	cnt := 0
	a := 0
	b := 0
	sum := 0
	return func() int {
		if cnt == 0 {
			sum = 0
		}
		if cnt == 1 {
			sum = 1
		}
		if cnt > 1 {
			a = b
			b = sum
			sum = a + b
		}
		cnt += 1
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
