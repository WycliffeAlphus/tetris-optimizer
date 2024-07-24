package main

import (
	"fmt"

	"tetris-optimizer/tetromino"
)

func main() {
	res, err := tetromino.Tetromino()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
