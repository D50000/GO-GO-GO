package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Replace the default io.Reader function:
//
//	type Reader interface {
//	    Read(p []byte) (n int, err error)
//	}
func (r MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A' // Replace the character to 'A'.
	}

	return len(p), nil // error nil for run successfully.
}

func main() {
	// Golang official library check tool.
	reader.Validate(MyReader{})
}
