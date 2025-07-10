package stack

type Stack struct {
	data string
}

func (s *Stack) Push(char rune) {
	s.data = s.data + string(char)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.data) == 0 {
		return rune(0), false
	}

	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return rune(top), true
}

func (s *Stack) Peak() (rune, bool) {
	if len(s.data) == 0 {
		return rune(0), false
	}

	top := rune(s.data[len(s.data)-1])
	return top, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
