package lib

import "math"

func IsPrime(i int) bool {
	sqrt := int(math.Floor(math.Sqrt(float64(i))))

	for d := 3; d <= sqrt; d += 2 {
		if i%d == 0 {
			return false
		}
	}

	return true
}

func PreviousPrime(i int) int {
	if i%2 == 0 {
		i++
	}

	for {
		i -= 2
		if IsPrime(i) {
			return i
		}
	}
}

func NextPrime(i int) int {
	if i%2 == 0 {
		i--
	}

	for {
		i += 2
		if IsPrime(i) {
			return i
		}
	}
}

func PrimeSieve(max int) []int {
	isntPrime := make([]bool, max+1)

	for i := 2; i*i < max; i++ {
		for j := i * 2; j <= max; j += i {
			isntPrime[j] = true
		}
	}

	primes := []int{}

	for ind, isnt := range isntPrime {
		if ind > 1 && !isnt {
			primes = append(primes, ind)
		}
	}

	return primes
}
