package main

import "fmt"

func main() {
	// The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding function returns.
	defer fmt.Println("world")

	fmt.Println("hello")
}
