package main

import (
	"fmt"
	"math"
)

var A2Length = math.Pow(2, -3.0/4.0)
var A2Width = math.Pow(2, -5.0/4.0)

func main() {
	var smallestAPaper int

	fmt.Scanf("%d", &smallestAPaper)
	papers := make([]int, smallestAPaper+1) // convention is papers[i] contains the count of Ai papers. papers[0] is unused
	for i := 2; i <= smallestAPaper; i++ {
		var nPaper int
		fmt.Scanf("%d", &nPaper)
		papers[i] = nPaper
	}
	totalTape, possible := makeNthAPaperITimes(1, 1, papers)
	if possible {
		fmt.Println(totalTape)
	} else {
		fmt.Println("impossible")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func makeNthAPaperITimes(n int, i int, papers []int) (tapeUsed float64, possible bool) {
	//fmt.Printf("Trying to make %d A%d papers from %v\n", i, n, papers)
	var totalTape float64 = 0
	if papers[n] >= i {
		//fmt.Printf("Already have enough A%d papers\n", i)
		return 0, true
	}
	if n+1 >= len(papers) { // need to produce a smaller paper than possible
		//fmt.Printf("Don't have A%d papers to use\n", n+1)
		return 0, false
	}
	need := i - papers[n]
	subTape, success := makeNthAPaperITimes(n+1, need*2, papers)
	if !success {
		return 0, false
	}
	totalTape += subTape
	papers[n+1] -= need * 2
	papers[n] = i
	//fmt.Printf("Making %d A%d papers from A%d papers\n", need, n, n+1)
	currentTape := float64(need) * NthWidth(n)
	//fmt.Printf("Used %f tape to do so (%d lengths of %f)\n", currentTape, need, NthWidth(n))
	totalTape += currentTape
	return totalTape, true
}

func NthWidth(n int) float64 {
	if n == 1 {
		return A2Length
	}
	if n == 2 {
		return A2Width
	}
	return NthLength(n) / math.Sqrt(2)
}
func NthLength(n int) float64 {
	if n == 2 {
		return A2Length
	}
	return NthWidth(n - 1)
}
