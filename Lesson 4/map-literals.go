package main

import "fmt"

type Vertex struct {
	Lat, long float64
}

// String key and Vertex value with initial value.
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},

	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
