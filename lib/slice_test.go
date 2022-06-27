package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RemoveAtIndex(t *testing.T) {
	assert.Equal(t, []string{"foo", "bar", "bazz"}, RemoveAtIndex(0, []string{"fizz", "foo", "bar", "bazz"}))
	assert.Equal(t, []string{"fizz", "foo", "bar"}, RemoveAtIndex(3, []string{"fizz", "foo", "bar", "bazz"}))
}
