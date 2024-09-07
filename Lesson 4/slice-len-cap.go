package main

import "fmt"

// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// Slice the slice to give it zero length.
	s = s[:0] // If the "capacity" doesn't delete it still can slice back the data.
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:] // When if remove element, it can slice back some data
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
