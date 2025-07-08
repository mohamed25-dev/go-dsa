package array

import "testing"

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		val      int
		expected int
	}{
		{
			name:     "first test",
			arr:      []int{1, 2, 2, 3, 4},
			val:      2,
			expected: 3,
		},
		{
			name:     "second test",
			arr:      []int{1, 4, 2, 3, 4, 5, 9, 6, 7},
			val:      4,
			expected: 7,
		},
	}

	for _, test := range tests {
		if got := removeElement(test.arr, test.val); got != test.expected {
			t.Fatalf("%s failed; expected %d, got %d", test.name, test.expected, got)
		}
	}

}
