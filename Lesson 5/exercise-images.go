package main

import "golang.org/x/tour/pic"

type Image struct{
	// Define your own Image type, implement the necessary methods, and call pic.ShowImage.
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
