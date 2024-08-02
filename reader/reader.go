package reader

import (
	"fmt"
	"os"
	"strings"
)

type Tetro struct {
	shape [4][4]rune
}

var t Tetro

func ReadTetrominoes(path string) ([]Tetro, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Split the file content by double newline to separate tetrominoes
	tetrominoes := strings.Split(strings.TrimSpace(string(file)), "\n\n")

	for _, tString := range tetrominoes {
		lines := strings.Split(strings.TrimSpace(tString), "\n")

		// Check the number of lines for each tetromino
		if len(lines) != 4 {
			return nil, fmt.Errorf("ERROR")
		}

		var hashCount int
		for i := 0; i < 4; i++ {
			if len(lines[i]) != 4 {
				return nil, fmt.Errorf("ERROR")
			}
			// Check the number of hashes
			for j := 0; j < 4; j++ {
				if lines[i][j] != '#' && lines[i][j] != '.' {
					return nil, fmt.Errorf("ERROR")
				}
				if lines[i][j] == '#' {
					hashCount++
				}
				t.shape[i][j] = rune(lines[i][j])
			}
		}

		if hashCount != 4 {
			return nil, fmt.Errorf("ERROR")
		}

		// Check number of touching sides
		if !isValidTetromino(t) {
			return nil, fmt.Errorf("ERROR")
		}
	}
	var tetromino []Tetro
	tetromino = append(tetromino, t)
	return tetromino, nil
}

func isValidTetromino(t Tetro) bool {
	touchingSides := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.shape[i][j] == '#' {
				if i > 0 && t.shape[i-1][j] == '#' {
					touchingSides++
				}
				if i < 3 && t.shape[i+1][j] == '#' {
					touchingSides++
				}
				if j > 0 && t.shape[i][j-1] == '#' {
					touchingSides++
				}
				if j < 3 && t.shape[i][j+1] == '#' {
					touchingSides++
				}
			}
		}
	}
	return touchingSides >= 6
}
