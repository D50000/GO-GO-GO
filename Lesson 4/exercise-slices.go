package main

// Use `go list -m golang.org/x/tour` to check the package install or not.
import (
	"golang.org/x/tour/pic"
)

// When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
func Pic(dx, dy int) [][]uint8 {
	// code here
	image := make([][]uint8, dy) // Declare 2D slices.
	for y := 0; y < dy; y++ {
        row := make([]uint8, dx) // Declare dx slices, set dx into axis y.
        for x := 0; x < dx; x++ {
            row[x] = uint8(x ^ y) // To calculate the pixel value x ^ y
        }
        image[y] = row
    }
    return image
}

func main() {
	pic.Show(Pic)
}
