package lib

func SumInts(arr []int) int {
	sum := 0
	for _, el := range arr {
		sum += el
	}
	return sum
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
