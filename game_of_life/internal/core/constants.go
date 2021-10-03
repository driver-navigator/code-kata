package core

const (
	ALIVE = true
	DEAD  = false
)

type coord struct {
	x int
	y int
}

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
