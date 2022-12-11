package main

func solveSudokuInner(board [][][]byte) [][]byte {

}

func valuesInRow(board [][][]byte, row int) []byte {
	vals := []byte{}
	for _, col := range board[row] {
		if len(col) == 1 {
			vals = append(vals, col[0])
		}
	}
	return vals
}

func valuesInCol(board [][][]byte, col int) []byte {
	vals := []byte{}
	for row := 0; row < 9; row++ {
		val := board[row][col]
		if len(val) == 1 {
			vals = append(vals, val[0])
		}
	}
	return vals
}

func valuesInSquare(board [][][]byte, row, col int) []byte {
	vals := []byte{}
	row = row % 3 * 3
	col = col % 3 * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			val := board[row+i][col+j]
			if len(val) == 1 {
				vals = append(vals, val[0])
			}
		}
	}
	return vals
}

func availableNumbers(board [][][]byte, row, col int) []byte {
	taken := append(append(valuesInCol(board, col), valuesInRow(board, row)...), valuesInSquare(board, row, col)...)
	takenMap := map[byte]bool{}
	for _, n := range taken {
		takenMap[n] = true
	}
	available := []byte{}
	for b := byte('1'); b <= '9'; b++ {
		if !takenMap[b] {
			available = append(available, b)
		}
	}
	return available
}

func reduceSuperPosition(board [][][]byte, row, col int) {
	board[row][col] = availableNumbers(board, row, col)
}

func toSuperpositions(board [][]byte) [][][]byte {
	superPositionBoard := make([][][]byte, 0, 9)
	for row := 0; row < 9; row++ {
		pRow := make([][]byte, 0, 9)
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				col := col
				row := row
				pRow[col] = []byte{}
				defer func() {
					reduceSuperPosition(superPositionBoard, row, col)
				}()
			} else {
				pRow[col] = []byte{board[row][col]}
			}
		}
		superPositionBoard[row] = pRow
	}
	return superPositionBoard
}

func solveSudoku(board [][]byte) {
	board = solveSudokuInner(toSuperpositions(board))
}
