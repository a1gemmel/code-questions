package lib

func SumInts(arr []int) int {
	sum := 0
	for _, el := range arr {
		sum += el
	}
	return sum
}

func SumInt64s(arr []int64) int64 {
	sum := int64(0)
	for _, el := range arr {
		sum += el
	}
	return sum
}

func MultiplyInts(arr []int) int {
	product := 1
	for _, el := range arr {
		product *= el
	}
	return product
}

func MultiplesOf(n int, lessThan int) []int {
	var s []int

	var i = n
	for i < lessThan {
		s = append(s, i)
		i += n
	}

	return s
}

func DivisibleByAll(n int, arr []int) bool {
	for _, el := range arr {
		if n%el != 0 {
			return false
		}
	}
	return true
}

func SquareAll(arr []int) []int {
	out := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		out[i] = arr[i] * arr[i]
	}
	return out
}

func GenerateN(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return arr
}

func Max(arr []int) int {
	max := arr[0]

	for _, el := range arr {
		if el > max {
			max = el
		}
	}
	return max
}

func AllDivisors(n int) []int {
	divisors := []int{n}
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
