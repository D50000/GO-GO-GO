package main

import (
	"fmt"
	"image"
)

// Package "image" defines the Image interface:
// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100)) // Draw Rectangle => return Rectangle.
	fmt.Println(m.Bounds())  // Show images boarder => return Rectangle.
	fmt.Println(m.At(0, 0).RGBA()) // Show the coordinate(0, 0) RGBA color => return Color.
}
