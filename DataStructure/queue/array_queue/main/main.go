package main

import (
	"DataStructure-And-Algorithms/DataStructure/trends_array"
	"fmt"
)

func NewQueue(capacity int) *Queue[any] {
	var queue = &Queue[any]{}
	return queue.init(capacity)
}

type Queue[T any] struct {
	arr *trends_array.Slice[any]
}

func (q *Queue[T]) init(capacity int) *Queue[T] {
	return &Queue[T]{
		arr: trends_array.NewSlice(capacity),
	}
}

func (q *Queue[T]) Enqueue(value T) {
	q.arr.Append(value)
}

func (q *Queue[T]) Dequeue() (res T) {
	if q.IsEmpty() {
		return
	}
	front, ok := q.arr.Get(0).(T)
	if !ok {
		return
	}
	q.arr.RemoveFirst()
	res = front
	return
}

func (q *Queue[T]) getFront() (res T) {
	if q.IsEmpty() {
		return
	}
	front, ok := q.arr.Get(0).(T)
	if !ok {
		return
	}
	res = front
	return
}

func (q *Queue[T]) getSize() int {
	return q.arr.Len()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.getSize() == 0
}

func main() {
	queue := NewQueue(10)
	queue.Enqueue(13)
	queue.Enqueue(1.7)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
