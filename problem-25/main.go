package main

import (
	"fmt"

	"github.com/a1gemmel/project-euler/lib"
)

func main() {

	a := lib.NewBigInt(1)
	b := lib.NewBigInt(1)
	index := 2

	for b.Length() < 1000 {
		f := a.Add(b)
		a = b
		b = f
		index++

		if b.Length() > a.Length() {
			fmt.Printf("Index: %d (Length %d)\n", index, b.Length())
		}

	}

	fmt.Println(index)
}
