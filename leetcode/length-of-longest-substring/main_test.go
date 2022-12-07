package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestSubstringLength(t *testing.T) {
	tests := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
		"a":        1,
		"ewok":     4,
	}

	for s, length := range tests {
		t.Run(s, func(t *testing.T) {
			assert.Equal(t, length, lengthOfLongestSubstring(s))
		})
	}
}
