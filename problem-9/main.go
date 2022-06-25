package main

import (
	"fmt"
)

func main() {

	for a := 1; a < 1000; a++ {
		for b := 1; a+b < 1000; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Printf("%d + %d + %d = %d\n", a, b, c, a+b+c)
				fmt.Printf("%d * %d * %d = %d\n", a, b, c, a*b*c)
			}
		}
	}

}
