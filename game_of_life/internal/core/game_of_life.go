package core

const (
	ALIVE = true
	DEAD  = false
)

var (
	LEFT         = coord{-1, 0}
	RIGHT        = coord{+1, 0}
	TOP          = coord{0, -1}
	BOTTOM       = coord{0, +1}
	TOP_LEFT     = coord{-1, -1}
	TOP_RIGHT    = coord{+1, -1}
	BOTTOM_LEFT  = coord{-1, +1}
	BOTTOM_RIGHT = coord{+1, +1}
)

type coord struct {
	x int
	y int
}

type gameOfLife struct {
	rows int16
	cols int16
	gen  [][]bool
}

func NewGameOfLife(rows, cols int16) gameOfLife {
	return gameOfLife{
		rows: rows,
		cols: cols,
		gen:  [][]bool{},
	}
}

func (g *gameOfLife) AdvanceGeneration() {
	genSpan := clone(g.gen)

	for i := 0; i < int(g.rows); i++ {
		for j := 0; j < int(g.cols); j++ {
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

	g.gen = clone(genSpan)
}

func (g *gameOfLife) CurrentGeneration() [][]bool {
	return g.gen
}

func (g *gameOfLife) SetCurrentGeneration(gen [][]bool) {
	if len(gen) != int(g.rows) {
		panic("mismatch rows")
	}
	if len(gen[0]) != int(g.cols) {
		panic("mismatch cols")
	}

	g.gen = gen
}

func (g *gameOfLife) numOfAliveNeighbors(x, y int) int {
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

func (g *gameOfLife) isAlive(x, y int) bool {
	if g.onTheEdge(x, y) {
		return false
	}

	return g.gen[x][y]
}

func (g *gameOfLife) onTheEdge(x, y int) bool {
	if x < 0 || y < 0 {
		return true
	}

	if x > int(g.rows-1) || y > int(g.cols-1) {
		return true
	}

	return false
}

func clone(matrix [][]bool) [][]bool {
	duplicate := make([][]bool, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]bool, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}
