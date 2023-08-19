package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	tmp := []int{6, 9, 1, 13, 8, 5, 2}
	sortV2(tmp, 0, len(tmp)-1)
	fmt.Println(tmp)

}

// 第一版快排版本
func sortV1(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partitionV1(arr, l, r)
	sortV1(arr, l, p-1)
	sortV1(arr, p+1, r)
}
func partitionV1(arr []int, l, r int) int {
	j := l
	// 左边区间：arr[l+1:j]<v  右边区间：arr[j+1:i-1] >=v
	// 假设索引为l的为比较值
	for i := l + 1; i <= r; i++ {
		if arr[l] > arr[i] {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[j], arr[l] = arr[l], arr[j]
	return j
}

func sortV2(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partitionV2(arr, l, r)
	sortV1(arr, l, p-1)
	sortV1(arr, p+1, r)
}

// 不再默认取第一个下标，而是随机取下标
func partitionV2(arr []int, l, r int) int {
	result, _ := rand.Int(rand.Reader, big.NewInt(int64((r-1)+1)))
	seed := int(int64(l) + result.Int64())
	arr[seed], arr[l] = arr[l], arr[seed]

	j := l
	for i := l + 1; i <= r; i++ {
		if arr[l] > arr[i] {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[j], arr[l] = arr[l], arr[j]
	return j
}
