package binary_search_tree

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type Avl[T Ordered] struct {
	root *Node[T]
}
type Node[T Ordered] struct {
	Value      T
	height     int
	LeftChild  *Node[T]
	RightChild *Node[T]
}

// 对节点y进行向右旋转操作，返回旋转后新的根节点x
//
//	      y								 x
//	     /  \         (向右旋转)			/  \
//	    x    T4      - - - - - - >	   z     y
//	   / \							  /  \	/ \
//	  z   T3						 T1  T2 T3 T4
//	 /  \
//	T1   T2
func (a *Avl[T]) rightRotate(node *Node[T]) *Node[T] {
	x := node.LeftChild
	t3 := x.RightChild
	x.RightChild = node
	node.LeftChild = t3

	// 更换高度（只更换x,y的高度）
	if a.getHeight(node.LeftChild) > a.getHeight(node.RightChild) {
		node.height = 1 + a.getHeight(node.LeftChild)
	} else {
		node.height = 1 + a.getHeight(node.RightChild)
	}
	if a.getHeight(x.LeftChild) > a.getHeight(x.RightChild) {
		x.height = 1 + a.getHeight(x.LeftChild)
	} else {
		x.height = 1 + a.getHeight(x.RightChild)
	}
	return x
}

// 对节点y进行向左旋转操作，返回旋转后新的根节点x
//
//		      y								  x
//		     /  \         (向左旋转)			 / \
//		    T4   x      - - - - - - >       y   z
//		   	    / \	                       / \  /\
//		       T3  z                      T4 T3T1 T2
//		          / \
//	 	     T1  T2
func (a *Avl[T]) leftRotate(node *Node[T]) *Node[T] {
	x := node.RightChild
	T3 := x.LeftChild
	x.LeftChild = node
	node.RightChild = T3

	// 只更改x,y高度
	if a.getHeight(node.LeftChild) > a.getHeight(node.RightChild) {
		node.height = 1 + a.getHeight(node.LeftChild)
	} else {
		node.height = 1 + a.getHeight(node.RightChild)
	}
	if a.getHeight(x.LeftChild) > a.getHeight(x.RightChild) {
		x.height = 1 + a.getHeight(x.LeftChild)
	} else {
		x.height = 1 + a.getHeight(x.RightChild)
	}
	return x
}

func (a *Avl[T]) isBST() bool {
	list := a.getOrderList()
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			return false
		}
	}
	return true
}

// 判断这棵树是否是平衡的
func (a *Avl[T]) isBalanced(root *Node[T]) bool {
	if root == nil {
		return true
	}

	balanceFactor := a.getBalanceFactor(root)
	if balanceFactor < 0 {
		balanceFactor = -balanceFactor
	}
	if balanceFactor > 1 {
		return false
	}
	return a.isBalanced(root.LeftChild) && a.isBalanced(root.RightChild)
}

func (a *Avl[T]) getHeight(node *Node[T]) int {
	if node == nil {
		return 0
	}
	return node.height
}

// 获取某个节点的平衡因子
func (a *Avl[T]) getBalanceFactor(node *Node[T]) int {
	if node == nil {
		return 0
	}
	return a.getHeight(node.LeftChild) - a.getHeight(node.RightChild)
}

func (a *Avl[T]) Add(value T) {
	// 递归方式
	a.root = a.addRecursion(a.root, value)
}

func (a *Avl[T]) addRecursion(node *Node[T], value T) *Node[T] {
	if node == nil {
		return &Node[T]{
			Value:  value,
			height: 1,
		}
	}

	// 这里相等不处理
	if node.Value < value {
		node.LeftChild = a.addRecursion(node.LeftChild, value)
	} else if node.Value > value {
		node.RightChild = a.addRecursion(node.RightChild, value)
	}

	// 计算高度
	if a.getHeight(node.LeftChild) > a.getHeight(node.RightChild) {
		node.height = 1 + a.getHeight(node.LeftChild)
	} else {
		node.height = 1 + a.getHeight(node.RightChild)
	}
	// rr 旋转
	if a.getBalanceFactor(node) > 1 && a.getBalanceFactor(node.LeftChild) >= 0 {
		return a.rightRotate(node)
	}
	// lr旋转
	if a.getBalanceFactor(node) > 1 && a.getBalanceFactor(node.LeftChild) < 0 {
		// 先左旋转
		node.LeftChild = a.leftRotate(node.LeftChild)
		return a.rightRotate(node)
	}
	// ll 旋转
	if a.getBalanceFactor(node) < -1 && a.getBalanceFactor(node.RightChild) <= 0 {
		return a.leftRotate(node)
	}
	// rl 旋转
	if a.getBalanceFactor(node) < -1 && a.getBalanceFactor(node.RightChild) > 0 {
		node.RightChild = a.rightRotate(node.RightChild)
		return a.leftRotate(node)
	}

	return node
}

func (a *Avl[T]) addIteration(rootNode *Node[T], value T) {
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

//func (s *SBT[T]) inOrderErgodic(node *Node[T]) {
//	if node == nil {
//		return
//	}
//
//	s.inOrderErgodic(node.LeftChild)
//	fmt.Println(node.Value)
//	s.inOrderErgodic(node.RightChild)
//}

// 中序遍历，是顺序的
func (a *Avl[T]) inOrderErgodic(node *Node[T], orderMap map[T]int) {
	if node == nil {
		return
	}

	a.inOrderErgodic(node.LeftChild, orderMap)
	orderMap[node.Value] = 1
	a.inOrderErgodic(node.RightChild, orderMap)
}

func (a *Avl[T]) getOrderList() []T {
	orderMap := make(map[T]int)
	a.inOrderErgodic(a.root, orderMap)
	var list []T
	for key := range orderMap {
		list = append(list, key)
	}
	return list
}

// 层序遍历（这里只说明伪代码，就不写了）
/*
	①：首先声明一个队列，然后先将头结点进行入队
	②：进行循环，循环条件队列是否为空
	③：出队打印
	④：如果左节点不为空就入队，如果右节点不为空就入队
*/

// FindMin 查找最小节点，并返回
func (a *Avl[T]) FindMin(node *Node[T]) *Node[T] {
	if node.LeftChild == nil {
		return node
	}
	return a.FindMin(node.LeftChild)
}

// FindMax 查找最大节点，并返回
func (a *Avl[T]) FindMax(node *Node[T]) *Node[T] {
	if node.RightChild == nil {
		return node
	}
	return a.FindMax(node.RightChild)
}

func (a *Avl[T]) RemoveMinAndReturn(node *Node[T]) *Node[T] {
	minNode := a.FindMin(node)
	node = a.removeMin(node)
	return minNode
}
func (a *Avl[T]) RemoveMaxAndReturn(node *Node[T]) *Node[T] {
	minNode := a.FindMax(node)
	node = a.removeMax(node)
	return minNode
}
func (a *Avl[T]) removeMin(node *Node[T]) *Node[T] {
	if node.LeftChild == nil {
		return node.RightChild
	}

	node.LeftChild = a.removeMin(node.LeftChild)
	return node
}
func (a *Avl[T]) removeMax(node *Node[T]) *Node[T] {
	if node.RightChild == nil {
		return node.LeftChild
	}
	node.RightChild = a.removeMax(node.RightChild)
	return node
}
func (a *Avl[T]) RemoveAny(value T) {
	a.root = a.remove(a.root, value)
}
func (a *Avl[T]) remove(node *Node[T], value T) *Node[T] {
	var finalNode *Node[T]
	if node == nil {
		return nil
	} else if node.Value < value {
		node.RightChild = a.remove(node.RightChild, value)
		finalNode = node
	} else if node.Value > value {
		node.LeftChild = a.remove(node.LeftChild, value)
		finalNode = node
	} else {
		if node.LeftChild == nil {
			rightChild := node.RightChild
			node.RightChild = nil
			finalNode = rightChild
		} else if node.RightChild == nil {
			leftChild := node.LeftChild
			node.LeftChild = nil
			finalNode = leftChild
		} else {
			// 有两个子结点,这里找后继节点（就是右子树最小的节点）
			succeedNode := a.FindMin(node.RightChild)
			succeedNode.RightChild = a.remove(node.RightChild, succeedNode.Value)
			succeedNode.LeftChild = node.LeftChild
			node.LeftChild, node.RightChild = nil, nil
			finalNode = succeedNode
		}
	}

	if finalNode == nil {
		return nil
	}

	// 计算高度
	if a.getHeight(finalNode.LeftChild) > a.getHeight(finalNode.RightChild) {
		finalNode.height = 1 + a.getHeight(finalNode.LeftChild)
	} else {
		finalNode.height = 1 + a.getHeight(finalNode.RightChild)
	}
	// rr 旋转
	if a.getBalanceFactor(finalNode) > 1 && a.getBalanceFactor(finalNode.LeftChild) >= 0 {
		return a.rightRotate(finalNode)
	}
	// lr旋转
	if a.getBalanceFactor(finalNode) > 1 && a.getBalanceFactor(finalNode.LeftChild) < 0 {
		// 先左旋转
		finalNode.LeftChild = a.leftRotate(finalNode.LeftChild)
		return a.rightRotate(finalNode)
	}
	// ll 旋转
	if a.getBalanceFactor(finalNode) < -1 && a.getBalanceFactor(finalNode.RightChild) <= 0 {
		return a.leftRotate(finalNode)
	}
	// rl 旋转
	if a.getBalanceFactor(finalNode) < -1 && a.getBalanceFactor(finalNode.RightChild) > 0 {
		finalNode.RightChild = a.rightRotate(finalNode.RightChild)
		return a.leftRotate(finalNode)
	}
	return finalNode
}
