package binary_search_tree

import "fmt"

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type SBT[T Ordered] struct {
	root *Node[T]
}
type Node[T Ordered] struct {
	Value      T
	LeftChild  *Node[T]
	RightChild *Node[T]
}

func (s *SBT[T]) Add(value T) {
	// 递归方式
	s.root = s.addRecursion(s.root, value)
}

func (s *SBT[T]) addRecursion(node *Node[T], value T) *Node[T] {
	if node == nil {
		return &Node[T]{
			Value: value,
		}
	}

	// 这里相等不处理
	if node.Value < value {
		node.LeftChild = s.addRecursion(node.LeftChild, value)
	} else if node.Value > value {
		node.RightChild = s.addRecursion(node.RightChild, value)
	}
	return node
}

func (s *SBT[T]) addIteration(rootNode *Node[T], value T) {
	node := &Node[T]{
		Value: value,
	}
	if rootNode == nil {
		// 代表还没有根节点
		rootNode = node
		return
	}
	pointer := rootNode

	for {
		if pointer.Value < value {
			if pointer.LeftChild == nil {
				pointer.LeftChild = node
				break
			}
			pointer = pointer.LeftChild
		} else if pointer.Value > value {
			if pointer.RightChild == nil {
				pointer.RightChild = node
				break
			}
			pointer = pointer.RightChild
		}
	}
}

func (s *SBT[T]) PreorderErgodicRecursion() {
	s.preorderRecursion(s.root)
}

// 前序递归遍历
func (s *SBT[T]) preorderRecursion(node *Node[T]) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)
	s.preorderRecursion(node.LeftChild)
	s.preorderRecursion(node.RightChild)
}

// 前序迭代遍历（这里只说明伪代码，就不写了）
/*
	①：首先声明一个栈，然后先将头结点压栈
	②：进行循环，循环条件栈是否为空
	③：出栈打印
	④：如果右节点不为空，右节点压栈，如果左节点不为空，左节点压栈
*/

// InOrderErgodic 中序递归遍历
func (s *SBT[T]) InOrderErgodic() {
	s.inOrderErgodic(s.root)
}

func (s *SBT[T]) inOrderErgodic(node *Node[T]) {
	if node == nil {
		return
	}

	s.inOrderErgodic(node.LeftChild)
	fmt.Println(node.Value)
	s.inOrderErgodic(node.RightChild)
}

// PostorderErgodic 后序递归遍历
func (s *SBT[T]) PostorderErgodic() {
	s.postorderErgodic(s.root)
}
func (s *SBT[T]) postorderErgodic(node *Node[T]) {
	if node == nil {
		return
	}

	s.postorderErgodic(node.LeftChild)
	s.postorderErgodic(node.RightChild)
	fmt.Println(node.Value)
}

// 层序遍历（这里只说明伪代码，就不写了）
/*
	①：首先声明一个队列，然后先将头结点进行入队
	②：进行循环，循环条件队列是否为空
	③：出队打印
	④：如果左节点不为空就入队，如果右节点不为空就入队
*/

// FindMin 查找最小节点，并返回
func (s *SBT[T]) FindMin(node *Node[T]) *Node[T] {
	if node.LeftChild == nil {
		return node
	}
	return s.FindMin(node.LeftChild)
}

// FindMax 查找最大节点，并返回
func (s *SBT[T]) FindMax(node *Node[T]) *Node[T] {
	if node.RightChild == nil {
		return node
	}
	return s.FindMax(node.RightChild)
}

func (s *SBT[T]) RemoveMinAndReturn(node *Node[T]) *Node[T] {
	minNode := s.FindMin(node)
	node = s.removeMin(node)
	return minNode
}
func (s *SBT[T]) RemoveMaxAndReturn(node *Node[T]) *Node[T] {
	minNode := s.FindMax(node)
	node = s.removeMax(node)
	return minNode
}
func (s *SBT[T]) removeMin(node *Node[T]) *Node[T] {
	if node.LeftChild == nil {
		return node.RightChild
	}

	node.LeftChild = s.removeMin(node.LeftChild)
	return node
}
func (s *SBT[T]) removeMax(node *Node[T]) *Node[T] {
	if node.RightChild == nil {
		return node.LeftChild
	}
	node.RightChild = s.removeMax(node.RightChild)
	return node
}
func (s *SBT[T]) RemoveAny(value T) {
	s.root = s.remove(s.root, value)
}
func (s *SBT[T]) remove(node *Node[T], value T) *Node[T] {
	if node == nil {
		return nil
	} else if node.Value < value {
		node.RightChild = s.remove(node.RightChild, value)
		return node
	} else if node.Value > value {
		node.LeftChild = s.remove(node.LeftChild, value)
		return node
	} else {
		if node.LeftChild == nil {
			rightChild := node.RightChild
			node.RightChild = nil
			return rightChild
		} else if node.RightChild == nil {
			leftChild := node.LeftChild
			node.LeftChild = nil
			return leftChild
		} else {
			// 有两个子结点,这里找后继节点（就是右子树最小的节点）
			succeedNode := s.FindMin(node.RightChild)
			succeedNode.RightChild = s.removeMin(node.RightChild)
			succeedNode.LeftChild = node.LeftChild
			node.LeftChild, node.RightChild = nil, nil
			return succeedNode
		}
	}
}
