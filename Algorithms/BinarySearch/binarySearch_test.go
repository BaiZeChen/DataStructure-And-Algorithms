package BinarySearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchV1(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		target int
		match  int
	}{
		{
			"测试1",
			[]int{1, 3, 5, 6, 7, 10, 21},
			7,
			4,
		},
		{
			"测试2",
			[]int{2, 4, 7, 13, 88, 23, 58},
			23,
			5,
		},
		{
			"测试3",
			[]int{2, 4, 7, 13, 88, 23, 58},
			110,
			-1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			match := SearchV1(tc.arr, tc.target, 0, len(tc.arr)-1)
			assert.Equal(t, tc.match, match)
		})
	}

}

func TestSearchV2(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		target int
		match  int
	}{
		{
			"测试1",
			[]int{1, 3, 5, 6, 7, 10, 21},
			7,
			4,
		},
		{
			"测试2",
			[]int{2, 4, 7, 13, 88, 23, 58},
			23,
			5,
		},
		{
			"测试3",
			[]int{2, 4, 7, 13, 88, 23, 58},
			110,
			-1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			match := SearchV2(tc.arr, tc.target)
			assert.Equal(t, tc.match, match)
		})
	}
}

func TestUpper(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		target int
		match  int
	}{
		{
			name:   "元素在数组内1",
			arr:    []int{1, 3, 5, 6, 7, 10, 21},
			target: 8,
			match:  5,
		},
		{
			name:   "元素不在数组内",
			arr:    []int{1, 3, 5, 6, 7, 10, 21},
			target: 100,
			match:  -1,
		},
		{
			name:   "元素在数组内2",
			arr:    []int{1, 3, 5, 6, 7, 10, 21},
			target: 13,
			match:  6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			match := Upper(tc.arr, tc.target)
			assert.Equal(t, tc.match, match)
		})
	}
}

func TestLower(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		target int
		match  int
	}{
		{
			name:   "元素在数组内1",
			arr:    []int{1, 3, 5, 6, 7, 10, 21},
			target: 8,
			match:  4,
		},
		{
			name:   "元素不在数组内",
			arr:    []int{1, 3, 5, 6, 7, 10, 21},
			target: -3,
			match:  -1,
		},
		{
			name:   "元素在数组内2",
			arr:    []int{1, 3, 5, 6, 7, 10, 21},
			target: 13,
			match:  5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			match := Lower(tc.arr, tc.target)
			assert.Equal(t, tc.match, match)
		})
	}
}
