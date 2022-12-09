package main

import (
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	magnitude := int(math.Pow(10, math.Ceil(math.Log10(float64(x)+0.01))-1))
	for {
		if magnitude < 10 {
			break
		}
		if x%10 != x/magnitude {
			return false
		}
		x = (x % magnitude) / 10
		magnitude /= 100
	}
	return true
}
