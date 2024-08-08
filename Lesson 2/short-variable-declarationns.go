package main

import "fmt"

// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

	fmt.Println(i, j, k, c, python, java)
}
