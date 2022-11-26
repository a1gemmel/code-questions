package lib

// TODO this is a poor man's prime factorization cancellation
func Choose(n, c int) int64 {
	numerators := GenerateN(n)
	denominators := append(GenerateN(c), GenerateN(n-c)...)

	for i, _ := range numerators {
		for j, _ := range denominators {
			if numerators[i]%denominators[j] == 0 {
				numerators[i] = numerators[i] / denominators[j]
				denominators[j] = 1
			}
		}
	}
	return int64(MultiplyInts(numerators)) / int64(MultiplyInts(denominators))
}

func Factorial(n int) int64 {
	if n == 1 {
		return int64(n)
	}
	return int64(n) * Factorial(n-1)
}
