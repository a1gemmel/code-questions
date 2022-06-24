package main

import (
	"fmt"

	"github.com/a1gemmel/project-euler/lib"
)

func main() {

	p := 2

	for i := 0; i < 10000; i++ {
		p = lib.NextPrime(p)
	}

	fmt.Println(p)
}
