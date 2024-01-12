package segment_tree

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type Merge[T Ordered] func(leftChild, rightChild T) T

func NewSegmentTree[T Ordered](arr []T, merge Merge[T]) *SegmentTree[T] {
	//初始化树的长度为4n
	segment := &SegmentTree[T]{
		treeStruct: make([]T, 4*len(arr)),
		arr:        arr,
		merge:      merge,
	}
	segment.create()
	return segment
}

type SegmentTree[T Ordered] struct {
	treeStruct []T
	arr        []T
	merge      Merge[T]
}

func (s *SegmentTree[T]) create() {
	s.segmentTreeCreate(0, 0, len(s.arr)-1)
}

func (s *SegmentTree[T]) segmentTreeCreate(index, leftIndex, rightIndex int) {
	if leftIndex == rightIndex {
		s.treeStruct[index] = s.arr[rightIndex]
		return
	}

	leftChildIndex := s.leftChild(index)
	rightChildIndex := s.rightChild(index)
	middle := (leftIndex + rightIndex) / 2

	s.segmentTreeCreate(leftChildIndex, leftIndex, middle)
	s.segmentTreeCreate(rightChildIndex, middle+1, rightIndex)

	s.treeStruct[index] = s.merge(s.treeStruct[leftChildIndex], s.treeStruct[rightChildIndex])
}

func (s *SegmentTree[T]) leftChild(index int) int {
	i := index*2 + 1
	return i
}

func (s *SegmentTree[T]) rightChild(index int) int {
	i := index*2 + 2
	return i
}

func (s *SegmentTree[T]) GetTree() []T {
	return s.treeStruct
}

func (s *SegmentTree[T]) Query(queryLeft, queryRight int) T {
	return s.query(0, 0, len(s.arr)-1, queryLeft, queryRight)
}

func (s *SegmentTree[T]) query(treeIndex, leftIndex, rightIndex, queryLeft, queryRight int) T {
	if leftIndex == queryLeft && rightIndex == queryRight {
		return s.treeStruct[treeIndex]
	}

	leftChildIndex := s.leftChild(treeIndex)
	rightChildIndex := s.rightChild(treeIndex)
	middle := (leftIndex + rightIndex) / 2

	if middle >= queryRight {
		return s.query(leftChildIndex, leftIndex, middle, queryLeft, queryRight)
	} else if middle+1 <= queryLeft {
		return s.query(rightChildIndex, middle+1, rightIndex, queryLeft, queryRight)
	} else {
		leftRes := s.query(leftChildIndex, leftIndex, middle, queryLeft, middle)
		rightRes := s.query(rightChildIndex, middle+1, rightIndex, middle+1, queryRight)
		return s.merge(leftRes, rightRes)
	}
}

func (s *SegmentTree[T]) Update(index int, value T) {
	s.arr[index] = value
	s.update(value, index, 0, 0, len(s.arr)-1)
}
func (s *SegmentTree[T]) update(value T, index, treeIndex, leftIndex, rightIndex int) {
	if leftIndex == rightIndex {
		s.treeStruct[leftIndex] = value
		return
	}

	leftChildIndex := s.leftChild(treeIndex)
	rightChildIndex := s.rightChild(treeIndex)
	middle := (leftIndex + rightIndex) / 2

	if middle >= index {
		s.update(value, index, leftChildIndex, leftIndex, middle)
	} else {
		s.update(value, index, rightChildIndex, middle+1, rightIndex)
	}
	s.treeStruct[treeIndex] = s.merge(s.treeStruct[leftChildIndex], s.treeStruct[rightChildIndex])
}
