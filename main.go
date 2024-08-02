package main

import (
	"fmt"

	"tetris-optimizer/tetromino"
)

func main() {
	if !tetromino.Tetromino() {
		fmt.Println("ERROR")
	} else {
		fmt.Println("SUCCESS")
	}
}
