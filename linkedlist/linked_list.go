package linkedlist

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (ll *LinkedList) InsertAtEnd(val int) {
	node := &Node{
		Val:  val,
		Next: nil,
	}

	ll.Length++

	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}

	ll.Tail.Next = node
	ll.Tail = node
}

func (ll *LinkedList) InsertAtHead(val int) {
	node := &Node{
		Val:  val,
		Next: nil,
	}

	ll.Length++

	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node

		return
	}

	node.Next = ll.Head
	ll.Head = node
}

func (ll *LinkedList) InsertAt(index, val int) bool {
	if index > ll.Length {
		return false
	}

	node := &Node{
		Val:  val,
		Next: nil,
	}

	if index == 0 {
		ll.InsertAtHead(val)
		return true
	}

	if index == ll.Length {
		ll.InsertAtEnd(val)
		return true
	}

	current := ll.Head
	previous := ll.Head
	for i := 0; i < index; i++ {
		previous = current
		current = current.Next
	}

	previous.Next = node
	node.Next = current

	ll.Length++
	return true
}

func (ll *LinkedList) DeleteAt(index int) bool {
	if index >= ll.Length {
		return false
	}

	if ll.Length == 0 {
		return true
	}

	ll.Length--

	if index == 0 {
		ll.Head = ll.Head.Next
		return true
	}

	current := ll.Head
	previous := ll.Head
	for i := 0; i < index; i++ {
		previous = current
		current = current.Next
	}

	previous.Next = current.Next

	return true

}

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) ToSlice() []int {
	result := []int{}

	if ll.Length == 0 {
		return result
	}

	current := ll.Head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}
