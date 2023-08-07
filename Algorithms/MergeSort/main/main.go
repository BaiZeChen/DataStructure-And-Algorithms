package main

import "fmt"

func main() {
	arr := []int{20, 7, 9, 6, 10, 8, 21, 24, 20}
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
	m.sort(0, len(m.arr)-1)
	return m.arr
}

func (m *MergeSort[T]) merge(l, middle, r int) {
	tmp := make([]T, (r-l)+1)
	copy(tmp, m.arr[l:r+1])

	lStart, rStrat := l, middle+1
	for i := l; i <= r; i++ {
		if lStart > middle {
			// 代表左边的数组已经没有值了
			m.arr[i] = tmp[rStrat-l]
			rStrat++
		} else if rStrat > r {
			m.arr[i] = tmp[lStart-l]
			lStart++
		} else if tmp[lStart-l] < tmp[rStrat-l] {
			m.arr[i] = tmp[lStart-l]
			lStart++
		} else {
			m.arr[i] = tmp[rStrat-l]
			rStrat++
		}
	}
}

func (m *MergeSort[T]) sort(l, r int) {
	if l >= r {
		return
	}

	middle := (l + r) / 2
	m.sort(l, middle)
	m.sort(middle+1, r)

	m.merge(l, middle, r)
}
