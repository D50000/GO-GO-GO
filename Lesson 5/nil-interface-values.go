package main

import "fmt"

type I interface {
	// A nil interface value holds neither value nor concrete type.
	M()
}

func main() {
	var i I
	describe(i)
	i.M() // nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
