package main

// To Fix no required module provides package golang.org/x/tour/wc
// `go install` and it will install module at: `$GOPATH/pkg/mod`
import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var words []string = strings.Split(s, " ") // strings.Split()
	fmt.Println(words)

	var result = make(map[string]int)
	// `_` for ignore the index parameter in range method.
	for _, v := range words {
		_, isExists := result[v]
		// Ignore the "value" return in map["key"]
		if isExists { // Shorthand for => if _, isExists := result[v]; isExists { }
			result[v] += 1
		} else {
			result[v] = 1
		}
	}
	return result
}

func main() {
	fmt.Println("Test run:")

	// golang lib for test suit for the method.
	wc.Test(WordCount)
}
