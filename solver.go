package main

func Solve(table Table, depth int) MoveList {
	return solveCore(table, depth, NewCache(), EmptyMove, GenerateMoveList(table), make(MoveList, 0))
}

func solveCore(src Table, depth int, cache Cache, lastMove Move, moves, solutionSoFar MoveList) MoveList {
	if depth == 0 {
		if src.IsSolution() {
			return solutionSoFar
		}
		return EmptyMoveList
	}

	for _, move := range moves {
		if move == lastMove || !move.Valid(src) {
			continue
		}

		table, ok := cache.Lookup(src, move)
		if !ok {
			table = src.Clone()
			move.Apply(table)
			cache.Record(src, move, table)
		}

		solution := solveCore(table, depth-1, cache, move, moves, append(solutionSoFar, move))
		if len(solution) > 0 {
			return solution
		}
	}

	return make(MoveList, 0)
}
