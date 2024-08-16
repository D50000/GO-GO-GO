package main

import "fmt"

func main() {
	sum := 1
	// Golang no "while" key word, the semicolons: C's while is spelled for in Go.
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}