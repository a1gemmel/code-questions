package main

import (
	"fmt"
	"math"
)

const delta = 0.0001
const cap = 100

func exponentForm(x int) (base, power int) {
	for i := float64(100); i >= 1; i-- {
		//for i := float64(2); i <= cap; i++ {
		nRoot := math.Pow(float64(x), 1/i)
		if math.Abs(nRoot-math.Round(nRoot)) <= delta {
			return int(math.Round(nRoot)), int(i)
		}
	}
	return x, 1
}

func main() {
	s := map[string]bool{}

	for a := 2; a <= cap; a++ {
		aBase, aPower := exponentForm(a)
		for b := 2; b <= cap; b++ {
			s[fmt.Sprintf("%d^%d", aBase, aPower*b)] = true
		}
	}

	fmt.Println(len(s))
}
