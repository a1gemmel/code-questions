package main

import (
	"fmt"

	"github.com/a1gemmel/project-euler/lib"
)

func main() {

	limit := 10000

	divisorSum := make([]int, limit+1)

	divisorSum[1] = 1
	for i := 2; i <= limit; i++ {
		allDivisors := lib.AllDivisors(i)[1:]
		divisorSum[i] = lib.SumInts(allDivisors)
		fmt.Printf("%d (%d) - %v\n", i, divisorSum[i], allDivisors)
	}

	amicableSum := 0

	for a := 1; a <= limit; a++ {
		b := divisorSum[a]

		if b != a && b <= limit && divisorSum[b] == a {
			amicableSum += a
		}
	}

	fmt.Println(amicableSum)
}
