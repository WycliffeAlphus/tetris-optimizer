package solver

import (
	"math"

	"tetris-optimizer/reader"
)

type Board struct {
	cells []rune
	size  int
}

// newBoard takes size as an integer and returns a pointer to a board of size by size
func newBoard(size int) *Board {
	cells := make([]rune, size*size)
	for i := range cells {
		cells[i] = '.'
	}
	return &Board{
		cells: cells,
		size:  size,
	}
}

// TetrominoesAssembler takes a slice of tetrominoes to be placed on the board and returns 2D slice of runes with all tetrominoes placed correctly
func TetrominoesAssembler(tetrominoes []reader.Tetro) [][]rune {
	totalCells := len(tetrominoes) * 4
	minSize := int(math.Ceil(math.Sqrt(float64(totalCells))))

	for size := minSize; ; size++ {
		board := newBoard(size)
		if solve(board, tetrominoes, 0) {
			return boardToRuneSlice(board)
		}
	}
}

// solve takes a pointer to a board, a slice of terominoes, current tetro index and returns a true if a solution for placement is found, defaults to false
func solve(board *Board, tetrominoes []reader.Tetro, index int) bool {
	if index == len(tetrominoes) {
		return true
	}

	tetro := tetrominoes[index]
	letter := rune('A' + index)

	for y := 0; y <= board.size-tetro.Height; y++ {
		for x := 0; x <= board.size-tetro.Width; x++ {
			if canPlace(board, tetro, x, y) {
				place(board, tetro, x, y, letter)
				if solve(board, tetrominoes, index+1) {
					return true
				}
				place(board, tetro, x, y, '.') // backtrack
			}
		}
	}
	return false
}

// canPlace takes a pointer to board, a tetro, the coordinates to place the tetro, and returns true or false if it can be placed at the position
func canPlace(board *Board, tetro reader.Tetro, x, y int) bool {
	for i := 0; i < tetro.Height; i++ {
		for j := 0; j < tetro.Width; j++ {
			if tetro.Shape[i][j] == '#' && board.get(x+j, y+i) != '.' {
				return false
			}
		}
	}
	return true
}

// place takes pointer to a board, a tetro, the coordinates to place the tetro, and the rune to place
// It has no return value
func place(board *Board, tetro reader.Tetro, x, y int, letter rune) {
	for i := 0; i < tetro.Height; i++ {
		for j := 0; j < tetro.Width; j++ {
			if tetro.Shape[i][j] == '#' {
				board.set(x+j, y+i, letter)
			}
		}
	}
}

// get retrieves the rune located at the specified (x, y) coordinates on the Board.
func (b *Board) get(x, y int) rune {
	return b.cells[y*b.size+x] // calculates the index in the 1D slice b.cells that corresponds to the 2D coordinates (x, y)
}

// set assigns a given rune value to the specified (x, y) coordinates on the Board.
func (b *Board) set(x, y int, value rune) {
	b.cells[y*b.size+x] = value
}

// boardToRuneSlice takes a pointer to a board and returns a 2D slice of runes
func boardToRuneSlice(board *Board) [][]rune {
	result := make([][]rune, board.size)
	for i := range result {
		result[i] = make([]rune, board.size)
		for j := range result[i] {
			result[i][j] = board.get(j, i)
		}
	}
	return result
}

// PrintSolution takes the board and prints it
// It has no return value
func PrintSolution(board [][]rune) {
	for _, row := range board {
		println(string(row))
	}
}
