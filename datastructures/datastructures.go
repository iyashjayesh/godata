package datastructures

type DataStructure interface {
}

type Node struct {
	Value int
	Next  *Node
}
