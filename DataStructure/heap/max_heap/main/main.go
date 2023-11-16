package main

import (
	"DataStructure-And-Algorithms/DataStructure/heap/max_heap"
	"fmt"
)

func main() {
	//heap := &max_heap.Heap[int]{}
	//heap.Add(9)
	//heap.Add(1)
	//heap.Add(3)
	//heap.Add(5)
	//heap.Add(4)
	//heap.Add(13)
	//fmt.Println(heap.Pop())
	//fmt.Println(heap.Pop())
	//fmt.Println(heap.Pop())
	//fmt.Println(heap.Pop())
	//fmt.Println(heap.Pop())
	//fmt.Println(heap.Pop())
	//fmt.Println(heap.Pop())

	heap := max_heap.NewHeap([]int{2, 9, 10, 17, 20, 1, 16, 75})
	fmt.Println(heap)
}
