package main

import (
	"fmt"
	"os"

	"tetris-optimizer/reader"
	"tetris-optimizer/solver"
)

func main() {
	if len(os.Args) != 2 {

		// fmt.Println("ERROR")
		// return
		fmt.Fprintln(os.Stderr, "Provide an argument\nUsage: go run . sample.txt | cat -e")
		return
	}

	tetrominoes, err := reader.ReadTetrominoes(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	solution, err := solver.TetrominoesAssembler(tetrominoes)
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	solver.PrintSolution(solution)
}
