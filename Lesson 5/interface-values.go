package main

import (
	"fmt"
	"math"
)

// Golang is implicit interface.
// No need to declare the type with implementation. 
type I interface {
	M()
}

// No need to declare implementation of I. 
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
	// Just "copy f type" to implement M() to print, after this it will "restore back to original value".
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"HELLO"} // Construct and initial the struct with "Dereferencing &p"
	describe(i)
	i.M() // T struct implement the M()

	i = F(math.Pi) // Initial the type with float.
	// Because no using "Dereferencing &p" just "copy i" variable.
	// Can't update it's value.
	describe(i)
	i.M() // F type implement the M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
	// %v  -> print value
	// %+v -> print key and value
	// %T  -> print value type
}
