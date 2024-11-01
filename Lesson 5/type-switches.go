package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) { // "Type Switch": Assertion type with switch case statement.
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2) // %v prints the "value" 
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v)) // %q prints the "strings with double-quoted" 
	default:
		fmt.Printf("I don't know about type %T!\n", v) // %T prints the "type" 
	}
}

func main() {
	do(21)
	do("Hello")
	do(true)
}
