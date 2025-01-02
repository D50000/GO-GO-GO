package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// It will run the case that channel ready to provide value.
		// It chooses one at random if multiple are ready. If all not ready will run default case.
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return // Shutdown the process.
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
