package adt

type Stack struct {
	size  int
	items [6]string
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Bury(item string) {
	s.items[s.size] = item
	s.size++
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Unbury() string {
	s.size--
	return s.items[s.size]
}

func NewStack() *Stack {
	return &Stack{0, [6]string{}}
}
