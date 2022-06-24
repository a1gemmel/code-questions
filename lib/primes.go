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
