package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/daviesjamie/mazes-in-go/algorithm"
	"github.com/daviesjamie/mazes-in-go/grid"
)

type mapping struct {
	algorithm func(*grid.Grid)
	aliases   []string
}

var mappings []mapping = []mapping{
	{algorithm: algorithm.BinaryTree, aliases: []string{"binary_tree", "binary-tree", "bt"}},
	{algorithm: algorithm.Sidewinder, aliases: []string{"sidewinder", "sw"}},
}

func main() {
	algorithmArg := flag.String("algorithm", "", "the algorithm to use to generate the maze")
	rowsArg := flag.Int("rows", 4, "the height of the maze, in cells")
	colsArg := flag.Int("cols", 4, "the width of the maze, in cells")

	flag.Parse()

	a, err := mapArgToAlgorithm(*algorithmArg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid algorithm: '%s'\n\n", *algorithmArg)
		printAlgorithmOptions()
		os.Exit(1)
	}

	g := grid.NewGrid(*rowsArg, *colsArg)
	a(g)
	fmt.Println(g)
}

func mapArgToAlgorithm(arg string) (func(*grid.Grid), error) {
	lcName := strings.ToLower(arg)

	for _, m := range mappings {
		if slices.Contains(m.aliases, lcName) {
			return m.algorithm, nil
		}
	}

	return nil, errors.New("unknown algorithm")
}

func printAlgorithmOptions() {
	fmt.Fprintln(os.Stderr, "Possible algorithm options:")
	for _, m := range mappings {
		fmt.Fprintf(os.Stderr, " - %s\n", strings.Join(m.aliases, ", "))
	}
}
