package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var queenPositions [8]int
	solveEightQueens(&queenPositions, 0)
}

func solveEightQueens(pos *[8]int, col int) {
	if col == 8 {
		printSolutions(pos)
		return
	}

	for row := 1; row <= 8; row++ {
		if isValid(pos, col, row) {
			pos[col] = row
			solveEightQueens(pos, col+1)
		}
	}
}

func isValid(pos *[8]int, colInd int, rowInd int) bool {
	for col := 0; col < colInd; col++ {
		row := pos[col]
		if row == rowInd ||
			(row-col) == (rowInd-colInd) ||
			(row+col) == (rowInd+colInd) {
			return false
		}
	}
	return true
}

func printSolutions(pos *[8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(pos[i]) + '0')
	}
	z01.PrintRune('\n')
}
