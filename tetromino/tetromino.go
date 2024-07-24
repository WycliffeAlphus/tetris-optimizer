package tetromino

import (
	"fmt"
	"os"
	"strings"
)

type Tetro struct {
	shape [4][4]rune
}

var t Tetro

func Tetromino() (*Tetro, error) {
	file, err := os.ReadFile(os.Args[1])

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	if err != nil {
		return nil, fmt.Errorf("ERROR")
	}

	if len(lines) != 4 {
		return nil, fmt.Errorf("ERROR")
	}
	var hashCount int
	for i := 0; i < 4; i++ {
		if len(lines[i]) != 4 {
			return nil, fmt.Errorf("ERROR")
		}
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
	return &t, nil
}
