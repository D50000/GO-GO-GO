package main

import "fmt"

type IPAddr [4]byte

// Add a "String() string" method to IPAddr and overwrite the original method in "fmt package".
func (ip IPAddr) String() string {
	// Use `strings.Join` to join the individual byte parts with "."
	// Convert each byte to a string representation.
	return fmt.Sprintf("\"%d.%d.%d.%d\"", ip[0], ip[1], ip[2], ip[3])
	// "%d": print decimal integers.
	// "%v": print generic placeholder for any value type.
}

func main(){
	// Hashmap: with string property with IPAddr value.
	hosts := map[string]IPAddr{ 
		// Initial
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}