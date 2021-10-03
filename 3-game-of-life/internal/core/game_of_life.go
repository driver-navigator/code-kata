package core

import (
	"time"

	"github.com/driver-navigator/code-kata/3-game-of-life/pkg/matrix"
)

type GameOfLife interface {
	Play()
}

type Generation interface {
	Population() [][]bool
	AdvanceGeneration()
}

type Printer interface {
	Print([][]bool)
}

type gameOfLife struct {
	generation Generation
	printer    Printer
}

func NewGameOfLife(initialGeneration Generation, printer Printer) GameOfLife {
	return &gameOfLife{
		generation: initialGeneration,
		printer:    printer,
	}
}

func (g *gameOfLife) Play() {
	for {
		g.generation.AdvanceGeneration()
		newGenPopulation := g.generation.Population()
		g.printer.Print(newGenPopulation)

		if !hasSurvived(newGenPopulation) {
			break
		}
		time.Sleep(time.Second)
	}
}

func hasSurvived(genPopulation [][]bool) bool {
	return matrix.Exists(genPopulation, true)
}
