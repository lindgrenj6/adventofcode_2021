package util

type Stack struct {
	store []string
}

func NewStack() *Stack {
	return &Stack{
		store: make([]string, 0),
	}
}

func (s *Stack) Push(str string) {
	s.store = append(s.store, str)
}

func (s *Stack) Pop() string {
	if len(s.store) == 0 {
		return ""
	}
	var last string
	s.store, last = s.store[:len(s.store)-1], s.store[len(s.store)-1]

	return last
}

func (s *Stack) Peek() string {
	if len(s.store) == 0 {
		return ""
	}

	return s.store[len(s.store)-1]
}
