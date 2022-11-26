package main

import (
	"fmt"
	"time"
)

const input = `7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450`

func runeToInt(r rune) int {
	return int(r - '0')
}

func byteToInt(b byte) int {
	return int(b - '0')
}

func stringProduct(s string) int {
	p := 1
	for _, r := range s {
		p *= runeToInt(r)
	}
	return p
}

func intProduct(arr []int) int {
	p := 1
	for _, i := range arr {
		p *= i
	}
	return p
}

func sliceOfN(input string, len int, endIndex int) string {
	return input[endIndex-len : endIndex]
}

func intSliceOfN(input []int, len int, endIndex int) []int {
	return input[endIndex-len : endIndex]
}

func simpleAlgo(n int, input string) int {
	max := 0
	for endInd := n; endInd < len(input)-1; endInd++ {
		product := stringProduct(sliceOfN(input, n, endInd))
		if product > max {
			max = product
		}
	}
	return max
}

func fastAlgo(n int, input string) int {
	max := 0
	endInd := n
	product := 0

	for endInd < len(input)-1 {
		s := sliceOfN(input, n, endInd)
		// If the last char in the string is 0, we can skip forward
		if s[n-1] == '0' {
			endInd += n
			product = 0
			continue
		}
		// if previous product was zero, we need to recompute entirely
		if product == 0 {
			product = stringProduct(sliceOfN(input, n, endInd))
		} else { // do a O(1) calculation of the next product
			product = product * byteToInt(input[endInd-1]) / byteToInt(input[endInd-1-n])
		}

		if product > max {
			max = product
		}
		endInd++
	}
	return max
}

// same as fastAlgo, but convert the input to ints once at the beginning
func fastAlgoOptimized(n int, s string) int {
	input := make([]int, len(s))
	for i, r := range s {
		input[i] = runeToInt(r)
	}

	max := 0
	endInd := n
	product := 0

	for endInd < len(input)-1 {
		s := intSliceOfN(input, n, endInd)
		// If the last char in the string is 0, we can skip forward
		if s[n-1] == '0' {
			endInd += n
			product = 0
			continue
		}
		// if previous product was zero, we need to recompute entirely
		if product == 0 {
			product = intProduct(s)
		} else { // do a O(1) calculation of the next product
			product = product * input[endInd-1] / input[endInd-1-n]
		}

		if product > max {
			max = product
		}
		endInd++
	}
	return max
}

func onePassScanAlgo(n int, s string) int {
	input := make([]int, len(s))
	for i, r := range s {
		input[i] = runeToInt(r)
	}
	// idea : keep running product of all non-zero numbers, and number of zeros currently in the stream
	// when number of zeros is zero, check the product against the max
	currentZeroCount := 0
	product := 1
	max := 0

	for i, el := range input {
		if el == 0 {
			currentZeroCount++
		} else {
			product *= el
		}
		// if we've processed more than n numbers, start removing the trailing number
		if i >= n {
			toRemove := input[i-n]
			if toRemove == 0 {
				currentZeroCount--
			} else {
				product /= toRemove
			}
		}
		// if we've processed at least n numbers and don't currently have a zero, start checking for max
		if i >= n-1 && currentZeroCount == 0 {
			if product > max {
				max = product
			}
		}
	}
	return max
}

func main() {
	len := 13

	t0 := time.Now()

	fmt.Printf("Simple algo result: %d\n", simpleAlgo(len, input))
	t1 := time.Now()
	fmt.Printf("Simple time: %v\n", t1.Sub(t0))

	fmt.Printf("Fast algo result: %d\n", fastAlgo(len, input))
	t2 := time.Now()
	fmt.Printf("Fast time: %v\n", t2.Sub(t1))

	fmt.Printf("Faster algo result: %d\n", fastAlgoOptimized(len, input))
	t3 := time.Now()
	fmt.Printf("Faster time: %v\n", t3.Sub(t2))

	fmt.Printf("Scan algo result: %d\n", onePassScanAlgo(len, input))
	t4 := time.Now()
	fmt.Printf("Scan time: %v\n", t4.Sub(t3))

}
