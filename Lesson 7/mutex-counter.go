package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int // map data type: "keyString": { 12345}
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock() // Lock so only "one goroutine" at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()         // Lock so only "one goroutine" at a time can access the map c.v.
	defer c.mu.Unlock() // Defer the unlock execute after function finish.
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)} // Initial the Struct v property.
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second) // Sleep 1s for make sure c.Inc finish before c.Value.
	fmt.Println(c.Value("somekey"))
}
