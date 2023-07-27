package main

import (
	"DataStructure-And-Algorithms/DataStructure/trends_array"
	"fmt"
)

func main() {
	slice := trends_array.NewSlice[int](10)
	slice.Append(1, 2, 3, 4, 5, 6, 7, 8, 9)
	slice.Insert(0, 0)
	fmt.Println(slice)
	slice.RemoveLast()
	slice.RemoveLast()
	slice.RemoveLast()
	slice.RemoveLast()
	slice.RemoveLast()
	fmt.Println(slice)
}
