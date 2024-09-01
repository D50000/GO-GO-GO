package main

import "fmt"

// An array's length is part of its type, so arrays cannot be resized.
// This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Word"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
