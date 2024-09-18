package main

import "fmt"

// Define struct
type Vertex struct {
	Lat, Long float64
}

// Initial
var m = map[string]Vertex{
	// Omit the struct name
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}