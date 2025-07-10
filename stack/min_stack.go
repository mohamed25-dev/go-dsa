package stack

type MinStack struct {
	data []int
	min  []int
}

func (s *MinStack) Push(num int) {
	min, ok := s.GetMin()
	if !ok || min > num {
		s.min = append(s.min, num)
	} else {
		s.min = append(s.min, min)
	}

	s.data = append(s.data, num)
}

func (s *MinStack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	s.min = s.min[:len(s.min)-1]

	return top, true
}

func (s *MinStack) Peak() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	return s.data[len(s.data)-1], true
}

func (s *MinStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *MinStack) GetMin() (int, bool) {
	if len(s.min) == 0 {
		return 0, false
	}
	return s.min[len(s.min)-1], true
}
