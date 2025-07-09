package array

import "testing"

func TestGetConcatenation2(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{
			name:     "first test",
			arr:      []int{1, 0, 3, 2},
			expected: []int{1, 0, 3, 2, 1, 0, 3, 2},
		},
		{
			name:     "second test",
			arr:      []int{3, 0, 8, 1, 6},
			expected: []int{3, 0, 8, 1, 6, 3, 0, 8, 1, 6},
		},
	}

	for _, test := range tests {
		got := getConcatenation(test.arr)
		if !areIntSlicesEqual(got, test.expected) {
			t.Fatalf("%s failed; expected %v, got %v", test.name, test.expected, got)
		}
	}

}

func areIntSlicesEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
