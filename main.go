package main

import "github.com/mohamed25-dev/go-dsa/linkedlist"

func main() {
	ll := linkedlist.LinkedList{}

	for i := 1; i <= 10; i++ {
		ll.InsertAtEnd(i)
	}

	ll.InsertAtHead(0)
	ll.InsertAtHead(-1)

	ll.InsertAt(5, 99)
	ll.InsertAt(0, 88)

	ll.DeleteAt(0)
	ll.DeleteAt(1)

	ll.Print()
}
