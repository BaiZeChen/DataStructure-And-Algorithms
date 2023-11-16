package main

import "fmt"

// 时间复杂度 O(nlogn)
func main() {
	arr := []int{989, 198, 10000, 1, 10, 8, 99, 37, 20, 3, 7, 12}
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
	m.sortRecursion(0, len(m.arr)-1, copyArr)
	// 非递归
	//m.sortIteration(copyArr)
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

// 递归的方式，自顶向下
func (m *MergeSort[T]) sortRecursion(l, r int, copyArr []T) {
	if l == r {
		return
	}

	middle := (l + r) / 2
	m.sortRecursion(l, middle, copyArr)
	m.sortRecursion(middle+1, r, copyArr)

	// 这里做一下优化，当middle的值已经小于右边最小的第一个值后，
	// 就不需要merge了
	if m.arr[middle] > m.arr[middle+1] {
		m.merge(l, middle, r, copyArr)
	}
}

// 迭代的方式，自底向上
func (m *MergeSort[T]) sortIteration(copyArr []T) {
	length := len(m.arr)
	for step := 1; step < length; step += step {
		// 这里合并的数组边界为 [i,i+step-1] [i+step, i+step+step-1]
		// 所以，l=i middle = i+step-1 r = i+step+step-1
		// 当然 i+step+step-1有可能越界，所以正确的r= min(n-1, i+step+step-1)
		for i := 0; i+step < length; i += step + step {
			// 当然这里也加一下优化：当middle的值小于middle+1的值，就不再进行merge了
			if m.arr[i+step-1] > m.arr[i+step] {
				if i+step+step-1 < length {
					m.merge(i, i+step-1, i+step+step-1, copyArr)
				} else {
					m.merge(i, i+step-1, length-1, copyArr)
				}
			}
		}
	}
}
