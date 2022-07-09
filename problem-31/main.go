package main

import "fmt"

func main() {

	const sum = 200

	// for each nth row in the backpack, the i-th entry is how many ways the sum i
	// can be made with the first n coins
	backpack := [][]int{}

	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}

	for _, _ = range coins {
		backpack = append(backpack, make([]int, sum+1))
	}

	// initialize the first row of the backpack with pennies - one way to make each number
	for i := 0; i <= sum; i++ {
		backpack[0][i] = 1
	}

	for coinIndex, coin := range coins {
		if coinIndex == 0 {
			continue
		}
		for i := 0; i <= sum; i += 1 {
			ways := 0
			if i >= coin {
				// ways if we choose the current coin
				ways += backpack[coinIndex][i-coin]
			}
			// ways if we don't choose the current coin
			ways += backpack[coinIndex-1][i]

			backpack[coinIndex][i] = ways
		}
	}

	fmt.Println(backpack[len(coins)-1][sum])

}
