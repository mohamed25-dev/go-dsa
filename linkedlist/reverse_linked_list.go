package linkedlist

func ReverseLinkedList(ll *LinkedList) {
	if ll.Length <= 1 {
		return
	}

	var prev *Node
	current := ll.Head

	ll.Tail = ll.Head

	// 1 -> 2 -> 3 -> 4
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	ll.Head = prev
}
