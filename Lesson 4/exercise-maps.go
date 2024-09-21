package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var words []string = strings.Split(s, " ")
	fmt.Println(words)

	var result = make(map[string]int)
	for _, v := range words {
		_, isExists := result[v];
		if isExists {
			result[v] += 1
		} else {
			result[v] = 1
		}
	}
	return result
}

func main() {
	fmt.Println("Test run:")
	wc.Test(WordCount)
}
