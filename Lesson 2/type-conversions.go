package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// Type conversions
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// := operator is used for short variable declaration and initialization.
	// This operator allows you to declare and initialize a variable in a single step "without having to explicitly specify the type" of the variable. 
	a := 42
	b := float64(a)
	c := uint(b)
	fmt.Println(a, b, c)
}
