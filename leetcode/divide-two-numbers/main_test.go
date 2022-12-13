package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Divide(t *testing.T) {
	assert.Equal(t, 2, divide(6, 3))
	assert.Equal(t, 1, divide(6, 5))
	assert.Equal(t, -1, divide(6, -5))
	assert.Equal(t, -2, divide(-6, 3))
	assert.Equal(t, -2, divide(6, -3))
	assert.Equal(t, 2, divide(-6, -3))
	assert.Equal(t, 4, divide(25, 6))
	assert.Equal(t, 1231234, divide(1231234, 1))
	assert.Equal(t, 2147483647, divide(1231234123123, 1))
	assert.Equal(t, 2147483647, divide(-2147483648, -1))
}
