package main

import "fmt"

// 时间复杂度 O(nlogn)
func main() {
	arr := []int{989, 198, 10000, 1, 10, 8, 99, 37, 20}
	sort := NewMergeSort[int](arr)
	fmt.Println(sort.Sort())
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func NewMergeSort[T Ordered](arr []T) *MergeSort[T] {
	return &MergeSort[T]{arr: arr}
}

type MergeSort[T Ordered] struct {
	arr []T
}

func (m *MergeSort[T]) Sort() []T {
	// 为了避免每次merge时都开辟空间，这里直接先定义好
	// 这里做了一下内存开辟的优化
	copyArr := make([]T, len(m.arr))
	m.sort(0, len(m.arr)-1, copyArr)
	return m.arr
}

func (m *MergeSort[T]) merge(l, middle, r int, copyArr []T) {
	copy(copyArr[l:r+1], m.arr[l:r+1])

	lStart, rStart := l, middle+1
	for i := l; i <= r; i++ {
		if lStart > middle {
			// 代表左边的数组已经没有值了
			m.arr[i] = copyArr[rStart]
			rStart++
		} else if rStart > r {
			m.arr[i] = copyArr[lStart]
			lStart++
		} else if copyArr[lStart] < copyArr[rStart] {
			m.arr[i] = copyArr[lStart]
			lStart++
		} else {
			m.arr[i] = copyArr[rStart]
			rStart++
		}
	}
}

func (m *MergeSort[T]) sort(l, r int, copyArr []T) {
	if l >= r {
		return
	}

	middle := (l + r) / 2
	m.sort(l, middle, copyArr)
	m.sort(middle+1, r, copyArr)

	// 这里做一下优化，当middle的值已经小于右边最小的第一个值后，
	// 就不需要merge了
	if m.arr[middle] > m.arr[middle+1] {
		m.merge(l, middle, r, copyArr)
	}
}
