package main

import (
	"fmt"
	"math"
)

const n = 600851475143

func previousPrime(i int) int {
	if i%2 == 0 {
		i++
	}

	for {
		i -= 2
		if isPrime(i) {
			return i
		}
	}
}

func isPrime(i int) bool {
	fmt.Printf("Testing %d for primeness\n", i)
	sqrt := int(math.Floor(math.Sqrt(float64(i))))

	for d := 3; d <= sqrt; d += 2 {
		if i%d == 0 {
			return false
		}
	}

	return true
}

func main() {

	p := int(math.Floor(math.Sqrt(float64(n))))

	if p%2 == 0 {
		p++
	}

	progress := 1

	for {
		if n%p == 0 && isPrime(p) {
			fmt.Println(p)
			return
		}
		p -= 2
		progress++
		if progress%100000 == 0 {
			fmt.Printf("Checked %d numbers, latest %d\n", progress, p)
		}

	}

	// for {
	// 	if n%p == 0 {
	// 		fmt.Println(p)
	// 		return
	// 	}
	// 	p = previousPrime(p)
	// 	if progress%10 == 0 {
	// 		fmt.Printf("Tested %d primes (latest %d)\n", progress, p)
	// 	}
	// }
}
