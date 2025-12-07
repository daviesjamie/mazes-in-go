package algorithm

import (
	"context"

	"github.com/daviesjamie/mazes-in-go/grid"
	"github.com/daviesjamie/mazes-in-go/util"
)

func BinaryTree(ctx context.Context, g grid.Grid) {
	for cell := range g.EachCell(ctx) {
		neighbours := []*grid.Cell{cell.North, cell.East}
		neighbours = util.FilterSlice(neighbours, func(cell *grid.Cell) bool { return cell != nil })

		if len(neighbours) > 0 {
			neighbour := util.SampleSlice(neighbours)
			cell.Link(neighbour)
		}
	}
}
