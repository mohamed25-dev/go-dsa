package stack

import (
	"slices"
)

func IsValid(s string) bool {
	openParenthesis := []rune{'(', '{', '['}
	closeParenthesis := []rune{')', '}', ']'}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := Stack{}

	for _, char := range s {
		if ok := slices.Contains(openParenthesis, rune(char)); ok {
			stack.Push(rune(char))
		}

		if ok := slices.Contains(closeParenthesis, rune(char)); ok {
			val, ok := stack.Pop()
			if !ok {
				return false
			}

			if val != pairs[rune(char)] {
				return false
			}
		}
	}

	return stack.IsEmpty()
}
