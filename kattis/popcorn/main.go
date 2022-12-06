package main

import (
	"fmt"
)

func main() {
	var participants int
	fmt.Scanf("%d", &participants)
	var totalBags int64 = 4
	groupSize := participants / 4
	if groupSize > 1 {
		totalBags += 4 * choose2(groupSize)
	}
	fmt.Println(totalBags)
}

func choose2(n int) int64 {
	if n%2 == 0 {
		return int64(n) / 2 * int64(n-1)
	}
	return int64(n-1) / 2 * int64(n)
}
