package main

import "github.com/mohamed25-dev/go-dsa/linkedlist"

func main() {
	ll := linkedlist.LinkedList{}
	ll2 := linkedlist.LinkedList{}

	for i := 1; i <= 8; i = i + 2 {
		ll.InsertAtEnd(i)
	}

	for i := -1; i <= 13; i = i + 3 {
		ll2.InsertAtEnd(i)
	}

	ll.Print()

	// linkedlist.ReverseLinkedList(&ll)

	ll2.Print()

	merged := linkedlist.MergeTwoLists(&ll, &ll2)
	merged.Print()
}
