package main

import "fmt"

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

// 1 2 3 4
func (l *LinkList[T]) Add(index int, value T) bool {
	if index < 0 || index > l.size {
		return false
	}
	defer func() {
		l.size++
	}()

	ptr := l.headPtr
	if index == 0 {
		if l.size == 0 {
			l.headPtr = &Node[T]{
				Value: value,
			}
		} else {
			l.headPtr = &Node[T]{
				Value: value,
				Next:  ptr,
			}
		}
	} else {
		for i := 0; i < index-1; i++ {
			ptr = ptr.Next
		}
		ptr.Next = &Node[T]{
			Value: value,
			Next:  ptr.Next,
		}
	}

	return true
}
func (l *LinkList[T]) AddFirst(value T) bool {
	return l.Add(0, value)
}
func (l *LinkList[T]) AddLast(value T) bool {
	return l.Add(l.size, value)
}

func (l *LinkList[T]) Get(index int) (res T) {
	if index < 0 || index >= l.size {
		return
	}

	ptr := l.headPtr
	for i := 0; i < index; i++ {
		ptr = ptr.Next
	}
	res = ptr.Value
	return
}
func (l *LinkList[T]) GetFirst() T {
	return l.Get(0)
}
func (l *LinkList[T]) GetLast() T {
	return l.Get(l.size - 1)
}

// 1 2
func (l *LinkList[T]) Del(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}
	defer func() {
		l.size--
	}()

	ptr := l.headPtr
	if index == 0 {
		l.headPtr = l.headPtr.Next
		ptr.Next = nil
	} else {
		for i := 0; i < index-1; i++ {
			ptr = ptr.Next
		}
		delNode := ptr.Next
		ptr.Next = delNode.Next
		delNode.Next = nil
	}
	return true
}
func (l *LinkList[T]) DelFirst() bool {
	return l.Del(0)
}
func (l *LinkList[T]) DelLast() bool {
	return l.Del(l.size - 1)
}

func (l *LinkList[T]) Size() int {
	return l.size
}

func (l *LinkList[T]) String() string {
	res := ""
	prev := l.headPtr
	for prev != nil {
		res += fmt.Sprintf("%v", prev.Value) + "->"
		prev = prev.Next
	}
	res += "nil"
	return res
}

// 用链表实现栈
func NewStack[T any]() *LinkListStack[T] {
	return &LinkListStack[T]{container: NewLinkList[T]()}
}

type LinkListStack[T any] struct {
	container *LinkList[T]
}

func (s *LinkListStack[T]) Push(value T) {
	s.container.AddFirst(value)
}
func (s *LinkListStack[T]) Pop() (res T) {
	if s.container.size == 0 {
		return
	}
	res = s.container.GetFirst()
	s.container.DelFirst()
	return
}

func main() {
	//linkList := NewLinkList[string]()
	//linkList.AddLast("1")
	//linkList.AddFirst("2")
	//linkList.AddLast("3")
	//fmt.Println(linkList.Del(1))
	//fmt.Println(linkList.DelLast())
	//fmt.Println(linkList.AddLast("3"))
	//linkList.DelFirst()
	//fmt.Println(linkList)

	// 栈测试
	stack := NewStack[string]()
	stack.Push("2")
	stack.Push("3")
	stack.Push("7")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

}
