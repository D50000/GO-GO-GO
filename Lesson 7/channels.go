package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // Send sum to channel c.
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // Create a int channel by build-in make function.
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // Receive from channel c.
	// Above statement will "block the process" until all the chanel operator is finnish.
	fmt.Println(x, y, x+y)
}
