package main

import (
	"fmt"

	"github.com/a1gemmel/project-euler/golang/lib"
)

func main() {

	/*
		The path can be encoded as a binary number where 0 is go right and 1 is go down.
		Then the problem is equivalent to how many binary numbers of length x+y have
		x 1's and y 0's. This is the nCr combinations formula.
	*/

	gridHeight := 20
	gridWidth := 20

	fmt.Println(lib.Choose(gridHeight+gridWidth, gridHeight))
}
