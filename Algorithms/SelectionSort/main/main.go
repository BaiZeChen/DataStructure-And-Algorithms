package main

import (
	"errors"
	"fmt"
	"math"
)

// 时间复杂度，O(n^2)
func main() {
	ints, err := task([]int{2, 3, 1, 4, 5, 10, 7, 8})
	if err != nil {
		panic(err)
	}
	fmt.Println(ints)
}

// AdvancedSort 选择排序（有B格的写法）
func AdvancedSort(arr []int) ([]int, error) {
	length := len(arr)
	if len(arr) == 0 {
		return nil, errors.New("没元素，你排个毛线~~~")
	}
	for i := 0; i < length-1; i++ {
		minKey := i
		for j := i; j < length; j++ {
			if arr[minKey] > arr[j] {
				minKey = j
			}
		}
		arr[i], arr[minKey] = arr[minKey], arr[i]
	}
	return arr, nil
}

// 选择排序（普通版的）
func sort(arr []int) ([]int, error) {
	length := len(arr)
	if len(arr) == 0 {
		return nil, errors.New("没元素，你排个毛线~~~")
	}
	for i := 0; i < length-1; i++ {
		minValue := math.MaxInt
		minKey := -1
		for j := i; j < length; j++ {
			if minValue > arr[j] {
				minValue = arr[j]
				minKey = j
			}
		}
		arr[minKey] = arr[i]
		arr[i] = minValue
	}
	return arr, nil
}

// 慕课网作业
func task(arr []int) ([]int, error) {
	length := len(arr)
	if len(arr) == 0 {
		return nil, errors.New("没元素，你排个毛线~~~")
	}

	for i := length - 1; i > 0; i-- {
		minKey := i
		for j := i; j >= 0; j-- {
			if arr[minKey] > arr[j] {
				minKey = j
			}
		}
		arr[i], arr[minKey] = arr[minKey], arr[i]
	}
	return arr, nil
}
