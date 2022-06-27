package main

import (
	"fmt"
	"sort"
)

func nameScore(name string) int {
	score := 0
	for _, r := range name {
		score += int(r - 'A' + 1)
	}
	return score
}

func main() {
	sort.StringSlice(names).Sort()

	totalScore := 0

	for i, name := range names {
		totalScore += (i + 1) * nameScore(name)
	}

	fmt.Println(totalScore)

}
