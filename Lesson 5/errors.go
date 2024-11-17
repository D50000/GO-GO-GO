package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// Overwrite default "Error() method".
func (e *MyError) Error() string {
	return fmt.Sprintf("Now is %v, %s", e.When, e.What)
}

// The "error type" is a "built-in interface" similar to fmt.Stringer:
//
//	type error interface {
//	    Error() string
//	}
func run() error {
	return &MyError{ // Initial the "MyError type" and return "non-nil error" type.
		time.Now(),
		"It don't work...",
	}
}

func main() {
	// A nil error denotes success; a non-nil error denotes failure.
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
