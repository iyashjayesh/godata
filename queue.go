package main

import "fmt"

type Queue struct {
	Head *Node
	Tail *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value int) {
	node := &Node{Value: value}
	if q.Tail == nil {
		q.Head = node
		q.Tail = node
		return
	}
	q.Tail.Next = node
	q.Tail = node
}

func (q *Queue) Dequeue() int {
	if q.Head == nil {
		return 0
	}
	value := q.Head.Value
	q.Head = q.Head.Next
	return value
}

func (q *Queue) Peek() int {
	if q.Head == nil {
		return 0
	}
	return q.Head.Value
}

func (q *Queue) Contains(value int) bool {
	current := q.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (q *Queue) IsEmpty() bool {
	return q.Head == nil
}

func (q *Queue) Clear() {
	q.Head = nil
	q.Tail = nil
}

func (q *Queue) Size() int {
	current := q.Head
	size := 0
	for current != nil {
		size++
		current = current.Next
	}
	return size
}

func (q *Queue) ToSlice() []int {
	slice := make([]int, 0)
	current := q.Head
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}
	return slice
}

func (q *Queue) FromSlice(slice []int) {
	for _, value := range slice {
		q.Enqueue(value)
	}
}

func (q *Queue) Reverse() {
	if q.Head == nil {
		return
	}
	var previous *Node
	current := q.Head
	q.Tail = current
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	q.Head = previous
}

func (q *Queue) Copy() *Queue {
	queue := NewQueue()
	current := q.Head
	for current != nil {
		queue.Enqueue(current.Value)
		current = current.Next
	}
	return queue
}

func (q *Queue) String() string {
	if q.Head == nil {
		return "[]"
	}
	current := q.Head
	str := "["
	for current != nil {
		str += fmt.Sprintf("%v", current.Value)
		if current.Next != nil {
			str += ", "
		}
		current = current.Next
	}
	str += "]"
	return str
}

func (q *Queue) Print() {
	fmt.Println(q)
}

// Path: GoData\datastructures\stack.go