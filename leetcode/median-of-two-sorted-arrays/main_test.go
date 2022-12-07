package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Median(t *testing.T) {
	tests := map[float64][]int{
		2:   {1, 2, 3},
		1:   {1},
		2.5: {1, 2, 3, 4},
	}

	for expected, arr := range tests {
		assert.Equal(t, expected, median(arr))
	}
}

func Test_MedianOf2(t *testing.T) {
	tests := []struct {
		expected float64
		arr1     []int
		arr2     []int
	}{
		{
			expected: 2.5,
			arr1:     []int{1, 2},
			arr2:     []int{3, 4},
		},
		{
			expected: 7,
			arr1:     []int{1, 2, 4, 6, 7, 8, 9},
			arr2:     []int{3, 4, 5, 7, 10, 12, 14, 15},
		},
		{
			expected: 2,
			arr1:     []int{1, 3},
			arr2:     []int{2},
		},
		{
			expected: 6,
			arr1:     []int{1, 2, 3, 4, 5},
			arr2:     []int{6, 7, 8, 9, 10, 11},
		},
		{
			expected: 16,
			arr1:     []int{1, 13, 14, 16, 18, 21, 34},
			arr2:     []int{3, 4, 5, 67, 89, 91},
		},
		{
			expected: 6,
			arr1:     []int{6},
			arr2:     []int{},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, findMedianSortedArrays(test.arr1, test.arr2))
		assert.Equal(t, test.expected, findMedianSortedArrays(test.arr2, test.arr1))
	}
}
