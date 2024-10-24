package main

import "fmt"

func main() {
	var i interface{} = "Hello"

	// Type Assertion
	S := i.(string)
	fmt.Println(s) // Hello

	// This kind of assertion won't compile error when assert fail.
	s, ok := i.(string) // assertion success
	fmt.Println(s, ok) // Hello true

	f, ok := i.(float64) // assertion fail
	fmt.Println(f, ok) // 0 false

	f = i.(float64) // compile error
	fmt.Println(f) // panic: interface conversion: interface {} is string, not float64
}
