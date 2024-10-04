package main

import "fmt"

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	// code here
	// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
	var a, b int = 0, 1
	// Initials
	return func() int {
		var temp int = a
		a = b
		b = temp + b
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
