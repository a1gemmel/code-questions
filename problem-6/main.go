package main

import "fmt"

func generateN(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return arr
}

func squareN(arr []int) []int {
	out := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		out[i] = arr[i] * arr[i]
	}
	//fmt.Println(out)
	return out
}

func sum(arr []int) int {
	sum := 0
	for _, el := range arr {
		sum += el
	}
	return sum
}

func square(i int) int {
	return i * i
}

func main() {

	n := generateN(100)

	fmt.Println(square(sum(n)) - sum(squareN(n)))
}
