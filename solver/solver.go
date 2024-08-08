package solver

import (
	"fmt"
	"math"

	"tetris-optimizer/reader"
)

// tetrominoesAssembler tries to place the tetrominoes on the smallest square
func TetrominoesAssembler(tetrominoes []reader.Tetro) ([][]rune, error) {
	n := len(tetrominoes)

	initialSize := int(math.Ceil(math.Sqrt(float64(n * 4))))

	for currentSize := initialSize; ; currentSize++ {
		board := createBoard(currentSize)
		if solve(board, tetrominoes, 0, currentSize) {
			return board, nil
		}
	}
}

// createBoard helps in initializing a board of a given size
func createBoard(size int) [][]rune {
	currentBoard := make([][]rune, size)
	for i := range currentBoard {
		currentBoard[i] = make([]rune, size)
		for j := range currentBoard[i] {
			currentBoard[i][j] = '.'
		}
	}
	return currentBoard
}

// solve uses recursion to try and place the tetriminoes on the board
func solve(board [][]rune, tetrominoes []reader.Tetro, index, size int) bool {
	if index == len(tetrominoes) {
		return true
	}

	for i := 0; i <= size-4; i++ {
		for j := 0; j <= size-4; j++ {
			if canPlace(board, tetrominoes[index], i, j, size) {
				place(board, tetrominoes[index], i, j, rune('A'+index))
				if solve(board, tetrominoes, index+1, size) {
					return true
				}
				remove(board, tetrominoes[index], i, j)
			}
		}
	}
	return false
}

// canPlace checks if a tetromino can be placed at the given position
func canPlace(board [][]rune, t reader.Tetro, x, y, size int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.Shape[i][j] == '#' {
				if x+i >= size || y+j >= size || board[x+i][y+j] != '.' {
					return false
				}
			}
		}
	}
	return true
}

// place puts a tetromino on the board at the given position
func place(board [][]rune, t reader.Tetro, x, y int, letter rune) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.Shape[i][j] == '#' {
				board[x+i][y+j] = letter
			}
		}
	}
}

// remove clears a tetromino from the board incase of unsuccessful placement
func remove(board [][]rune, t reader.Tetro, x, y int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.Shape[i][j] == '#' {
				board[x+i][y+j] = '.'
			}
		}
	}
}

func PrintSolution(board [][]rune) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
	}
	fmt.Println()
}
