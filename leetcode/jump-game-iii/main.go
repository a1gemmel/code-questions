package main

func canReachInner(arr []int, pos int, visited map[int]bool) bool {
	if visited[pos] || pos < 0 || pos >= len(arr) {
		return false
	}
	if arr[pos] == 0 {
		return true
	}
	visited[pos] = true
	return canReachInner(arr, pos+arr[pos], visited) || canReachInner(arr, pos-arr[pos], visited)
}

func canReach(arr []int, start int) bool {
	return canReachInner(arr, start, map[int]bool{})
}
