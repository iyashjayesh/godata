package main

type Stack struct {
	Head *Node
	Tail *Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value int) {
	node := &Node{Value: value}
	if s.Tail == nil {
		s.Head = node
		s.Tail = node
		return
	}
	s.Tail.Next = node
	s.Tail = node
}

func (s *Stack) Pop() int {
	if s.Head == nil {
		return 0
	}
	if s.Head.Next == nil {
		value := s.Head.Value
		s.Head = nil
		s.Tail = nil
		return value
	}
	current := s.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	value := current.Next.Value
	current.Next = nil
	s.Tail = current
	return value
}

func (s *Stack) Peek() int {
	if s.Tail == nil {
		return 0
	}
	return s.Tail.Value
}

func (s *Stack) Contains(value int) bool {
	current := s.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (s *Stack) IsEmpty() bool {
	return s.Head == nil
}

func (s *Stack) Clear() {
	s.Head = nil
	s.Tail = nil
}
