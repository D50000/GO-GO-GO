package main
// TODO: Fix no required module provides package golang.org/x/tour/pic
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
}

func main() {
	pic.Show(Pic)
}
