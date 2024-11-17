package main

import (
	"fmt"
)

// Error datatype.
type ErrNegativeSqrt float64

// "Implement the Error interface" and "override" the build-in one.
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
	// Use "float64(e)" to prevent infinite error log.
	// Root Cause:
	// e ErrNegativeSqrt implement the Error() method will Sprintf(e) and check e type infinite.
	// float64() is build-in simple type that doesn't implement Error() method.
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) // Return {x, error}
	}
	
	// Calculate the square root method and return nil error.
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
