package main

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func NewQueue[T any]() *LinkListQueue[T] {
	return &LinkListQueue[T]{}
}

type LinkListQueue[T any] struct {
	size             int
	headPtr, tailPtr *Node[T]
}

func (q *LinkListQueue[T]) Enqueue(value T) bool {
	defer func() {
		q.size++
	}()
	node := &Node[T]{
		Value: value,
	}
	if q.size == 0 {
		q.headPtr = node
		q.tailPtr = node
	} else {
		q.tailPtr.Next = node
		q.tailPtr = node
	}
	return true
}

func (q *LinkListQueue[T]) Dequeue() (res T) {
	if q.size == 0 {
		return
	}
	defer func() {
		q.size--
	}()

	node := q.headPtr
	q.headPtr = q.headPtr.Next
	node.Next = nil
	if q.headPtr == nil {
		// 代表一个也没有了，此时tail也得置为nil
		q.tailPtr = nil
	}
	res = node.Value
	return
}

func main() {
	queue := NewQueue[string]()
	queue.Enqueue("1")
	queue.Enqueue("2")
	fmt.Println(queue.tailPtr)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
