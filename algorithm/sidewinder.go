package algorithm

import (
	"math/rand"

	"github.com/daviesjamie/mazes-in-go/grid"
	"github.com/daviesjamie/mazes-in-go/util"
)

func Sidewinder(g *grid.Grid) {
	for row := range g.Rows() {
		var run []*grid.Cell

		for _, cell := range row {
			run = append(run, cell)

			atEastBoundary := cell.East == nil
			atNorthBoundary := cell.North == nil

			shouldCloseOut := atEastBoundary || (!atNorthBoundary && rand.Intn(2) == 0)

			if shouldCloseOut {
				c := util.SampleSlice(run)
				if c.North != nil {
					c.Link(c.North)
				}
				run = []*grid.Cell{}
			} else {
				cell.Link(cell.East)
			}
		}
	}
}
