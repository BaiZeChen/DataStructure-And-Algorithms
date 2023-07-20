package main

import (
	"errors"
	"fmt"
)

// 时间复杂度 o(n^2)，对于比如有序数据则是o(n)
func main() {
	ints, err := test([]int{2, 3, 1, 4, 5, 10, 7, 8})
	if err != nil {
		panic(err)
	}
	fmt.Println(ints)
}

// 插入排序
func sort(arr []int) ([]int, error) {
	length := len(arr)
	if len(arr) == 0 {
		return nil, errors.New("没元素，你排个毛线~~~")
	}

	for i := 1; i < length; i++ {
		for j := i; j >= 0; j-- {
			if j > 0 && arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				break
			}
		}
	}
	return arr, nil
}

// 优化的排序
func optimizeSort(arr []int) ([]int, error) {
	length := len(arr)
	if len(arr) == 0 {
		return nil, errors.New("没元素，你排个毛线~~~")
	}

	for i := 1; i < length; i++ {
		tmp := arr[i]
		lastIndex := -1
		for j := i - 1; j >= 0; j-- {
			if arr[j] > tmp {
				lastIndex = j
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		if !(lastIndex < 0) {
			arr[lastIndex] = tmp
		}
	}
	return arr, nil
}

func test(arr []int) ([]int, error) {
	length := len(arr)
	if len(arr) == 0 {
		return nil, errors.New("没元素，你排个毛线~~~")
	}

	for i := length - 2; i >= 0; i-- {
		tmp := arr[i]
		lastIndex := -1
		for j := i + 1; j < length; j++ {
			if arr[j] > tmp {
				arr[j-1] = arr[j]
				lastIndex = j
			} else {
				break
			}
		}
		if !(lastIndex < 0) {
			arr[lastIndex] = tmp
		}
	}
	return arr, nil
}
