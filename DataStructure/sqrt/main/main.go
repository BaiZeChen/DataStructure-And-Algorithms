package main

import (
	"DataStructure-And-Algorithms/DataStructure/sqrt"
	"fmt"
)

func main() {
	sq := sqrt.NewSqrt([]int{-1, 3, 4, 2, 9}, func(front, rear int) int {
		return front + rear
	})
	fmt.Println(sq.RangeResult(0, 0))
}
