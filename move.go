package main

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
