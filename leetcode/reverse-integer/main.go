package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	isNegative := x < 0
	s := fmt.Sprintf("%d", x)
	if isNegative {
		s = s[1:]
	}
	reversedBytes := make([]byte, len(s))
	for i, b := range []byte(s) {
		reversedBytes[len(s)-1-i] = b
	}
	reversed := string(reversedBytes)
	if isNegative {
		reversed = "-" + reversed
	}
	integer, err := strconv.ParseInt(reversed, 10, 32)
	if err != nil {
		return 0
	}
	return int(integer)
}
