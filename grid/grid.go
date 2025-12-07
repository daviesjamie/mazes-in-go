package grid

import (
	"context"
	"math/rand"
)

type Grid struct {
	Rows, Columns int
	cells         [][]*Cell
}

func NewGrid(ctx context.Context, rows, columns int) *Grid {
	grid := &Grid{Rows: rows, Columns: columns}
	grid.PrepareGrid()
	grid.ConfigureCells(ctx)
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

func (g *Grid) ConfigureCells(ctx context.Context) {
	for cell := range g.EachCell(ctx) {
		row, col := cell.Row, cell.Column
		cell.North = g.CellAt(row-1, col)
		cell.South = g.CellAt(row+1, col)
		cell.West = g.CellAt(row, col-1)
		cell.East = g.CellAt(row, col+1)
	}
}

func (g *Grid) EachRow(ctx context.Context) <-chan []*Cell {
	rowStream := make(chan []*Cell)
	go func() {
		defer close(rowStream)

		for _, row := range g.cells {
			select {
			case <-ctx.Done():
				return
			case rowStream <- row:
			}
		}
	}()
	return rowStream
}

func (g *Grid) EachCell(ctx context.Context) <-chan *Cell {
	cellStream := make(chan *Cell)
	go func() {
		defer close(cellStream)

		for row := range g.EachRow(ctx) {
			for _, cell := range row {
				select {
				case <-ctx.Done():
					return
				case cellStream <- cell:
				}
			}
		}
	}()
	return cellStream
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
