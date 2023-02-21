# GoData

[![Build Status](https://travis-ci.org/your_username/your_library_name.svg?branch=master)](https://travis-ci.org/your_username/your_library_name)

A Golang library that provides implementations of common data structures like linked lists, stacks, and queues.

## Installation

To install, simply run:

`go get github.com/iyashjayesh/godata`


## Usage

```go
import (
    "github.com/iyashjayesh/godata"
)

func main() {
    // Create a new linked list
    ll := godata.NewLinkedList()

    // Add some elements to the linked list
    ll.PushBack(1)
    ll.PushBack(2)
    ll.PushBack(3)

    // Print out the linked list
    for e := ll.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }

    // Create a new stack
    stack := godata.NewStack()

    // Push some elements onto the stack
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    // Pop elements off the stack
    for !stack.IsEmpty() {
        fmt.Println(stack.Pop())
    }

    // Create a new queue
    queue := godata.NewQueue()

    // Enqueue some elements into the queue
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)

    // Dequeue elements from the queue
    for !queue.IsEmpty() {
        fmt.Println(queue.Dequeue())
    }
}
```

### Contributing
Contributions are always welcome! If you'd like to contribute, please open an issue or a pull request on the [GitHub repository](https://github.com/iyashjayesh/GoData)

### License
This library is licensed under the MIT License. See the LICENSE file for more information.