package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Binary Tree Data Type:
// type Tree struct {
// 	Left  *Tree // Left tree
// 	Value int   // Node value
// 	Right *Tree // Right tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch using in-order traversal.
func Walk(t *tree.Tree, ch chan int) {
	// Go through the tree.
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	// Walk both trees in separate goroutines.
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	// Compare the output of both channels.
	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 { // ch2 empty or not same.
			return false
		}
	}

	// Ensure the second channel is also fully consumed.
	_, ok := <-ch2
	return !ok
}

func main() {
	ch := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
		// tree.New(1) means create a tree with "1" for cardinality.
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
