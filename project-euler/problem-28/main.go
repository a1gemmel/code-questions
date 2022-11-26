package main

import "fmt"

func sumOfCorners(sideLength int) int {
	n := sideLength
	topRightDownOne := (n-2)*(n-2) + 1
	bottomRight := topRightDownOne + (n - 2)
	bottomLeft := bottomRight + (n - 1)
	topLeft := bottomLeft + (n - 1)
	topRight := topLeft + (n - 1)
	fmt.Println(bottomRight, bottomLeft, topLeft, topRight)
	return bottomLeft + bottomRight + topLeft + topRight
}

func main() {

	sum := 1
	n := 1001

	for i := 3; i <= n; i += 2 {
		sum += sumOfCorners(i)
	}
	fmt.Println(sum)
}
