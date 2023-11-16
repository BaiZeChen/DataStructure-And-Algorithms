package max_heap

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// 二叉堆的性质：
// 1. 它是一个完全二叉树
// 2. 所有节点元素都大于等于其左右孩子树（最大堆）或者都小于等于其左右子树（最小堆）
// 3. 可以采用数组存储接口

func NewHeap[T Ordered](arr []T) *Heap[T] {
	heap := &Heap[T]{}
	heap.heapify(arr)
	return heap
}

type Heap[T Ordered] struct {
	data []T
}

// 将任意一个数组变成最大堆
func (h *Heap[T]) heapify(arr []T) {
	h.data = arr
	for i := h.parentsIndex(h.Size() - 1); i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}

func (h *Heap[T]) Empty() bool {
	if h.Size() == 0 {
		return true
	}
	return false
}

func (h *Heap[T]) Add(value T) {
	h.data = append(h.data, value)
	h.siftUp(h.Size() - 1)
}

func (h *Heap[T]) siftUp(index int) {
	for {
		parent := h.parentsIndex(index)
		if parent >= 0 && h.data[parent] < h.data[index] {
			h.data[parent], h.data[index] = h.data[index], h.data[parent]
			index = parent
		} else {
			break
		}
	}
}

func (h *Heap[T]) Pop() (value T) {
	if h.Size() == 0 {
		return
	}
	value = h.data[0]
	h.data[0] = h.data[h.Size()-1]
	h.data = h.data[:h.Size()-1]
	h.siftDown(0)
	return
}

func (h *Heap[T]) siftDown(index int) {
	for {
		child := h.leftChild(index)
		if child <= 0 {
			break
		}
		if child+1 < h.Size() && h.data[child] < h.data[child+1] {
			child++
		}
		if h.data[child] < h.data[index] {
			break
		}
		h.data[index], h.data[child] = h.data[child], h.data[index]
		index = child
	}
}

func (h *Heap[T]) parentsIndex(index int) int {
	if index == 0 {
		// 根节点没有父节点
		return -1
	}
	return (index - 1) / 2
}

func (h *Heap[T]) leftChild(index int) int {
	i := index*2 + 1
	if i >= h.Size() {
		return -1
	}
	return i
}

func (h *Heap[T]) rightChild(index int) int {
	i := index*2 + 2
	if i >= h.Size() {
		return -1
	}
	return i
}
