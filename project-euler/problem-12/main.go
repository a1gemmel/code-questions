package main

import (
	"fmt"
	"math"
)

func countDivisors(n int) int {
	count := 0
	sqrtN := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			count += 2
		}
	}
	return count
}

func main() {
	i := 1
	triangleNumber := 0

	for {
		triangleNumber += i
		i++

		divisors := countDivisors(triangleNumber)
		//fmt.Println(triangleNumber, divisors)
		if divisors > 500 {
			fmt.Println(triangleNumber)
			return
		}
	}
}
