package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Factorial(t *testing.T) {
	assert.Equal(t, int64(1), Factorial(1))
	assert.Equal(t, int64(2), Factorial(2))
	assert.Equal(t, int64(6), Factorial(3))
}

func Test_Choose(t *testing.T) {
	assert.Equal(t, int64(2), Choose(2, 1))
	assert.Equal(t, int64(184756), Choose(20, 10))
}
