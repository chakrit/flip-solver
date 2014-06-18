package main

func Solve(table Table, depth int) MoveList {
	return solveCore(NewContext(table, depth))
}

func solveCore(c *Context) MoveList {
	if c.depth == 0 {
		if c.table.IsSolution() {
			return c.solution
		}

		return EmptyMoveList
	}

	for _, move := range c.possibleMoves {
		if move == c.lastMove || !move.Valid(c.table) {
			continue
		}

		table, ok := c.cache.Lookup(c.table, move)
		if !ok {
			table = c.table.Clone()
			move.Apply(table)
			c.cache.Record(c.table, move, table)
		}

		solution := solveCore(c.AppendStep(table, move))
		if len(solution) > 0 {
			return solution
		}
	}

	return EmptyMoveList
}
