package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RomanToInteger(t *testing.T) {
	cases := map[string]int{
		"IV":      4,
		"III":     3,
		"XII":     12,
		"XXVII":   27,
		"MCMXCIV": 1994,
		"LVIII":   58,
	}

	for roman, integer := range cases {
		assert.Equal(t, integer, romanToInt(roman))
	}
}
