package main

// Use `go list -m golang.org/x/tour` to check the package install or not.
import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// code here
}

func main() {
	pic.Show(Pic)
}
