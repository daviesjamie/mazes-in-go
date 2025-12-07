package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/daviesjamie/mazes-in-go/algorithm"
	"github.com/daviesjamie/mazes-in-go/grid"
)

func main() {
	algorithmArg := flag.String("algorithm", "", "the algorithm to use to generate the maze")
	rowsArg := flag.Int("rows", 4, "the height of the maze, in cells")
	colsArg := flag.Int("cols", 4, "the width of the maze, in cells")

	flag.Parse()

	a, err := mapArgToAlgorithm(*algorithmArg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid algorithm: '%s'\n", *algorithmArg)
		os.Exit(1)
	}

	g := grid.NewGrid(*rowsArg, *colsArg)
	a(g)
	fmt.Println(g)
}

func mapArgToAlgorithm(arg string) (func(*grid.Grid), error) {
	switch strings.ToLower(arg) {
	case "binary_tree", "binary-tree", "bt":
		return algorithm.BinaryTree, nil
	case "sidewinder", "sw":
		return algorithm.Sidewinder, nil
	default:
		return nil, errors.New("unknown algorithm")
	}
}
