package main

import (
	"fmt"
	"sort"

	"github.com/a1gemmel/project-euler/golang/lib"
)

func main() {

	var palindromes []int

	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			product := i * j

			productStr := fmt.Sprintf("%d", product)
			if lib.ReverseStr(productStr) == productStr {
				fmt.Println(i, j, product)
				palindromes = append(palindromes, product)
			}
		}
	}

	sort.IntSlice(palindromes).Sort()

	fmt.Println(palindromes)

}
