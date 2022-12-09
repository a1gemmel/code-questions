package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reverse(t *testing.T) {
	tests := map[int]int{
		12:         21,
		12345:      54321,
		2147483647: 0,
	}

	for in, expected := range tests {
		assert.Equal(t, expected, reverse(in))
	}
}
