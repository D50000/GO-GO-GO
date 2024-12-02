package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Build-it image Package: "image" defines the Image interface:
// type Image interface {
//     Bounds() Rectangle
//     ColorModel() color.Model
//     At(x, y int) color.Color
// }

type Image struct {
	Width  int
	Height int
}

// Implement and replace the Bounds().
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

// Implement and replace the ColorModel().
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Implement and replace the At().
func (img Image) At(x, y int) color.Color {
	v := uint8((x ^ y) % 256)         // Generate the v coefficient of 8-bit unsigned integer.
	return color.RGBA{v, v, 255, 255} // Return color.
}

func main() {
	// Initial the Image instance.
	m := Image{Width: 256, Height: 256}

	pic.ShowImage(m)
}
