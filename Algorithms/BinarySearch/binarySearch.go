package BinarySearch

// SearchV1 递归版本
func SearchV1(arr []int, target int, l, r int) int {
	if l > r {
		return -1
	}

	middle := (l + r) / 2
	if arr[middle] < target {
		return SearchV1(arr, target, middle+1, r)
	} else if arr[middle] > target {
		return SearchV1(arr, target, l, middle-1)
	} else {
		return middle
	}
}

// SearchV2 迭代版本
func SearchV2(arr []int, target int) int {
	l, r := 0, len(arr)-1

	for l <= r {
		middle := (l + r) / 2
		if arr[middle] < target {
			l = middle + 1
		} else if arr[middle] > target {
			r = middle - 1
		} else {
			return middle
		}
	}
	return -1
}

// Upper 查找大于target的最小值
// 范围[l,r)，如果当前元素比数组最大值都大，则返回-1
func Upper(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		middle := (l + r) / 2
		if arr[middle] <= target {
			l = middle + 1
		} else {
			r = middle
		}
	}
	if r == len(arr) {
		return -1
	}
	return r
}

// Lower 查找小于target的最大值
// 范围(l,r]，如果当前元素比数组最小值都小，则返回-1
func Lower(arr []int, target int) int {
	l, r := -1, len(arr)-1
	for l < r {
		middle := (l + r) / 2
		if arr[middle] < target {
			l = middle
			// 相邻时容易出现死循环，所以这里判断一下
			if r-l == 1 {
				r--
			}
		} else {
			r = middle - 1
		}
	}
	return r
}
