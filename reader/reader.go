package reader

import (
	"fmt"
	"os"
	"strings"
)

type Tetro struct {
	Shape  [4][4]rune
	Width  int
	Height int
}

func ReadTetrominoes(path string) ([]Tetro, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	tetrominoes := strings.Split(strings.TrimSpace(string(file)), "\n\n")
	result := make([]Tetro, 0, len(tetrominoes))

	for _, tString := range tetrominoes {
		lines := strings.Split(strings.TrimSpace(tString), "\n")
		if len(lines) != 4 {
			return nil, fmt.Errorf("ERROR")
		}

		t := Tetro{}
		hashCount := 0

		for i, line := range lines {
			if len(line) != 4 {
				return nil, fmt.Errorf("ERROR")
			}

			for j, char := range line {
				if char != '#' && char != '.' {
					return nil, fmt.Errorf("ERROR")
				}
				if char == '#' {
					hashCount++
				}
				t.Shape[i][j] = char
			}
		}

		if hashCount != 4 {
			return nil, fmt.Errorf("ERROR")
		}

		if !isValidTetromino(t) {
			return nil, fmt.Errorf("ERROR")
		}

		t = trimTetromino(t)
		result = append(result, t)
	}

	return result, nil
}

func trimTetromino(t Tetro) Tetro {
	minRow, minCol := 4, 4
	maxRow, maxCol := -1, -1

	for i, row := range t.Shape {
		for j, cell := range row {
			if cell == '#' {
				minRow = min(minRow, i)
				maxRow = max(maxRow, i)
				minCol = min(minCol, j)
				maxCol = max(maxCol, j)
			}
		}
	}

	var trimmed Tetro
	for i := minRow; i <= maxRow; i++ {
		for j := minCol; j <= maxCol; j++ {
			trimmed.Shape[i-minRow][j-minCol] = t.Shape[i][j]
		}
	}
	trimmed.Width = maxCol - minCol + 1
	trimmed.Height = maxRow - minRow + 1
	return trimmed
}

func isValidTetromino(t Tetro) bool {
	touchingSides := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.Shape[i][j] == '#' {
				touchingSides += countTouchingSides(t, i, j)
			}
		}
	}
	return touchingSides >= 6
}

func countTouchingSides(t Tetro, i, j int) int {
	count := 0
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range directions {
		newI, newJ := i+dir[0], j+dir[1]
		if newI >= 0 && newI < 4 && newJ >= 0 && newJ < 4 && t.Shape[newI][newJ] == '#' {
			count++
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
