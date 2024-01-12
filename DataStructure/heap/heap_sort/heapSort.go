package heap_sort

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func HeapSort[T Ordered](arr []T) []T {
	length := len(arr)
	// 现将数组变成一个最大堆
	for i := (length - 2) / 2; i >= 0; i-- {
		// siftDown
		tmp := i
		for {
			childIndex := 2*tmp + 1
			if childIndex >= length {
				break
			}
			if childIndex+1 < length && arr[childIndex] < arr[childIndex+1] {
				childIndex++
			}
			if arr[tmp] > arr[childIndex] {
				break
			}
			arr[tmp], arr[childIndex] = arr[childIndex], arr[tmp]
			tmp = childIndex
		}
	}
	// 然后开始排序
	for i := length - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		tmp := 0
		for {
			childIndex := 2*tmp + 1
			if childIndex > i-1 {
				break
			}
			if childIndex+1 <= i-1 && arr[childIndex] < arr[childIndex+1] {
				childIndex++
			}
			if arr[tmp] > arr[childIndex] {
				break
			}
			arr[tmp], arr[childIndex] = arr[childIndex], arr[tmp]
			tmp = childIndex
		}
	}
	return arr
}
