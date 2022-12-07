package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestCommonPrefix(t *testing.T) {
	tests := map[string][]string{
		"flo": {"flower", "float", "floss"},
		"add": {"adder", "add", "added"},
		"":    {"dog", "water", "cat"},
	}

	for prefix, set := range tests {
		assert.Equal(t, prefix, longestCommonPrefix(set))
	}

}
