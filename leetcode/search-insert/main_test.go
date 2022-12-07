package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SearchInsert(t *testing.T) {
	tests := map[string]struct {
		expected int
		target   int
		arr      []int
	}{
		"exists-in-middle": {
			expected: 2,
			target:   5,
			arr:      []int{0, 1, 5, 7, 8, 9},
		},
		"exists-in-middle-2": {
			expected: 2,
			target:   5,
			arr:      []int{1, 3, 5, 6},
		},
		"missing-in-first-half": {
			expected: 1,
			target:   2,
			arr:      []int{1, 3, 5, 6},
		},
		"missing-in-second-half": {
			expected: 3,
			target:   8,
			arr:      []int{3, 5, 7, 9, 10},
		},
		"missing-at-end": {
			expected: 4,
			target:   7,
			arr:      []int{1, 3, 5, 6},
		},
		"missing-at-beginning": {
			expected: 0,
			target:   0,
			arr:      []int{1, 3, 5, 6},
		},
		"exists-at-beginning": {
			expected: 0,
			target:   1,
			arr:      []int{1, 3, 5, 6},
		},
		"exists-at-end": {
			expected: 3,
			target:   6,
			arr:      []int{1, 3, 5, 6},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			fmt.Println(name)
			assert.Equal(t, test.expected, searchInsert(test.arr, test.target))
		})
	}
}
