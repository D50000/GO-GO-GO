package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Regular function style for manipulating the struct properties by "Dereferencing `&p`" and "Indirecting `*p`".
// For updating the struct properties.
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Here we see the Abs and Scale methods rewritten as functions.
func main() {
	v := Vertex{3, 4}
	Scale(&v, 10) // With no "Pointer" it will just copy of struct. When invoke "v" again will refresh back to initial.
	fmt.Println(Abs(v))
}
