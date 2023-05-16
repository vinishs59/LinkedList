package main

import (
	"fmt"

	List "github.com/vinishs59/LinkedList/SinglyLinkedList"
)

func main() {
	// Create a new singly linked list with the values 1, 2, and 3.
	list := List.New(1, 2, 3, 4, 5)

	list.Delete(2)

	// Print the contents of the list.
	for curr := list.Head; curr != nil; curr = curr.Next {
		fmt.Println(curr.Value)
	}

	node := list.Reverse()

	for curr := node; curr != nil; curr = curr.Next {
		fmt.Println(curr.Value)
	}
}
