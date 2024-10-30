package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Ubiquitous interfaces is Stringer
// Default interface. A Stringer is a type that can describe itself as a string.
// type Stringer interface {
//     String() string
// }
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
