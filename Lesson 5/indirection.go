package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5)
// since the Scale method has a pointer receiver.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) // Whatever use "&v" or not it's OK
	ScaleFunc(&v, 10) // no "&v" will compile error.

	// For the statement v.Scale(5), even though v is a value and not a pointer,
	// the method with the pointer receiver is called "automatically".

	p := &Vertex{4, 3}
	p.Scale(3) // OK
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
