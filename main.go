package main

import (
	"fmt"

	"github.com/iyashjayesh/GoData/datastructures"
)

func main() {

	linklist := datastructures.New()
	linklist.Append(1)
	linklist.Append(2)
	linklist.Append(3)
	linklist.Append(4)
	linklist.Append(5)

	fmt.Println("Linked List:")
	node := linklist.Head
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}

	linklist.Remove(3)
	fmt.Println("Linked List after removing 3:")
}
