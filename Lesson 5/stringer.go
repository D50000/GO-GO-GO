package main

// Default interface: "Stringer" is a type that can describe itself as a string.
//
// type Stringer interface {
//     String() string
// }
import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string { // Overwrite the "String()" method in fmt package.
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}