package queue

import (
	"testing"
)

func TestEnqueueAndDequeue(t *testing.T) {
	q := &Queue{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	if val, ok := q.Dequeue(); !ok || val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	if val, ok := q.Dequeue(); !ok || val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}

	if val, ok := q.Dequeue(); !ok || val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}

	if _, ok := q.Dequeue(); ok {
		t.Errorf("Expected dequeue on empty queue to fail")
	}
}

func TestPeek(t *testing.T) {
	q := &Queue{}

	if _, ok := q.Peek(); ok {
		t.Errorf("Expected Peek on empty queue to fail")
	}

	q.Enqueue(42)

	if val, ok := q.Peek(); !ok || val != 42 {
		t.Errorf("Expected Peek to return 42, got %d", val)
	}

	if size := q.Size(); size != 1 {
		t.Errorf("Expected size to still be 1, got %d", size)
	}
}
