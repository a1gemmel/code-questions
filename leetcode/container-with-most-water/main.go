package main

import "sort"

/*
Idea:

Compute maps from each height[i] to first and last index of a number at least height[i]
loop through height, take max of height[i] * (first(height[i]) - last(height[i]))

*/

func maxArea(height []int) int {

	heightsSorted := make([]int, len(height))
	for i, h := range height {
		heightsSorted[i] = h
	}
	sort.Ints(heightsSorted)

	firstAtLeastN := map[int]int{}
	lastAtLeastN := map[int]int{}

	scan := 0
	for _, h := range heightsSorted {
		for {
			if height[scan] >= h {
				firstAtLeastN[h] = scan
				break
			} else {
				scan++
			}
		}
	}

	scan = len(height) - 1
	for _, h := range heightsSorted {
		for {
			if height[scan] >= h {
				lastAtLeastN[h] = scan
				break
			} else {
				scan--
			}
		}
	}

	max := 0

	for _, h := range height {
		volume := h * (lastAtLeastN[h] - firstAtLeastN[h])
		if volume > max {
			max = volume
		}
	}

	return max
}
