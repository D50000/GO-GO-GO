package main

import (
	"fmt"
)

/*
 *	This general approach is called "Newton's method".
 *	It works well for many functions but especially well for square root.
 **/
func sqrt(x float64) float64 {
	// Initial guess z
	z := x / 2

	// "Newton's method"
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Println(sqrt(2))
}
