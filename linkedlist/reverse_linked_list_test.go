package linkedlist

import (
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty list", []int{}, []int{}},
		{"Single element", []int{1}, []int{1}},
		{"Two elements", []int{1, 2}, []int{2, 1}},
		{"Multiple elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{}
			for _, v := range tt.input {
				ll.InsertAtEnd(v)
			}

			ReverseLinkedList(ll)
			result := ll.ToSlice()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ReverseLinkedList() = %v, expected %v", result, tt.expected)
			}

			if len(tt.expected) > 0 {
				if ll.Head == nil || ll.Head.Val != tt.expected[0] {
					t.Errorf("Expected head to be %d, got %v", tt.expected[0], ll.Head)
				}
				if ll.Tail == nil || ll.Tail.Val != tt.expected[len(tt.expected)-1] {
					t.Errorf("Expected tail to be %d, got %v", tt.expected[len(tt.expected)-1], ll.Tail)
				}
			}
		})
	}
}
