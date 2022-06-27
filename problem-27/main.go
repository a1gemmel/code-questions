package main

import (
	"fmt"
	"time"

	"github.com/a1gemmel/project-euler/lib"
)

func formula(n, a, b int) int {
	return n*n + a*n + b
}

func main() {

	maxPrimes := 0
	maxA := 0
	maxB := 0

	t0 := time.Now()

	for a := -999; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			n := 0
			for {
				p := formula(n, a, b)
				if p < 2 || !lib.IsPrime(p) {
					break
				}
				n++
			}

			if n > maxPrimes {
				maxPrimes = n
				maxA = a
				maxB = b
			}
		}
	}

	fmt.Printf("a=%d,b=%d produces %d primes\n", maxA, maxB, maxPrimes)
	fmt.Println("Solution in ", time.Now().Sub(t0))

}
