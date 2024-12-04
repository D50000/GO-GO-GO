package main

import "fmt"

// Multiple types using "type parameters".
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints.
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings.
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}