package main

import "fmt"

type Vertex struct {
	// X, Y call fields
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// Struct fields are accessed using a dot.
	v.X = 4
	fmt.Println(v.X)
}
