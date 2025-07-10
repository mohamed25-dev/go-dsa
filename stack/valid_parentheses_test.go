package stack

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected bool
	}{
		{
			name:     "success test",
			str:      "(){{}}",
			expected: true,
		},
		{
			name:     "failed test",
			str:      "()){{}}",
			expected: false,
		},
	}

	for _, test := range tests {
		if got := IsValid(test.str); got != test.expected {
			t.Fatalf("%s failed; expected %v, got %v", test.name, test.expected, got)
		}
	}

}
