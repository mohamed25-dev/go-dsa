package stack

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	s := MinStack{}

	if !s.IsEmpty() {
		t.Error("stack should be empty")
	}
	if _, ok := s.GetMin(); ok {
		t.Error("GetMin should fail on empty stack")
	}

	s.Push(3) // 3
	s.Push(5) // 3, 3
	s.Push(2) // 3, 3, 2
	s.Push(1) // 3, 3, 2, 1

	if min, _ := s.GetMin(); min != 1 {
		t.Errorf("expected min to be 1, got %d", min)
	}

	// Pop 1 → expect min to become 2
	s.Pop()
	if min, _ := s.GetMin(); min != 2 {
		t.Errorf("expected min to be 2 after popping 1, got %d", min)
	}

	// Pop 2 → expect min to become 3
	s.Pop()
	if min, _ := s.GetMin(); min != 3 {
		t.Errorf("expected min to be 3 after popping 2, got %d", min)
	}

	// Test Top
	if top, _ := s.Peek(); top != 5 {
		t.Errorf("expected top to be 5, got %d", top)
	}

	// Pop remaining
	s.Pop()
	s.Pop()

	if !s.IsEmpty() {
		t.Error("stack should be empty after popping all")
	}
	if _, ok := s.GetMin(); ok {
		t.Error("GetMin should fail after emptying stack")
	}
	if _, ok := s.Peek(); ok {
		t.Error("Peak should fail on empty stack")
	}
}
