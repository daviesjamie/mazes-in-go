package grid

import "github.com/daviesjamie/mazes-in-go/util"

type Cell struct {
	Row, Column              int
	North, South, East, West *Cell
	links                    map[*Cell]bool
}

func NewCell(row, column int) *Cell {
	return &Cell{Row: row, Column: column, links: map[*Cell]bool{}}
}

func (c *Cell) Link(cell *Cell) {
	c.links[cell] = true
	cell.links[c] = true
}

func (c *Cell) Unlink(cell *Cell) {
	delete(c.links, cell)
	delete(cell.links, c)
}

func (c *Cell) Links() []*Cell {
	cells := make([]*Cell, len(c.links))
	for l := range c.links {
		cells = append(cells, l)
	}

	return cells
}

func (c *Cell) IsLinked(cell *Cell) bool {
	_, linked := c.links[cell]
	return linked
}

func (c *Cell) Neighbours() []*Cell {
	cells := []*Cell{c.North, c.South, c.East, c.West}
	return util.FilterSlice(cells, func(c *Cell) bool { return c != nil })
}
