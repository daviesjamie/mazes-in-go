package grid

import "strings"

func (g *Grid) String() string {
	output := "+" + strings.Repeat("---+", g.Columns) + "\n"

	for row := range g.EachRow() {
		middle := "|"
		bottom := "+"

		for _, cell := range row {
			body := "   "
			if cell.IsLinked(cell.East) {
				middle += body + " "
			} else {
				middle += body + "|"
			}

			corner := "+"
			if cell.IsLinked(cell.South) {
				bottom += "   " + corner
			} else {
				bottom += "---" + corner
			}
		}

		output += middle + "\n"
		output += bottom + "\n"
	}

	return output
}
