package main

type Stack struct {
	items []int
}

func NewStack(items ...int) *Stack {
	return &Stack{items: items}
}

func (s *Stack) Push(item int) {
	s.push(item)
}

func (s *Stack) Pop() int {
	return s.pop()
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() int {
	last := len(s.items) - 1
	s.items = s.items[:last]
	return last
}
