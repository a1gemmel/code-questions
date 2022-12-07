package main

func median(arr []int) float64 {
	if len(arr)%2 == 0 {
		return float64(arr[len(arr)/2]+arr[len(arr)/2-1]) / 2.0
	}
	return float64(arr[len(arr)/2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge(nums1 []int, nums2 []int) []int {
	merged := make([]int, 0, len(nums1)+len(nums2))
	for len(nums1) > 0 || len(nums2) > 0 {
		if len(nums1) == 0 {
			return append(merged, nums2...)
		}
		if len(nums2) == 0 {
			return append(merged, nums1...)
		}
		if nums1[0] < nums2[0] {
			merged = append(merged, nums1[0])
			nums1 = nums1[1:]
		} else {
			merged = append(merged, nums2[0])
			nums2 = nums2[1:]
		}
	}
	return merged
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for len(nums1) > 2 && len(nums2) > 2 {
		minLen := min(len(nums1), len(nums2))
		var toCut int
		if minLen%2 == 0 {
			toCut = (minLen - 2) / 2
		} else {
			toCut = (minLen - 1) / 2
		}
		if median(nums1) < median(nums2) {
			nums1 = nums1[toCut:]
			nums2 = nums2[0 : len(nums2)-toCut]
		} else {
			nums2 = nums2[toCut:]
			nums1 = nums1[0 : len(nums1)-toCut]
		}
	}

	merged := merge(nums1, nums2)
	return median(merged)
}

// func main() {
// 	fmt.Println(findMedianSortedArrays([]int{1, 13, 14, 16, 18, 21, 34}, []int{3, 4, 5, 67, 89, 91}))
// }
