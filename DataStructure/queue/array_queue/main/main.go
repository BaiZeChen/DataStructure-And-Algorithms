package main

import (
	"DataStructure-And-Algorithms/DataStructure/trends_array"
	"fmt"
)

func NewQueue[T any](capacity int) *Queue[T] {
	var queue = &Queue[T]{}
	return queue.init(capacity)
}

type Queue[T any] struct {
	arr *trends_array.Slice[T]
}

func (q *Queue[T]) init(capacity int) *Queue[T] {
	return &Queue[T]{
		arr: trends_array.NewSlice[T](capacity),
	}
}

func (q *Queue[T]) Enqueue(value T) {
	q.arr.Append(value)
}

func (q *Queue[T]) Dequeue() (res T) {
	if q.IsEmpty() {
		return
	}
	res = q.arr.Get(0)
	q.arr.RemoveFirst()
	return
}

func (q *Queue[T]) getFront() (res T) {
	if q.IsEmpty() {
		return
	}
	res = q.arr.Get(0)
	return
}

func (q *Queue[T]) getSize() int {
	return q.arr.Len()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.getSize() == 0
}

func main() {
	queue := NewQueue[int](10)
	queue.Enqueue(13)
	queue.Enqueue(17)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
