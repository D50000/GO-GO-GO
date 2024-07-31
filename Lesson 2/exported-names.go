package main

import (
	"fmt"
	"math"
)

// In Go, a name is exported if it begins with a capital letter (Public). For example, Pizza is an exported name, as is Pi, which is exported from the math package.
// pizza and pi do not start with a capital letter, so they are not exported (Private).
func main() {
	// Print PI（π）
	// To fix the error, rename math.pi to math.Pi and try it again.
	fmt.Println(math.Pi)
}
