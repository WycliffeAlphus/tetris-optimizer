package solver

import (
	"math"

	"tetris-optimizer/reader"
)

// tetrominoesAssembler tries to place the tetrominoes on the smallest square
func tetrominoesAssembler(tetrominoes []reader.Tetro) ([][]rune, error) {
	n := len(tetrominoes)

	initialSize := int(math.Ceil(math.Sqrt(float64(n * 4))))

	for currentSize := initialSize; ; currentSize++ {
		board := createBoard(currentSize)
		if solve(currentSize, tetrominoes, 0, currentSize) {
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
