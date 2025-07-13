package linkedlist

func MergeTwoLists(ll1, ll2 *LinkedList) *LinkedList {

	if ll1.Head == nil {
		return ll2
	}

	if ll2.Head == nil {
		return ll1
	}

	merged := &LinkedList{}

	current1 := ll1.Head
	current2 := ll2.Head

	for current1 != nil && current2 != nil {
		if current1.Val < current2.Val {
			merged.InsertAtEnd(current1.Val)
			current1 = current1.Next
		} else {
			merged.InsertAtEnd(current2.Val)
			current2 = current2.Next
		}
	}

	for current1 != nil {
		merged.InsertAtEnd(current1.Val)
		current1 = current1.Next
	}

	for current2 != nil {
		merged.InsertAtEnd(current2.Val)
		current2 = current2.Next
	}

	return merged
}
