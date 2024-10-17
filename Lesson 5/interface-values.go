package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

// Indirecting `*p`, pointer receiver.
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// Golang simple type also can declare method for it. Concept similar with "Class".
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"HELLO"} // Construct and initial the struct with "Dereferencing &p"
	describe(i)
	i.M() // T struct implement the M()

	i = F(math.Pi) // Initial the type with float.
	describe(i)
	i.M() // F type implement the M()
}

func describe(i I) {
	fmt.Printf("(%V, %T)\n", i, i)
}
