package stack

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		arr      []rune
		expected bool
	}{
		{
			name:     "success test",
			arr:      []rune{'(', ')', '{', '{', '}', '}'},
			expected: true,
		},
		{
			name:     "failed test",
			arr:      []rune{'(', ')', ')', '{', '{', '}'},
			expected: false,
		},
	}

	for _, test := range tests {
		if got := IsValid(test.arr); got != test.expected {
			t.Fatalf("%s failed; expected %v, got %v", test.name, test.expected, got)
		}
	}

}
