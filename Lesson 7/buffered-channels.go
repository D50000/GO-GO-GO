package main

import "fmt"

func main() {
	// Provide the buffer length as the second argument to make to initialize a buffered channel.
	// Sends to a buffered channel block only when the buffer is full.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
