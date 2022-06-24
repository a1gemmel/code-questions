package main

import "fmt"

func multiplesOf(n int, lessThan int) []int {
	var s []int

	var i = n
	for i < lessThan {
		s = append(s, i)
		i += n
	}

	return s
}

func main() {

	fives := multiplesOf(5, 1000)
	threes := multiplesOf(3, 1000)

	sum := 0
	for _, i := range fives {
		sum += i
	}
	for _, i := range threes {
		if i%5 != 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
