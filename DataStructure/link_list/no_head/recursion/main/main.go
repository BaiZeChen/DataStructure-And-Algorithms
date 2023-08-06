package main

import (
	"fmt"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func NewLinkList[T any]() *LinkList[T] {
	linkList := &LinkList[T]{}
	return linkList
}

type LinkList[T any] struct {
	size    int
	headPtr *Node[T]
}

func (l *LinkList[T]) add(ptr *Node[T], index int, value T) (*Node[T], bool) {
	if index < 0 || index > l.size {
		return nil, false
	} else if index == 0 {
		// 代表这就是插入点了，进行插入
		insertNode := &Node[T]{
			Value: value,
			Next:  ptr,
		}
		l.size++
		return insertNode, true
	} else {
		node, ok := l.add(ptr.Next, index-1, value)
		ptr.Next = node
		return ptr, ok
	}
}
func (l *LinkList[T]) AddFirst(value T) (ok bool) {
	l.headPtr, ok = l.add(l.headPtr, 0, value)
	return ok
}
func (l *LinkList[T]) AddLast(value T) (ok bool) {
	l.headPtr, ok = l.add(l.headPtr, l.size, value)
	return ok
}
func (l *LinkList[T]) AddByIndex(index int, value T) (ok bool) {
	l.headPtr, ok = l.add(l.headPtr, index, value)
	return ok
}

func (l *LinkList[T]) Get(index int, ptr *Node[T]) (res T) {
	if index < 0 || index > l.size {
		return
	} else if index == 0 {
		res = ptr.Value
	} else {
		index--
		res = l.Get(index, ptr.Next)
	}
	return
}
func (l *LinkList[T]) GetFirst() T {
	return l.Get(0, l.headPtr)
}
func (l *LinkList[T]) GetLast() T {
	return l.Get(l.size-1, l.headPtr)
}

func (l *LinkList[T]) del(ptr *Node[T], index int) (*Node[T], bool) {
	if ptr == nil {
		return ptr, true
	}
	// 注意，如果当index等于0时，会出现负值，但不受印象
	node, ok := l.del(ptr.Next, index-1)
	if index == 0 {
		l.size--
		return node, ok
	} else {
		ptr.Next = node
		return ptr, ok
	}
}
func (l *LinkList[T]) DelFirst() (ok bool) {
	l.headPtr, ok = l.del(l.headPtr, 0)
	return
}
func (l *LinkList[T]) DelLast() (ok bool) {
	l.headPtr, ok = l.del(l.headPtr, l.size-1)
	return
}
func (l *LinkList[T]) DelByIndex(index int) (ok bool) {
	l.headPtr, ok = l.del(l.headPtr, index)
	return ok
}

func (l *LinkList[T]) String() string {
	return l.toString(l.headPtr)
}
func (l *LinkList[T]) toString(ptr *Node[T]) string {
	if ptr == nil {
		return "null"
	} else {
		return fmt.Sprintf("%v", ptr.Value) + "->" + l.toString(ptr.Next)
	}
}

func main() {
	linkList := NewLinkList[string]()
	linkList.AddFirst("2")
	linkList.AddLast("1")
	linkList.AddLast("3")
	fmt.Println(linkList)
	linkList.AddByIndex(2, "4")
	fmt.Println(linkList)
	fmt.Println(linkList.size)
	linkList.DelByIndex(2)
	fmt.Println(linkList)
	linkList.DelFirst()
	linkList.DelLast()
	linkList.DelFirst()
	fmt.Println(linkList)
	fmt.Println(linkList.size)
}
