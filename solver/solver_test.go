package solver

import (
	"testing"

	"tetris-optimizer/reader"
)

// mockTetro creates a simple mock of a Tetromino for testing purposes
func mockTetro(shape [4][4]rune) reader.Tetro {
	height := 0
	width := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if shape[i][j] == '#' {
				if i+1 > height {
					height = i + 1
				}
				if j+1 > width {
					width = j + 1
				}
			}
		}
	}
	return reader.Tetro{
		Shape:  shape,
		Width:  width,
		Height: height,
	}
}

// TestTetrominoesAssembler tests the TetrominoesAssembler function
func TestTetrominoesAssembler(t *testing.T) {
	// Define test cases
	tests := []struct {
		name        string
		tetrominoes []reader.Tetro
		expected    [][]rune
	}{
		{
			name: "Single tetromino",
			tetrominoes: []reader.Tetro{
				mockTetro([4][4]rune{
					{'#', '#', '.', '.'},
					{'#', '#', '.', '.'},
					{'.', '.', '.', '.'},
					{'.', '.', '.', '.'},
				}),
			},
			expected: [][]rune{
				{'A', 'A'},
				{'A', 'A'},
			},
		},
		{
			name: "Two tetrominoes",
			tetrominoes: []reader.Tetro{
				mockTetro([4][4]rune{
					{'#', '#', '.', '.'},
					{'#', '#', '.', '.'},
					{'.', '.', '.', '.'},
					{'.', '.', '.', '.'},
				}),
				mockTetro([4][4]rune{
					{'#', '.', '.', '.'},
					{'.', '.', '.', '.'},
					{'.', '.', '.', '.'},
					{'.', '.', '.', '.'},
				}),
			},
			expected: [][]rune{
				{'A', 'A', 'B'},
				{'A', 'A', '.'},
				{'.', '.', '.'},
			},
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TetrominoesAssembler(tt.tetrominoes)
			if !equalBoard(got, tt.expected) {
				t.Errorf("TetrominoesAssembler() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// equalBoard checks if two boards are equal
func equalBoard(a, b [][]rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
