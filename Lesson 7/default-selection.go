package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Microsecond)  // Tick: Periodic Timer
	boom := time.After(500 * time.Microsecond) // After: One-Time Timer
	for {
		select {
		case <-tick:
			fmt.Println("tick. ")
		case <-boom:
			fmt.Println("BOOM! ")
			return
		default:
			// The default case in a select is run if no other case is ready.
			{
				fmt.Println("    .")
				time.Sleep(50 * time.Microsecond)
			}
		}
	}
}
