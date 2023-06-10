package main

import (
	"os"

	"github.com/01-edu/z01"
)

const sudokuSize = 9

type Board [sudokuSize][sudokuSize]int

func main() {
	gamesetup := os.Args[1:]

	// Checking if rows are equal to the sudokuSize
	if len(gamesetup) != sudokuSize {
		printError()
		return
	}

	var board Board
	for i := 0; i < sudokuSize; i++ {

		// Checking if columns are equal to the sudokuSize
		if len(gamesetup[i]) != sudokuSize {
			printError()
			return
		}

		for j := 0; j < sudokuSize; j++ {
			if gamesetup[i][j] != '.' {
				board[i][j] = int(gamesetup[i][j] - '0')
			}
		}
	}

	// Checking if there are duplicates in the initial setup
	for n := 1; n <= sudokuSize; n++ {

		// Check rows and columns
		count := 0
		for i := 0; i < sudokuSize; i++ {
			for j := 0; j < sudokuSize; j++ {
				if board[i][j] == n {
					count += 1
				}
			}
		}
		if count > 9 {
			printError()
			return
		}

		// Check boxes
		for i := 0; i < sudokuSize; i += 3 {
			for j := 0; j < sudokuSize; j += 3 {
				countInBox := 0
				for r := 0; r < 3; r++ {
					for c := 0; c < 3; c++ {
						if board[i-i%3+r][j-j%3+c] == n {
							countInBox += 1
						}
					}
				}
				if countInBox > 1 {
					printError()
					return
				}
			}
		}

	}

	if !solved(&board) {
		printError()
		return
	}

	printBoard(&board)
}

func solved(board *Board) bool {
	row, col := findBlanks(board)
	if row == -1 && col == -1 {
		return true
	}

	for n := 1; n <= sudokuSize; n++ {
		if isCorrect(board, row, col, n) {
			board[row][col] = n
			if solved(board) {
				return true
			}
			board[row][col] = 0
		}
	}

	return false
}

func findBlanks(board *Board) (int, int) {
	for row := 0; row < sudokuSize; row++ {
		for col := 0; col < sudokuSize; col++ {
			if board[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

func isCorrect(board *Board, row, col, n int) bool {
	// Check the row and column
	for i := 0; i < sudokuSize; i++ {
		if board[row][i] == n || board[i][col] == n {
			return false
		}
	}

	// Check the corresponding box
	boxRow := row - row%3
	boxCol := col - col%3
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[boxRow+r][boxCol+c] == n {
				return false
			}
		}
	}

	return true
}

func printBoard(board *Board) {
	for row := 0; row < sudokuSize; row++ {
		for col := 0; col < sudokuSize; col++ {
			z01.PrintRune(rune(board[row][col] + '0'))
			if col != 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func printError() {
	err := "Error"
	for _, letter := range err {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
