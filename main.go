package main

import (
	"fmt"

	"github.com/daviesjamie/mazes-in-go/algorithm"
	"github.com/daviesjamie/mazes-in-go/grid"
)

func main() {
	g := grid.NewGrid(4, 4)
	algorithm.BinaryTree(g)
	fmt.Println(g)
}
