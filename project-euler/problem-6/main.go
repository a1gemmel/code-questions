package main

import (
	"fmt"

	"github.com/a1gemmel/project-euler/golang/lib"
)

func square(i int) int {
	return i * i
}

func main() {
	n := lib.GenerateN(100)
	fmt.Println(square(lib.SumInts(n)) - lib.SumInts(lib.SquareAll(n)))
}
