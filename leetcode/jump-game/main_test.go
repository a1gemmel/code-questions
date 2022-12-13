package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanJump(t *testing.T) {
	tests := map[string]struct {
		arr    []int
		result bool
	}{
		"case-1": {
			arr:    []int{2, 3, 1, 1, 4},
			result: true,
		},
		"case-2": {
			arr:    []int{3, 2, 1, 0, 4},
			result: false,
		},
		"degen": {
			arr:    []int{0},
			result: true,
		},
		"simple-impossible": {
			arr:    []int{0, 1},
			result: false,
		},
		"simple-possible": {
			arr:    []int{1, 0},
			result: true,
		},
		"backtrack": {
			arr:    []int{4, 6, 1, 1, 0, 0, 0, 1},
			result: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.result, canJump(test.arr))
		})
	}
}
