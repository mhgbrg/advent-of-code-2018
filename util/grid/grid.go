package grid

import (
	"fmt"
	"math"
)

type Position struct {
	X int
	Y int
}

func Origo() Position {
	return Position{0, 0}
}

func (pos Position) Above() Position {
	return Position{pos.X, pos.Y - 1}
}

func (pos Position) Below() Position {
	return Position{pos.X, pos.Y + 1}
}

func (pos Position) Left() Position {
	return Position{pos.X - 1, pos.Y}
}

func (pos Position) Right() Position {
	return Position{pos.X + 1, pos.Y}
}

func (pos Position) String() string {
	return fmt.Sprintf("(%d, %d)", pos.X, pos.Y)
}

func (pos Position) Around1() []Position {
	return []Position{
		Position{pos.X, pos.Y - 1},
		Position{pos.X - 1, pos.Y},
		Position{pos.X + 1, pos.Y},
		Position{pos.X, pos.Y + 1},
	}
}

func (pos Position) Around2() []Position {
	return []Position{
		Position{pos.X - 1, pos.Y - 1},
		Position{pos.X, pos.Y - 1},
		Position{pos.X + 1, pos.Y - 1},
		Position{pos.X - 1, pos.Y},
		Position{pos.X + 1, pos.Y},
		Position{pos.X - 1, pos.Y + 1},
		Position{pos.X, pos.Y + 1},
		Position{pos.X + 1, pos.Y + 1},
	}
}

type Limits struct {
	XMin int
	XMax int
	YMin int
	YMax int
}

func CalcLimits(posMap map[Position]bool) Limits {
	xMin, xMax, yMin, yMax := math.MaxInt32, -math.MaxInt32, math.MaxInt32, -math.MaxInt32
	for p := range posMap {
		if p.X < xMin {
			xMin = p.X
		}
		if p.X > xMax {
			xMax = p.X
		}
		if p.Y < yMin {
			yMin = p.Y
		}
		if p.Y > yMax {
			yMax = p.Y
		}
	}
	return Limits{xMin, xMax, yMin, yMax}
}
