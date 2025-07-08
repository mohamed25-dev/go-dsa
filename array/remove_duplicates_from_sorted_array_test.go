package array

import "testing"

func TestRemoveDuplicatesOption1(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "first test",
			arr:      []int{1, 2, 2, 3, 4},
			expected: 4,
		},
		{
			name:     "second test",
			arr:      []int{1, 2, 2, 3, 4, 5, 5, 6, 6, 6, 7},
			expected: 7,
		},
	}

	for _, test := range tests {
		if got := removeDuplicatesOption1(test.arr); got != test.expected {
			t.Fatalf("%s failed; expected %d, got %d", test.name, test.expected, got)
		}
	}

}

func TestRemoveDuplicatesOption2(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "first test",
			arr:      []int{1, 2, 2, 3, 4},
			expected: 4,
		},
		{
			name:     "second test",
			arr:      []int{1, 2, 2, 3, 4, 5, 5, 6, 6, 6, 7},
			expected: 7,
		},
	}

	for _, test := range tests {
		if got := removeDuplicatesOption2(test.arr); got != test.expected {
			t.Fatalf("%s failed; expected %d, got %d", test.name, test.expected, got)
		}
	}

}
