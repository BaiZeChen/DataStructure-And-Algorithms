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

func (l *LinkList[T]) Add(index int, value T) bool {
	if index < 0 || index > l.size {
		return false
	}
	defer func() {
		l.size++
	}()

	node := &Node[T]{
		Value: value,
	}
	if l.size == 0 {
		l.headPtr = node
		return true
	} else if index == 0 {
		tmp := l.headPtr
		l.headPtr = node
		node.Next = tmp
		return true
	} else {
		return l.recursionAdd(index, l.headPtr, node)
	}
}
func (l *LinkList[T]) recursionAdd(index int, ptr *Node[T], addNode *Node[T]) bool {
	if index == 1 {
		node := ptr.Next
		ptr.Next = addNode
		addNode.Next = node
		return true
	} else {
		index--
		return l.recursionAdd(index, ptr.Next, addNode)
	}
}
func (l *LinkList[T]) AddFirst(value T) bool {
	return l.Add(0, value)
}
func (l *LinkList[T]) AddLast(value T) bool {
	return l.Add(l.size, value)
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

func (l *LinkList[T]) Del(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}
	defer func() {
		l.size--
	}()

	if index == 0 {
		delNode := l.headPtr
		l.headPtr = l.headPtr.Next
		delNode.Next = nil
		return true
	} else {
		return l.recursionDel(index, l.headPtr)
	}
}
func (l *LinkList[T]) DelFirst() bool {
	return l.Del(0)
}
func (l *LinkList[T]) DelLast() bool {
	return l.Del(l.size - 1)
}
func (l *LinkList[T]) recursionDel(index int, ptr *Node[T]) bool {
	if index == 1 {
		delNode := ptr.Next
		ptr.Next = ptr.Next.Next
		delNode.Next = nil
		return true
	} else {
		index--
		return l.recursionDel(index, ptr.Next)
	}
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
	linkList.AddLast("1")
	linkList.AddFirst("2")
	linkList.AddLast("3")
	linkList.Add(2, "4")
	fmt.Println(linkList)
	linkList.Del(2)
	fmt.Println(linkList)
	linkList.DelFirst()
	linkList.DelLast()
	fmt.Println(linkList)
	fmt.Println(linkList.size)
}
