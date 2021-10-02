package printer

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/driver-navigator/code-kata/game-of-life/internal/core"
	"github.com/driver-navigator/code-kata/game-of-life/pkg/matrix"
)

type printer struct {
}

func NewPrinter() core.Printer {
	return &printer{}
}

func (p *printer) Print(population [][]bool) {
	clear()
	matrix.Print(population)
}

var clearFunc = map[string]func(){
	"linux": func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
	"windows": func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	},
}

func clear() {
	value, ok := clearFunc[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("unsupported platform: " + runtime.GOOS)
	}
}
