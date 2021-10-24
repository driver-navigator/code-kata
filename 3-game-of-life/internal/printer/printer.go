package printer

import (
	"github.com/driver-navigator/code-kata/3-game-of-life/internal/core"
	clearscreen "github.com/driver-navigator/code-kata/3-game-of-life/pkg/clear-screen"
	matrixprettyprint "github.com/driver-navigator/code-kata/3-game-of-life/pkg/matrix-pretty-print"
)

type printer struct {
}

func NewPrinter() core.Printer {
	return &printer{}
}

func (p *printer) Print(population [][]bool) {
	clearscreen.Clear()
	// matrix.Print(population)
	matrixprettyprint.Print(population)
}
