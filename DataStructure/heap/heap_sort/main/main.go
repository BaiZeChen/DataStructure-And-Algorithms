package main

import (
	"DataStructure-And-Algorithms/DataStructure/heap/heap_sort"
	"fmt"
)

func main() {
	arr := []int{9, 1, 100, 7, 14, 88, 65}
	fmt.Println(heap_sort.HeapSort(arr))
}
