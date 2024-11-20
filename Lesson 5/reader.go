package main

import (
	"fmt"
	"io"
	"strings"
)

// The io.Reader interface, which represents the read end of a stream of data.
// The Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.
func main() {
	r := strings.NewReader("Hello Reader!") // Similar with files class.

	b := make([]byte, 8) // Initial 8 bytes data with [0 0 0 0 0 0 0 0].
	for { // Infinity for loop until break.
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b =%v\n", n, err, b) // n = 8 err = <nil> b =[72 101 108 108 111 32 82 101] b[:n] = "Hello Re"
		fmt.Printf("b[:n] = %q\n", b[:n]) // n = 5 err = <nil> b =[97 100 101 114 33 32 82 101] b[:n] = "ader!"
		if err == io.EOF { // n = 0 err = EOF b =[97 100 101 114 33 32 82 101] b[:n] = ""
			break
		}
	}
}
