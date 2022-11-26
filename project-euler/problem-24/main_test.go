package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Prepend(t *testing.T) {
	assert.Equal(t, []string{"012", "021"}, prepend("0", []string{"12", "21"}))
}
