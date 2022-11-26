package main

import (
	"fmt"

	"github.com/a1gemmel/project-euler/golang/lib"
)

func prepend(s string, strings []string) []string {
	for i := 0; i < len(strings); i++ {
		strings[i] = s + strings[i]
	}
	return strings
}

func allPermutations(chars []string) []string {
	if len(chars) == 1 {
		return chars
	}

	var permutations []string

	for i, c := range chars {
		permutations = append(permutations, prepend(c, allPermutations(lib.RemoveAtIndex(i, chars)))...)
	}
	return permutations
}

func main() {

	set := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	p := allPermutations(set)

	fmt.Println(p[999999])

}
