package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // 1000000000
	// Golang struct pointer auto dereference the field.
	// So it equally mean (*p).X like C & C++
	fmt.Println(v)
}
