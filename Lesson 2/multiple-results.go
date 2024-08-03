package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// Use ":=" for multiple response
	a, b := swap("hello", "word")
	fmt.Println(a, b)
}
