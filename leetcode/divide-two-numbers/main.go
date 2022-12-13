package main

func abs(n int) int {
	if n > 0 {
		return n
	}
	return 0 - n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func divide(dividend int, divisor int) int {
	var signOp func(int, int) int
	isPositiveResult := (dividend < 0) == (divisor < 0)
	if isPositiveResult {
		signOp = func(a, b int) int { return a - b }
	} else {
		signOp = func(a, b int) int { return a + b }
	}

	divisors := []int{divisor}
	divisorCounts := []int{1}

	res := 0
	for {
		if len(divisors) < 1 {
			break
		}
		d := divisors[len(divisors)-1]
		count := divisorCounts[len(divisorCounts)-1]

		if abs(dividend) >= abs(d) {
			dividend = signOp(dividend, d)
			res += count
			divisors = append(divisors, d+d)
			divisorCounts = append(divisorCounts, count+count)
		} else {
			divisorCounts = divisorCounts[0 : len(divisorCounts)-1]
			divisors = divisors[0 : len(divisors)-1]
		}

	}
	if !isPositiveResult {
		return max((0 - res), -2147483648)
	}
	return min(res, 2147483647)
}
