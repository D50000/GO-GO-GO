package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2) // -1 * "√2"
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser's Abs()
	a = &v // a *Vertex implements Abser's Abs()

	// In the following line, v is a Vertex (not *Vertex) and does NOT implement Abser.
	a = v // Compile error. It can't adapt func (v *Vertex) Abs() for implement.

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
