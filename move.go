package main

type Move struct {
	// Assume Dest is always to the bottom right.
	Src  Point
	Dest Point
}

var EmptyMove = Move{Point{0, 0}, Point{0, 0}}

func NewMove(src Point, dest Point) Move {
	s, d := src, dest
	if s.X > d.X {
		s.X, d.X = d.X, s.X
	}
	if s.Y > d.Y {
		s.Y, d.Y = d.Y, s.Y
	}

	return Move{src, dest}
}

func (m Move) String() string {
	return m.Src.String() + " -> " + m.Dest.String()
}

func (m Move) Valid(table Table) bool {
	for x := m.Src.X; x <= m.Dest.X; x++ {
		for y := m.Src.Y; y <= m.Dest.Y; y++ {
			if table[y][x] == HOLLOW {
				return false
			}
		}
	}

	return true
}

func (m Move) Apply(t Table) {
	// assume m.Src < m.Dest always
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
