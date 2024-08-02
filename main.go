package main

import (
	"fmt"
	"os"

	"tetris-optimizer/reader"
)

func main() {
	if len(os.Args) != 2 {

		fmt.Println("ERROR")
		return
		// 		fmt.Fprintf(os.Stderr, `provide an argument
		// Usage: go run . sample.txt | cat -e`)
		// 		fmt.Println()
		// 		os.Exit(0)
	}

	tetrominoes, err := reader.ReadTetrominoes(os.Args[1])
	if err != nil {
		fmt.Println("Error")
		return
	}
}
