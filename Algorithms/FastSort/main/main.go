package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// 快速排序：
// 最坏：O(n^2) 概率非常低，平均：O(nlogn)
func main() {
	tmp := []int{6, 10, 1, 13, 8, 5, 2, 87, 23, 37, 13, 13, 17, 23, 10, 9, 6, 6, 6, 1, 1, 1}
	sortV4(tmp, 0, len(tmp)-1)
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
	sortV2(arr, l, p-1)
	sortV2(arr, p+1, r)
}

// 不再默认取第一个下标，而是随机取下标
func partitionV2(arr []int, l, r int) int {
	result, _ := rand.Int(rand.Reader, big.NewInt(int64((r-l)+1)))
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

func sortV3(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partitionV3(arr, l, r)
	sortV3(arr, l, p-1)
	sortV3(arr, p+1, r)
}

// 双路快排
func partitionV3(arr []int, l, r int) int {
	result, _ := rand.Int(rand.Reader, big.NewInt(int64((r-l)+1)))
	seed := int(int64(l) + result.Int64())
	arr[seed], arr[l] = arr[l], arr[seed]

	// 左边区间：arr[l+1:lte-1]<v  右边区间：arr[gte+1:r] >=v
	lte, gte := l+1, r
	for true {
		for lte <= gte && arr[l] > arr[lte] {
			lte++
		}
		for gte >= lte && arr[l] < arr[gte] {
			gte--
		}
		if lte >= gte {
			break
		}
		arr[lte], arr[gte] = arr[gte], arr[lte]
		lte++
		gte--
	}
	arr[l], arr[gte] = arr[gte], arr[l]
	return gte
}

// 三路快排
func sortV4(arr []int, l, r int) {
	if l >= r {
		return
	}

	lt, rt := partitionV4(arr, l, r)
	sortV4(arr, l, lt)
	sortV4(arr, rt, r)
}
func partitionV4(arr []int, l, r int) (int, int) {
	result, _ := rand.Int(rand.Reader, big.NewInt(int64((r-l)+1)))
	seed := int(int64(l) + result.Int64())
	arr[seed], arr[l] = arr[l], arr[seed]

	lt, index, gt := l, l+1, r+1
	// 左边区间：arr[l+1:lt]< v 中间区间 arr[lt+1:gte-1]= v  右边区间：arr[gt:r] >v
	for index < gt {
		if arr[index] < arr[l] {
			lt++
			arr[lt], arr[index] = arr[index], arr[lt]
			index++
		} else if arr[index] == arr[l] {
			index++
		} else {
			gt--
			arr[index], arr[gt] = arr[gt], arr[index]
		}
	}
	// 循环完后，将lt与l位置元素交换，此时区间为：
	// 左边区间：arr[l:lt-1]< v 中间区间 arr[lt:gte-1]==v  右边区间：arr[gt:r] >v
	arr[l], arr[lt] = arr[lt], arr[l]
	return lt - 1, gt
}
