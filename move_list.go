package main

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
