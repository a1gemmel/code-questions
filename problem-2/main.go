package main

import "fmt"

func nextFib(prev [2]int) (int, [2]int) {
	n := prev[0] + prev[1]
	return n, [2]int{prev[1], n}
}

func main() {

	var rollingFib = [2]int{1, 2}

	sum := 2
	f := 0

	for f < 4000000 {
		if f%2 == 0 {
			sum += f
		}
		f, rollingFib = nextFib(rollingFib)
	}

	fmt.Println(sum)

}
