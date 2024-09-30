package main

import "fmt"

// The adder function returns a closure. Each closure is bound to its own sum variable.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), // pos := adder() same 'sum'
			neg(-2*i), // neg := adder() same 'sum'
		)
	}
}
