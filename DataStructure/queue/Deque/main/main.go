package main

import "fmt"

func NewDeque[T any](capacity int) *Deque[T] {
	deque := &Deque[T]{}
	return deque.init(capacity)
}

type Deque[T any] struct {
	arr               []T
	front, tail, size int // 用size来维护，不需要多一个空间
}

func (d *Deque[T]) init(capacity int) *Deque[T] {
	d.arr = make([]T, capacity)
	return d
}

func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque[T]) IsFull() bool {
	return d.size == len(d.arr)
}

func (d *Deque[T]) AddFront(value T) bool {
	if d.IsFull() {
		d.expansion()
	}

	// 为什么这里是先确定下标在赋值跟tail不一样呢？
	// 因为front指的就是当前元素，而tail指的是下一个要添加的元素
	if d.front == 0 {
		d.front = len(d.arr) - 1
	} else {
		d.front--
	}
	d.arr[d.front] = value
	d.size++

	return true
}

func (d *Deque[T]) AddLast(value T) bool {
	if d.IsFull() {
		d.expansion()
	}
	d.arr[d.tail] = value
	d.tail = (d.tail + 1) % len(d.arr)
	d.size++
	return true
}

func (d *Deque[T]) RemoveFront() (res T) {
	if d.IsEmpty() {
		return
	}

	res = d.arr[d.front]
	d.front = (d.front + 1) % len(d.arr)
	d.size--
	d.shrinkage()
	return
}

func (d *Deque[T]) RemoveLast() (res T) {
	if d.IsEmpty() {
		return
	}

	if d.tail == 0 {
		d.tail = len(d.arr) - 1
	} else {
		d.tail--
	}
	res = d.arr[d.tail]
	d.shrinkage()
	return
}

// 扩容按照2倍处理
func (d *Deque[T]) expansion() {
	tmp := make([]T, 2*len(d.arr))
	for i := 0; i < d.size; i++ {
		tmp[i] = d.arr[(d.front+i)%len(d.arr)]
	}
	d.front, d.tail, d.arr = 0, d.size, tmp
}

func (d *Deque[T]) shrinkage() {
	if len(d.arr) > 2 && d.size == len(d.arr)/4 {
		tmp := make([]T, len(d.arr)/2)
		for i := 0; i < d.size; i++ {
			tmp[i] = d.arr[(d.front+i)%len(d.arr)]
		}
		d.front, d.tail, d.arr = 0, d.size, tmp
	}
}

func main() {
	deque := NewDeque[int](3)
	deque.AddFront(1)
	deque.AddFront(2)
	deque.AddLast(3)
	fmt.Println(deque.arr)
	deque.AddFront(4)
	fmt.Println(deque.arr)
}
