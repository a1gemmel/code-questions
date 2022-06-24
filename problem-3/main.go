package main

import (
	"fmt"
	"math"

	"github.com/a1gemmel/project-euler/lib"
)

const n = 600851475143

func main() {

	p := int(math.Floor(math.Sqrt(float64(n))))

	if p%2 == 0 {
		p++
	}

	progress := 1

	for {
		if n%p == 0 && lib.IsPrime(p) {
			fmt.Println(p)
			return
		}
		p -= 2
		progress++
		if progress%100000 == 0 {
			fmt.Printf("Checked %d numbers, latest %d\n", progress, p)
		}

	}

}
