package main

import "fmt"

// A type of struct is a collection of fields.
type Vertex struct {
	Lat, Long float64
}

// "map" is a key, value data type.
// String key and Vertex value
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	// Input float64 constructor into Vertex struct type
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
