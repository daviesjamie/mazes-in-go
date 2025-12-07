package grid

import (
	"iter"
	"math/rand"
	"slices"
)

type Grid struct {
	rows, columns int
	cells         [][]*Cell
}

func NewGrid(rows, columns int) *Grid {
	g := &Grid{rows: rows, columns: columns}
	g.cells = make([][]*Cell, g.rows)

	// Create all the cells
	for row := range g.rows {
		g.cells[row] = make([]*Cell, g.columns)
		for column := range g.columns {
			g.cells[row][column] = NewCell(row, column)
		}
	}

	// Tell each cell which cells are its neighbours
	for cell := range g.Cells() {
		row, col := cell.Row, cell.Column
		cell.North = g.CellAt(row-1, col)
		cell.South = g.CellAt(row+1, col)
		cell.West = g.CellAt(row, col-1)
		cell.East = g.CellAt(row, col+1)
	}
	return g
}

func (g *Grid) Rows() iter.Seq[[]*Cell] {
	return slices.Values(g.cells)
}

func (g *Grid) Cells() iter.Seq[*Cell] {
	return func(yield func(*Cell) bool) {
		for row := range g.Rows() {
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
	row := rand.Intn(g.rows)
	col := rand.Intn(g.columns)
	return g.CellAt(row, col)
}
