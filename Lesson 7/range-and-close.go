package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// A sender can close a channel to indicate that no more values will be sent.
	// Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	// Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c) // cap(c) return the capacity of the channel.
	for i := range c {      // "i" is "false" if "c" are no more values to receive and the channel is closed.
		fmt.Println(i)
	}
}
