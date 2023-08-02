package main

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func NewLinkList[T any]() *LinkList[T] {
	linkList := &LinkList[T]{}
	linkList.headPtr = &Node[T]{
		Next: nil,
	}
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

	prev := l.headPtr
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = &Node[T]{Value: value, Next: prev.Next}

	return true
}

func (l *LinkList[T]) AddFirst(value T) bool {
	return l.Add(0, value)
}

func (l *LinkList[T]) AddLast(value T) bool {
	return l.Add(l.size, value)
}

func (l *LinkList[T]) Get(index int) (value T) {
	if index < 0 || index >= l.size {
		return
	}

	prev := l.headPtr.Next
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	value = prev.Value
	return
}

func (l *LinkList[T]) GetFirst() T {
	return l.Get(0)
}

func (l *LinkList[T]) GetLast() T {
	return l.Get(l.size - 1)
}

func (l *LinkList[T]) Del(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}
	defer func() {
		l.size--
	}()
	prev := l.headPtr
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	delNode := prev.Next
	prev.Next = delNode.Next
	delNode.Next = nil
	return true
}

func (l *LinkList[T]) DelFirst() bool {
	return l.Del(0)
}
func (l *LinkList[T]) DelLast() bool {
	return l.Del(l.size - 1)
}

func (l *LinkList[T]) String() string {
	res := ""
	prev := l.headPtr.Next
	for prev != nil {
		res += fmt.Sprintf("%v", prev.Value) + "->"
		prev = prev.Next
	}
	res += "nil"
	return res
}

func main() {
	linkList := NewLinkList[string]()
	linkList.AddLast("1")
	linkList.AddFirst("2")
	linkList.AddLast("3")
	fmt.Println(linkList.Del(1))
	fmt.Println(linkList.DelLast())
	fmt.Println(linkList)
}
