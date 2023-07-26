package main

import "fmt"

func NewStack(size int) *Stack {
	stack := &Stack{}
	stack.init(size)
	return stack
}

type Stack struct {
	arr []string
}

func (s *Stack) init(size int) {
	s.arr = make([]string, 0, size)
}

func (s *Stack) Push(value string) {
	s.arr = append(s.arr, value)
}

func (s *Stack) Pop() string {
	if s.getSize() == 0 {
		return ""
	}
	length := len(s.arr) - 1
	value := s.arr[length]
	s.arr = s.arr[:length]
	return value
}

func (s *Stack) Peek() string {
	if s.getSize() == 0 {
		return ""
	}
	return s.arr[len(s.arr)-1]
}

func (s *Stack) IsEmpty() bool {
	return s.getSize() == 0
}

func (s *Stack) getSize() int {
	return len(s.arr)
}

func main() {
	fmt.Println(test("(){]{}"))
}

// leetcode 20题
func test(s string) bool {
	stack := NewStack(len(s))
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
