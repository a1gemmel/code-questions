package main

import (
	"fmt"

	"github.com/a1gemmel/project-euler/golang/lib"
)

func doubleLongNum(n []int) []int {
	length := len(n)
	for i := length - 1; i >= 0; i-- {
		n[i] *= 2
		if n[i] > 9 {
			n[i] %= 10
			carry := 1
			for j := i + 1; carry != 0; j++ {
				if j >= len(n) {
					n = append(n, 0)
				}
				n[j] += carry
				if n[j] > 9 {
					n[j] %= 10
				} else {
					carry = 0
				}
			}
		}
	}
	return n
}

func main() {

	longNum := []int{1}

	for i := 1; i <= 1000; i++ {
		longNum = doubleLongNum(longNum)
	}

	fmt.Println(longNum)
	fmt.Println(lib.SumInts(longNum))

}
