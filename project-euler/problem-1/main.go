package main

import (
	"fmt"

	"github.com/a1gemmel/project-euler/golang/lib"
)

func main() {

	fives := lib.MultiplesOf(5, 1000)
	threes := lib.MultiplesOf(3, 1000)

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
