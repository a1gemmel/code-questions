package main

import "fmt"

func main() {
	var boardSize int
	var startingPos int
	var magicNum int
	fmt.Scanf("%d %d %d", &boardSize, &startingPos, &magicNum)

	board := []int{}
	for i := 0; i < boardSize; i++ {
		var tile int
		fmt.Scanf("%d", &tile)
		board = append(board, tile)
	}

	// fmt.Println(board, startingPos, magicNum)

	gameLength(board, startingPos-1, magicNum)
}

func gameLength(board []int, pos, magicNum int) {
	length := 0
	visited := map[int]bool{}
	for {
		if board[pos] == magicNum {
			fmt.Println("magic")
			break
		}
		if visited[pos] {
			fmt.Println("cycle")
			break
		}
		visited[pos] = true
		pos = pos + board[pos]
		length++
		if pos >= len(board) {
			fmt.Println("right")
			break
		}
		if pos < 0 {
			fmt.Println("left")
			break
		}
	}
	fmt.Println(length)
}
