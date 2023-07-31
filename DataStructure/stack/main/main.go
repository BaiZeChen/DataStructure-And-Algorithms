package main

import "fmt"

func NewStack[T any](size int) *Stack[T] {
	stack := &Stack[T]{}
	stack.init(size)
	return stack
}

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) init(size int) {
	s.arr = make([]T, 0, size)
}

func (s *Stack[T]) Push(value T) {
	s.arr = append(s.arr, value)
}

func (s *Stack[T]) Pop() (res T) {
	if s.getSize() == 0 {
		return
	}
	length := len(s.arr) - 1
	res = s.arr[length]
	s.arr = s.arr[:length]
	return
}

func (s *Stack[T]) Peek() (res T) {
	if s.getSize() == 0 {
		return
	}
	res = s.arr[len(s.arr)-1]
	return
}

func (s *Stack[T]) IsEmpty() bool {
	return s.getSize() == 0
}

func (s *Stack[T]) getSize() int {
	return len(s.arr)
}

func main() {
	//fmt.Println(test("(){]{}"))
	queue := NewQueue[int](5)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(7)
	queue.Enqueue(9)
	queue.Enqueue(18)
	fmt.Println(queue.Enqueue(12))
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Enqueue(12))
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Enqueue(13))
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

}

// leetcode 20题
func test(s string) bool {
	stack := NewStack[string](len(s))
	for _, value := range s {
		switch string(value) {
		case "{", "(", "[":
			stack.Push(string(value))
		case ")":
			str := stack.Pop()
			if str != "(" {
				return false
			}
		case "}":
			str := stack.Pop()
			if str != "{" {
				return false
			}
		case "]":
			str := stack.Pop()
			if str != "[" {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func NewQueue[T any](capacity int) *Queue[T] {
	loop := &Queue[T]{}
	return loop.init(capacity)
}

type Queue[T any] struct {
	stack1         *Stack[T]
	stack2         *Stack[T]
	size, capacity int
}

func (q *Queue[T]) init(capacity int) *Queue[T] {
	q.stack1 = NewStack[T](capacity)
	q.stack2 = NewStack[T](capacity)
	q.capacity = capacity
	return q
}

func (q *Queue[T]) IsEmpty() bool {
	return q.stack1.IsEmpty() && q.stack2.IsEmpty()
}

func (q *Queue[T]) IsFull() bool {
	return q.size == q.capacity
}

func (q *Queue[T]) Enqueue(value T) bool {
	if q.IsFull() {
		return false
	}
	q.stack1.Push(value)
	q.size++
	return true
}

func (q *Queue[T]) Dequeue() (res T) {
	if q.IsEmpty() {
		return
	}
	q.size--
	if !q.stack2.IsEmpty() {
		return q.stack2.Pop()
	}
	for q.stack1.getSize() > 1 {
		q.stack2.Push(q.stack1.Pop())
	}
	res = q.stack1.Pop()
	return
}
