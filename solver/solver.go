package solver

import (
	"math"

	"tetris-optimizer/reader"
)

type Board struct {
	cells []rune
	size  int
}

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

func (b *Board) get(x, y int) rune {
	return b.cells[y*b.size+x]
}

func (b *Board) set(x, y int, value rune) {
	b.cells[y*b.size+x] = value
}

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

func place(board *Board, tetro reader.Tetro, x, y int, letter rune) {
	for i := 0; i < tetro.Height; i++ {
		for j := 0; j < tetro.Width; j++ {
			if tetro.Shape[i][j] == '#' {
				board.set(x+j, y+i, letter)
			}
		}
	}
}

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

func PrintSolution(board [][]rune) {
	for _, row := range board {
		println(string(row))
	}
}
