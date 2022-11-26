package main

import (
	"fmt"

	"github.com/a1gemmel/project-euler/golang/lib"
)

func properDivisorSum(n int) int {
	if n == 1 {
		return 0
	}
	return lib.SumInts(lib.AllDivisors(n)[2:])
}

func main() {
	isAbundant := map[int]bool{}
	upperBound := 28123 + 1

	for i := 1; i < upperBound; i++ {
		isAbundant[i] = properDivisorSum(i) > i
	}

	sum := 0

	for i := 1; i < upperBound; i++ {
		expressable := false
		for x := 1; x < i; x++ {
			y := i - x
			if isAbundant[x] && isAbundant[y] {
				expressable = true
				break
			}
		}
		if !expressable {
			sum += i
		}
	}

	fmt.Println(sum)

}
