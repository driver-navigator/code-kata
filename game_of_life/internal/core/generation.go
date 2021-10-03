package core

import (
	"errors"

	"github.com/driver-navigator/code-kata/game-of-life/pkg/matrix"
)

type generation struct {
	rows       int
	cols       int
	population [][]bool
}

func NewGeneration(rows, cols int, population [][]bool) (Generation, error) {
	if len(population) != rows {
		return nil, errors.New("mismatch number of rows")
	}
	if len(population[0]) != cols {
		return nil, errors.New("mismatch number of cols")
	}

	return &generation{
		rows:       rows,
		cols:       cols,
		population: population,
	}, nil
}

func (g *generation) AdvanceGeneration() {
	genSpan := matrix.Clone(g.population)

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			var numOfAliveNeighbors = g.numOfAliveNeighbors(i, j)
			if genSpan[i][j] && (numOfAliveNeighbors < 2 || numOfAliveNeighbors > 3) {
				genSpan[i][j] = DEAD
			} else if genSpan[i][j] && (numOfAliveNeighbors == 2 || numOfAliveNeighbors == 3) {
				continue
			} else if !genSpan[i][j] && numOfAliveNeighbors == 3 {
				genSpan[i][j] = ALIVE
			}
		}
	}

	g.population = matrix.Clone(genSpan)
}

func (g *generation) Population() [][]bool {
	return g.population
}

func (g *generation) numOfAliveNeighbors(x, y int) int {
	alive := 0

	if g.isAlive(x+LEFT.x, y+LEFT.y) {
		alive++
	}

	if g.isAlive(x+RIGHT.x, y+RIGHT.y) {
		alive++
	}

	if g.isAlive(x+TOP.x, y+TOP.y) {
		alive++
	}

	if g.isAlive(x+BOTTOM.x, y+BOTTOM.y) {
		alive++
	}

	if g.isAlive(x+TOP_LEFT.x, y+TOP_LEFT.y) {
		alive++
	}

	if g.isAlive(x+TOP_RIGHT.x, y+TOP_RIGHT.y) {
		alive++
	}

	if g.isAlive(x+BOTTOM_LEFT.x, y+BOTTOM_LEFT.y) {
		alive++
	}

	if g.isAlive(x+BOTTOM_RIGHT.x, y+BOTTOM_RIGHT.y) {
		alive++
	}

	return alive
}

func (g *generation) isAlive(x, y int) bool {
	if g.isOnTheEdge(x, y) {
		return false
	}

	return g.population[x][y]
}

func (g *generation) isOnTheEdge(x, y int) bool {
	if x < 0 || y < 0 {
		return true
	}

	if x > g.rows-1 || y > g.cols-1 {
		return true
	}

	return false
}
