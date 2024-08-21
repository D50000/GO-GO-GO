package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	fmt.Printf("time.Saturday (numeric): %d\n", time.Saturday) // Outputs: time.Saturday (numeric): 6
	fmt.Println("time.Saturday:", time.Saturday) // Used for simple, unformatted output: Saturday
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}