package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Multiply(t *testing.T) {
	assert.Equal(t, NewBigInt(52), NewBigInt(26).Multiply(2))
	assert.Equal(t, NewBigInt(581152), NewBigInt(4576).Multiply(127))
}
