package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Receiver method
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Receiver method
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// There are two reasons to use a "pointer receiver".
// The first is so that the method can "modify the value" that its receiver points to.
// The second is to "avoid copying" the value on each method call. This can be "more efficient" if the receiver is a large struct.
func main() {
	v := &Vertex{3, 4}
	// %v: just print struct value, %+v: will also print the key and value.
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
