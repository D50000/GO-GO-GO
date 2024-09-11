package main

import "fmt"

func main() {
	// Declare a slice
	var s []int
	printSlice(s)

	// Built-in append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	// Allocated bigger array strategy (https://go.dev/blog/slices-intro)
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
