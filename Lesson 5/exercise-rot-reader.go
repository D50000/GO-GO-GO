package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader // Execute the Read method.
}

// Implement and replace the build-it default Read function in io.Reader interface.
func (rot *rot13Reader) Read(p []byte) (int, error) {
	// Read data p and out put length n and error.
	n, err := rot.r.Read(p)
	if err != nil {
		return n, err
	}

	// Convert into ROT13.
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}

	return n, nil
}

// Encrypt characters into ROT13.
func rot13(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return (b-'A'+13)%26 + 'A'
	} else if b >= 'a' && b <= 'z' {
		return (b-'a'+13)%26 + 'a'
	}
	return b // No change if not alphabet.
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s} // Initial the rot13Reader struct type. Read and encrypt the content.
	io.Copy(os.Stdout, &r) // "io.Copy" will copy the io.Reader content and output to os.Stdout.
}
