package main

import (
	"fmt"
	"time"
)

func collatzLength(n int, memoize map[int]int) int {
	var chain []int

	for n != 1 {
		chain = append(chain, n)
		if l, found := memoize[n]; found {
			for i, el := range chain {
				memoize[el] = len(chain) - i + 1 + l
			}
			return l + len(chain)
		}
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}

	for i, el := range chain {
		memoize[el] = len(chain) - i + 1
	}
	return len(chain) + 1
}

func main() {
	m := map[int]int{}

	max := 0
	maxItem := 0

	t0 := time.Now()

	for i := 1000000; i > 1; i-- {
		l := collatzLength(i, m)
		if l > max {
			max = l
			maxItem = i
		}
	}

	fmt.Println(maxItem, max)
	fmt.Printf("Run time: %v\n", time.Now().Sub(t0))

}
