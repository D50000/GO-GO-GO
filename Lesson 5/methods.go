package main

import (
	"fmt"
	"math"
)

// Declare a struct type
type Vertex struct {
	X, Y float64
}

// "Go" does not have "classes". However, you can define methods on types.
// A method is a function with a special "receiver argument".
// The receiver appears in its own argument list between the func keyword and the method name.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // Vertex.Abs() type 
}
