package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		// The zero value of a slice is nil.
		fmt.Println("nil!")
	}

}
