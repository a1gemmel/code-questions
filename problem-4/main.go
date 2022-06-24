package main

import (
	"fmt"
	"sort"
)

func reverseStr(s string) string {
	b := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	return string(b)
}

func main() {

	var palindromes []int

	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			product := i * j

			productStr := fmt.Sprintf("%d", product)
			if reverseStr(productStr) == productStr {
				fmt.Println(i, j, product)
				palindromes = append(palindromes, product)
			}
		}
	}

	sort.IntSlice(palindromes).Sort()

	fmt.Println(palindromes)

}
