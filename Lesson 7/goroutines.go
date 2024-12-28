package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func main() {
	// A goroutine is a lightweight thread managed by the Go runtime.
	// Goroutines run in the same address space, so access to shared memory must be synchronized.
	go say("world") // Concurrent, erupt simultaneously.
	say("hello")
}
