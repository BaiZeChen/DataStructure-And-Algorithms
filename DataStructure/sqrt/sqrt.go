package sqrt

import (
	"errors"
	"math"
)

func NewSqrt[T Ordered](data []T, merge Merge[T]) *Sqrt[T] {
	sq := &Sqrt[T]{}
	return sq.init(data, merge)
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

type Merge[T Ordered] func(front, rear T) T

type Sqrt[T Ordered] struct {
	data, block                     []T // data 代表元数据，block代表每一组的数据（具体根据merge是和、最大值、最小值）
	length, groupValueLen, groupLen int // length代表元数据长度，groupValueLen代表每组元素个数，groupLen代表组数
	merge                           Merge[T]
}

func (s *Sqrt[T]) init(data []T, merge Merge[T]) *Sqrt[T] {
	sq := &Sqrt[T]{
		data:   data,
		merge:  merge,
		length: len(data),
	}
	sq.groupValueLen = int(math.Floor(math.Sqrt(float64(sq.length))))
	if sq.length%sq.groupValueLen != 0 {
		sq.groupLen = sq.length/sq.groupValueLen + 1
	} else {
		sq.groupLen = sq.length / sq.groupValueLen
	}
	sq.block = make([]T, sq.groupLen)
	for i := 0; i < s.length; i++ {
		sq.block[i/s.groupValueLen] = s.merge(sq.block[i/s.groupValueLen], s.data[i])
	}
	return sq
}

func (s *Sqrt[T]) RangeResult(x, y int) (res T, err error) {
	if x < 0 || x >= s.length || y < 0 || y >= s.length || x > y {
		err = errors.New("边界不正确")
		return res, err
	}

	begin, end := x/s.groupValueLen, y/s.groupValueLen
	if begin == end {
		for i := x; i <= y; i++ {
			if i == x {
				res = s.data[i]
			} else {
				res = s.merge(res, s.data[i])
			}
		}
		return
	}
	for i := x; i < (begin+1)*s.groupValueLen; i++ {
		if i == x {
			res = s.data[i]
		} else {
			res = s.merge(res, s.data[i])
		}
	}
	for i := begin + 1; i < end; i++ {
		res = s.merge(res, s.block[i])
	}
	for i := end * s.groupValueLen; i <= y; i++ {
		res = s.merge(res, s.data[i])
	}
	return
}

func (s *Sqrt[T]) Update(index int, val T) bool {
	if index >= s.length {
		return false
	}

	groupNum := index / s.groupValueLen
	s.block[groupNum] = s.block[groupNum] - s.data[index]
	s.block[groupNum] += val
	s.data[index] = val

	return true
}
