package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsPalindrome(t *testing.T) {
	tests := map[int]bool{
		0:       true,
		-1:      false,
		-12315:  false,
		1:       true,
		10:      false,
		1000:    false,
		11:      true,
		99:      true,
		100:     false,
		101:     true,
		666:     true,
		1000021: false,
	}

	for n, res := range tests {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			fmt.Println(n)
			assert.Equal(t, res, isPalindrome(n))
		})
	}
}
