package main

import (
	"fmt"
	"math"
)

// Declare a "non-struct" type
type MyFloat float64

// Define "Abs()" methods on MyFloat types.
// A method is a function with a special "receiver argument".
// "receiver argument" also support normal function. It's similar with "constructor".
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
func main() {
	f := MyFloat(-math.Sqrt2)  // "math.Sqrt2" = constant of "√2"
	fmt.Println(f.Abs())
}
