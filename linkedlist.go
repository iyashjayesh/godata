package main

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Append(value int) {
	node := &Node{Value: value}
	if l.Tail == nil {
		l.Head = node
		l.Tail = node
		return
	}
	l.Tail.Next = node
	l.Tail = node
}

func (l *LinkedList) Prepend(value int) {
	node := &Node{Value: value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	node.Next = l.Head
	l.Head = node
}

func (l *LinkedList) Remove(value int) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (l *LinkedList) Contains(value int) bool {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (l *LinkedList) Length() int {
	length := 0
	current := l.Head
	for current != nil {
		length++
		current = current.Next
	}
	return length
}
