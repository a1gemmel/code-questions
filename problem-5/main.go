package main

import "fmt"

func divisibleBy(n int, arr []int) bool {
	for _, el := range arr {
		if n%el != 0 {
			return false
		}
	}
	return true
}

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

func main() {
	n := 20

	for {
		if divisibleBy(n, nums) {
			fmt.Println(n)
			return
		}
		n += 2
	}

}
