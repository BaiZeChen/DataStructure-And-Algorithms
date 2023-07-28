package main

import (
	"fmt"
)

func NewLoopQueue[T any](capacity int) *LoopQueue[T] {
	loop := &LoopQueue[T]{}
	return loop.init(capacity)
}

type LoopQueue[T any] struct {
	arr         []T
	front, tail int
}

func (l *LoopQueue[T]) init(capacity int) *LoopQueue[T] {
	l.arr = make([]T, capacity+1)
	return l
}

func (l *LoopQueue[T]) IsEmpty() bool {
	return l.front == l.tail
}

func (l *LoopQueue[T]) IsFull() bool {
	return (l.tail+1)%len(l.arr) == l.front
}

func (l *LoopQueue[T]) Size() (sum int) {
	sum = 0
	for i := l.front; i != l.tail; i = (i + 1) % len(l.arr) {
		sum++
	}
	return
}

func (l *LoopQueue[T]) Enqueue(value T) bool {
	if l.IsFull() {
		l.expansion()
	}

	l.arr[l.tail] = value
	l.tail = (l.tail + 1) % len(l.arr)
	return true
}

func (l *LoopQueue[T]) Dequeue() (res T) {
	if l.IsEmpty() {
		return
	}

	res = l.arr[l.front]
	l.front = (l.front + 1) % len(l.arr)
	l.shrinkage()
	return
}

// 扩容按照2倍处理
func (l *LoopQueue[T]) expansion() {
	tmp := make([]T, 2*len(l.arr))
	for i := 0; i < l.Size(); i++ {
		tmp[i] = l.arr[(l.front+i)%len(l.arr)]
	}
	l.front, l.tail, l.arr = 0, l.Size(), tmp
}

func (l *LoopQueue[T]) shrinkage() {
	if len(l.arr) > 2 && l.Size() == len(l.arr)/4 {
		tmp := make([]T, len(l.arr)/2)
		for i := 0; i < l.Size(); i++ {
			tmp[i] = l.arr[(l.front+i)%len(l.arr)]
		}
		l.front, l.tail, l.arr = 0, l.Size(), tmp
	}
}

func (l *LoopQueue[T]) String() string {
	res := fmt.Sprintf("Array: realSize = %d ,capacity = %d\n", len(l.arr), cap(l.arr))
	res += "["
	for i := 0; i < len(l.arr); i++ {
		res += fmt.Sprintf("%v", l.arr[i])
		if i != len(l.arr)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}

func main() {
	loop := NewLoopQueue[int](6)
	for i := 0; i < 6; i++ {
		loop.Enqueue(i + 1)
	}
	fmt.Println(loop.Dequeue())
	fmt.Println(loop.Dequeue())
	fmt.Println(loop.front)
	fmt.Println(loop.tail)
	fmt.Println(loop.Enqueue(8))
	fmt.Println(loop.Enqueue(9))
	fmt.Println(loop.Enqueue(10))
	fmt.Println(loop)
}
