package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// Default [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(pow)

	for i := range pow {
		// Shift 1 left i undefined-int
		pow[i] = 1 << uint(i)
	}

	// You can skip the index or value by assigning to _.
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
