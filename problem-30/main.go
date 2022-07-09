package main

import "fmt"

func pow(n int64, p int) int64 {
	if p == 1 {
		return n
	}
	return n * pow(n, p-1)
}

func sumOfDigitsToFifthPower(n int64) int64 {
	var sum int64 = 0
	for n != 0 {
		sum += pow(n%10, 5)
		n /= 10
	}
	return sum
}

func main() {

	// The upper bound of number to test is when n * 9^5 < 10*n
	// n = 5.51
	// So we will test up to 6 digits

	var sum int64 = 0
	var i int64
	for i = 10; i < 999999; i++ {
		if i == sumOfDigitsToFifthPower(i) {
			sum += i
			fmt.Println(i)
		}
	}

	fmt.Println(sum)
}
