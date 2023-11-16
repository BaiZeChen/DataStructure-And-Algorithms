package main

import "fmt"

func main() {
	fmt.Println(bubblingSort([]int{22, 1, 8, 16, 25, 7, 10}))
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// 比较高级的冒泡算法
func bubblingSort[T Ordered](arr []T) []T {
	length := len(arr)
	for i := 0; i < length-1; {
		swapLastIndex := 0
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapLastIndex = j + 1
			}
		}
		i = length - swapLastIndex
	}
	return arr
}
