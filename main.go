package main

import "github.com/mohamed25-dev/go-dsa/linkedlist"

func main() {
	ll := linkedlist.LinkedList{}

	for i := 1; i <= 10; i++ {
		ll.InsertAtEnd(i)
	}

	ll.Print()

	linkedlist.ReverseLinkedList(&ll)

	ll.Print()
}
