package main

import (
	"fmt"
	"sort"
)

func median(arr []int) float64 {
	if len(arr)%2 == 0 {
		return float64(arr[len(arr)/2]+arr[len(arr)/2-1]) / 2.0
	}
	return float64(arr[len(arr)/2])
}

// returns the new arrays and number of elements removed
func removeNLessThanOrEqualToX(arr1, arr2 []int, n int, x float64) ([]int, []int, int) {
	removed := 0
	for n > 0 {
		if len(arr1) == 0 || len(arr2) == 0 {
			return arr1, arr2, removed
		}
		if arr1[0] < arr2[0] {
			arr1 = arr1[1:]
		} else {
			arr2 = arr2[1:]
		}
		removed += 1
	}
	return arr1, arr2, removed
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func beforeMedian(length int) int {
	if length%2 == 0 {
		return length / 2
	}
	return length / 2
}

func take2(a []int) []int {
	switch len(a) {
	case 0:
		return []int{}
	case 1:
		return a[:1]
	default:
		return a[:2]
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	toRemove := totalLen/2 - 1

	for len(nums1) > 0 && len(nums2) > 0 {

		median1 := median(nums1)
		median2 := median(nums2)
		removed := 0
		if median1 < median2 {
			removed = min(toRemove, beforeMedian(len(nums1)))
			nums1 = nums1[removed:]
			toRemove -= removed
		} else {
			removed = min(toRemove, beforeMedian(len(nums2)))
			nums2 = nums2[removed:]
			toRemove -= removed
		}
		fmt.Println("To Remove: ", toRemove, nums1, nums2)

		if toRemove == 0 || removed == 0 {
			break
		}
	}

	if totalLen%2 == 1 {
		if len(nums1) == 0 {
			return float64(nums2[0])
		}
		if len(nums2) == 0 {
			return float64(nums1[0])
		}
		return float64(min(nums1[0], nums2[0]))
	}

	concat := append(append([]int{}, take2(nums1)...), take2(nums2)...)
	sort.Ints(concat)
	return float64(concat[0]+concat[1]) / 2

}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
