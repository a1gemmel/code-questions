package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Multiply(t *testing.T) {
	assert.Equal(t, NewBigInt(52), NewBigInt(26).Multiply(2))
	assert.Equal(t, NewBigInt(581152), NewBigInt(4576).Multiply(127))
}

func Test_Add(t *testing.T) {
	assert.Equal(t, NewBigInt(121), NewBigInt(50).Add(NewBigInt(71)))
	assert.Equal(t, NewBigInt(12), NewBigInt(0).Add(NewBigInt(12)))
	assert.Equal(t, NewBigInt(200), NewBigInt(100).Add(NewBigInt(100)))
}
