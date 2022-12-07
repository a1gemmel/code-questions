package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if target < nums[0] {
		return 0
	}

	start := 0
	end := len(nums) - 1

	for {
		middle := (start + end) / 2
		middleVal := nums[middle]
		if middleVal == target {
			return middle
		}
		if start == end {
			if middleVal < target {
				return start + 1
			} else {
				return start
			}
		}
		if middleVal < target {
			start = middle + 1
		} else {
			end = middle
		}
	}
}

func main() {
	fmt.Println(searchInsert([]int{3, 5, 7, 9, 10}, 8))
}

/*

1,3,5,6 (1)
(0,3) -> 1 (3)
(0,0) -> 0 (1)

*/
