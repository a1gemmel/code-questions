package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var rooms, teams int
	fmt.Scanf("%d\n%d", &rooms, &teams)
	for rooms > 0 {
		toAdd := int(math.Round(float64(teams) / float64(rooms)))
		rooms -= 1
		teams -= toAdd
		fmt.Println(strings.Repeat("*", toAdd))
	}
}
