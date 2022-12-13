package main

/*
	Idea - always greedily jump as far as possible until reaching a zero,
	then back up until finding a number that can jump past that zero.
	If the beginning is reached, it's not possible
*/

func canJump(nums []int) bool {
	furthest := 0
	current := 0
	for {
		if current >= len(nums)-1 {
			return true
		}
		if current < 0 {
			return false
		}
		if nums[current] > 0 && (furthest == current || current+nums[current] > furthest) {
			current += nums[current]
			furthest = current
		} else {
			current--
		}
	}
}
