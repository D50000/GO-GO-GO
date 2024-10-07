package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Need to use "Indirecting `*p`" to link back to "Vertex{3, 4}".
// Similar with "self.", "this." in OOP.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// With a value receiver, the Scale method operates on a copy of the original Vertex value. 
	// When the v.Scale(10) finish, it will reset back to initial "Vertex{3, 4}".
	v.Scale(10)
	fmt.Println(v.Abs())
}
