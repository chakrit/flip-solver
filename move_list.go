package main

type MoveList []Move

var EmptyMoveList = make(MoveList, 0)

func GenerateMoveList(t Table) MoveList {
	height := len(t)
	width := len(t[0])

	moves := []Move{}
	yield := func(srcX, srcY, destX, destY int) {
		move := NewMove(Point{srcX, srcY}, Point{destX, destY})
		if move.Valid(t) {
			moves = append(moves, move)
		}
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
