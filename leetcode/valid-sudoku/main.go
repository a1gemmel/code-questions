package main

func bytesToInts(arr []byte) []int {
	ints := []int{}
	for _, b := range arr {
		if b != '.' {
			ints = append(ints, int(b-48))
		}
	}
	return ints
}

func getRows(board [][]byte) [][]int {
	rows := [][]int{}
	for _, row := range board {
		rows = append(rows, bytesToInts(row))
	}
	return rows
}

func getColumns(board [][]byte) [][]int {
	cols := [][]int{}
	for i := 0; i < 9; i++ {
		col := make([]byte, 0, 9)
		for j := 0; j < 9; j++ {
			col = append(col, board[j][i])
		}
		cols = append(cols, bytesToInts(col))
	}
	return cols
}

func getSquares(board [][]byte) [][]int {
	squares := [][]int{}
	for top := 0; top < 9; top += 3 {
		for left := 0; left < 9; left += 3 {
			square := []byte{}
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					square = append(square, board[top+i][left+j])
				}
			}
			squares = append(squares, bytesToInts(square))
		}
	}
	return squares
}

func isValidSet(s []int) bool {
	nums := map[int]bool{}
	for _, n := range s {
		if nums[n] {
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	sets := append(append(getRows(board), getColumns(board)...), getSquares(board)...)

	for _, s := range sets {
		if !isValidSet(s) {
			return false
		}
	}
	return true
}
