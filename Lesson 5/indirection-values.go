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

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v)) // Compile OK!
	// fmt.Println(AbsFunc(&v)) // Compile error!

	p := &Vertex{4, 3}
	fmt.Println(p.Abs()) // In this case, the method call p.Abs() is interpreted as (*p).Abs(). The pointer receiver is called "automatically".
	fmt.Println(AbsFunc(*p))
}
