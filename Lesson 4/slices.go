package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // Slice's "length is dynamic"
	// The following expression creates a slice which includes elements 1 through 3.
	fmt.Println(s)
}
