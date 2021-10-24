package matrixprettyprint

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func Print(matrix [][]bool) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	for i := 0; i < len(matrix); i++ {
		var r table.Row
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] {
				r = append(r, "■")
			} else {
				r = append(r, " ")
			}
		}
		t.AppendRow(r)
		t.AppendSeparator()
	}

	t.SetStyle(table.StyleLight)
	t.Render()
}
