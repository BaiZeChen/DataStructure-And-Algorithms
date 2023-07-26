package main

import (
	"DataStructure-And-Algorithms/DataStructure/trends_array"
	"fmt"
)

func main() {
	slice := trends_array.NewSlice(2)
	slice.Append(1.2, 2.3, 3, 4.6, 5, 6, 7, 8, 9)
	slice.Insert(0, 0)
	fmt.Println(slice)
	slice.RemoveLast()
	slice.RemoveLast()
	slice.RemoveLast()
	slice.RemoveLast()
	slice.RemoveLast()
	fmt.Println(slice)
}
