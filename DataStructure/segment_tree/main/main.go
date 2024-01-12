package main

import (
	"DataStructure-And-Algorithms/DataStructure/segment_tree"
	"fmt"
)

func main() {
	s := segment_tree.NewSegmentTree([]int{-1, 3, 4, 2, 9}, func(leftChild, rightChild int) int {
		return rightChild + leftChild
	})
	fmt.Println(s.Query(0, 3))
}
