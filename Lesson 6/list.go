package main

import (
	"fmt"
)

// To make the "list functional", it will need:
// - A method to append elements to the list.
// - A method to traverse and print the elements.
// - Additional helper functions as needed.

// "List" represents a singly-linked list that holds values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// "Append" adds a new value to the end of the list.
func (l *List[T]) Append(value T) {
	current := l
	// Traverse to the end of the list.
	for current.next != nil {
		current = current.next
	}
	// Add a new node with the given values.
	current.next = &List[T]{val: value}
}

// "Traverse prints" all elements in the list.
func (l *List[T]) Traverse() {
	current := l
	// Loop through the list and print values.
	for current != nil {
		fmt.Printf("%v : ", current.val)
		current = current.next
	}
	fmt.Println("nil")
}

// "Length" returns the number of elements in the list.
func (l *List[T]) Length() int {
	count := 0
	current := l
	// Loop through the list and count nodes.
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// "Create" a new linked list with an initial value.
func NewList[T any](value T) *List[T] { // return pointer list.
	return &List[T]{val: value} // Returns a pointer (heap-allocated).
}

func main() {
	// Create a new list.
	list := NewList(10)

	// Append elements.
	list.Append(20)
	list.Append(30)
	list.Append(40)

	// Traverse and print.
	fmt.Println("List elements:")
	list.Traverse()

	// Length of list.
	fmt.Printf("Length of the list: %d\n", list.Length())
}
