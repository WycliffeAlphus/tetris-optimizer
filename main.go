package main

import (
	"fmt"
	"os"
	"time"

	"tetris-optimizer/reader"
	"tetris-optimizer/solver"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Provide an argument\nUsage: go run . sample.txt | cat -e")
		return
	}

	start := time.Now()

	tetrominoes, err := reader.ReadTetrominoes(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	solution := solver.TetrominoesAssembler(tetrominoes)

	duration := time.Since(start)

	solver.PrintSolution(solution)

	fmt.Printf("\nExecution time: %v\n", duration)
}
