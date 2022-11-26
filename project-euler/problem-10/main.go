package main

import (
	"fmt"

	"github.com/a1gemmel/project-euler/golang/lib"
)

func main() {
	const max = 2000000 - 1
	fmt.Println(lib.SumInts(lib.PrimeSieve(max)))
}
