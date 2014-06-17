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
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}
