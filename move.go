package main

import "strconv"

type Point struct {
	X int
	Y int
}

func (p Point) DistanceTo(another Point) int {
	// dumb version, since our problem space doesn't have diagonoal moves
	distX := another.X - p.X
	if distX < 0 {
		distX = -distX
	}
	distY := another.Y - p.Y
	if distY < 0 {
		distY = -distY
	}
	return distX + distY
}

func (p Point) String() string {
	return "{" + strconv.Itoa(p.X) + ", " + strconv.Itoa(p.Y) + "}"
}

// move
type Move struct {
	Src  Point
	Dest Point
}

func NewMove(src Point, dest Point) Move {
	return Move{src, dest}
}

func (m Move) String() string {
	return m.Src.String() + " -> " + m.Dest.String()
}

func (m Move) ApplyTo(table *Table) {
	// assume m.Src < m.Dest always
	t := *table
	dist := m.Src.DistanceTo(m.Dest) + 1
	dist /= 2

	if m.Src.X == m.Dest.X {
		for i := 0; i < dist; i++ {
			t.SwapY(m.Src.X, m.Src.Y+i, m.Dest.Y-i)
		}

	} else if m.Src.Y == m.Dest.Y {
		for i := 0; i < dist; i++ {
			t.SwapX(m.Src.Y, m.Src.X+i, m.Dest.X-i)
		}

	} else {
		panic("Diagonal move is not supported!")
	}
}

// move list
type MoveList []Move

func GenerateMovesFromTable(table *Table) MoveList {
	t := *table
	height := len(t)
	width := len(t[0])

	moves := []Move{}
	yield := func(srcX, srcY, destX, destY int) Move {
		src := Point{srcX, srcY}
		dest := Point{destX, destY}
		move := NewMove(src, dest)
		moves = append(moves, move)
		return move
	}

	// horizontal moves
	for y := 0; y < height; y++ {
		for x := 0; x < (width - 2); x++ {
			for destX := x + 2; destX < width; destX++ {
				yield(x, y, destX, y)
			}
		}
	}

	// vertical moves
	for x := 0; x < width; x++ {
		for y := 0; y < (height - 2); y++ {
			for destY := y + 2; destY < height; destY++ {
				yield(x, y, x, destY)
			}
		}
	}

	return moves
}

func (list MoveList) String() string {
	s := ""
	for _, move := range list {
		s += move.String() + "\n"
	}
	return s
}
