package grid

import (
	"iter"
	"math/rand"
	"slices"
)

type Grid struct {
	Rows, Columns int
	cells         [][]*Cell
}

func NewGrid(rows, columns int) *Grid {
	grid := &Grid{Rows: rows, Columns: columns}
	grid.PrepareGrid()
	grid.ConfigureCells()
	return grid
}

func (g *Grid) PrepareGrid() {
	g.cells = make([][]*Cell, g.Rows)

	for row := range g.Rows {
		g.cells[row] = make([]*Cell, g.Columns)
		for column := range g.Columns {
			g.cells[row][column] = NewCell(row, column)
		}
	}
}

func (g *Grid) ConfigureCells() {
	for cell := range g.EachCell() {
		row, col := cell.Row, cell.Column
		cell.North = g.CellAt(row-1, col)
		cell.South = g.CellAt(row+1, col)
		cell.West = g.CellAt(row, col-1)
		cell.East = g.CellAt(row, col+1)
	}
}

func (g *Grid) EachRow() iter.Seq[[]*Cell] {
	return slices.Values(g.cells)
}

func (g *Grid) EachCell() iter.Seq[*Cell] {
	return func(yield func(*Cell) bool) {
		for row := range g.EachRow() {
			for _, cell := range row {
				if !yield(cell) {
					return
				}
			}
		}
	}
}

func (g *Grid) CellAt(row, column int) *Cell {
	if row < 0 || row > len(g.cells)-1 {
		return nil
	}

	if column < 0 || column > len(g.cells[row])-1 {
		return nil
	}

	return g.cells[row][column]
}

func (g *Grid) RandomCell() *Cell {
	row := rand.Intn(g.Rows)
	col := rand.Intn(g.Columns)
	return g.CellAt(row, col)
}
