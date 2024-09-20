package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	
	// Insert or update an element in map.
	m["Answer"] = 48
	// Retrieve an element.
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	// Delete an element.
	fmt.Println("The value:", m["Answer"])

	// Test that a key is present with a two-value assignment.
	v, ok := m["Answer"]
	// If key is in m, ok is true. If not, ok is false.
	fmt.Println("The value:", v, "Present?", ok)
}
