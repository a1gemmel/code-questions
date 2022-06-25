package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PrimeSieve(t *testing.T) {
	assert.Equal(t, []int{2, 3, 5, 7}, PrimeSieve(10))
	assert.Equal(t, []int{2, 3, 5, 7, 11, 13, 17}, PrimeSieve(17))
	assert.Equal(t, []int{2}, PrimeSieve(2))
	assert.Equal(t, []int{}, PrimeSieve(1))
}
