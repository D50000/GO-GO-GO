package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Random seed between 0~10
	fmt.Println("My favorite number is", rand.Intn(10))
}
