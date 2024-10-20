package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Others languages would trigger a null pointer exception, 
// but in Go it is common to write methods that gracefully handle being called with a nil receiver
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T // Declare t with "*T" will default be "nil".
	i = t // No need to use "&t" because t is type of "*T".
	describe(i)
	i.M()

	i = &T{"hello"} // T struct implement the M().
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%V, %T)\n", i, i)
}
