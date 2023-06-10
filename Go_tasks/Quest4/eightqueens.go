package piscine

import "github.com/01-edu/z01"

const boardSize = 8

func EightQueens() {
	board := make([][]bool, boardSize)
	for i := range board {
		board[i] = make([]bool, boardSize)
	}

	// Start with the first row
	placeQueen(board, 0)
}

func placeQueen(board [][]bool, row int) {
	// Base case: if all queens are placed, print the solution
	if row == boardSize {
		printSolution(board)
		return
	}

	for col := 0; col < boardSize; col++ {
		if isSafe(board, row, col) {
			board[row][col] = true
			placeQueen(board, row+1)
			board[row][col] = false
		}
	}
}

func isSafe(board [][]bool, row, col int) bool {
	// Check if there is a queen in the same column
	for i := 0; i < row; i++ {
		if board[i][col] {
			return false
		}
	}

	// Check if there is a queen in the upper-left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}

	// Check if there is a queen in the upper-right diagonal
	for i, j := row-1, col+1; i >= 0 && j < boardSize; i, j = i-1, j+1 {
		if board[i][j] {
			return false
		}
	}

	return true
}

func printSolution(board [][]bool) {
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			if board[row][col] {
				// Add 1 to the column index since the puzzle is 1-based
				r := '0' + rune(col+1)
				if err := z01.PrintRune(r); err != nil {
					panic(err)
				}
			}
		}
	}
	if err := z01.PrintRune('\n'); err != nil {
		panic(err)
	}
}
