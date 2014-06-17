package main

func Solve(table Table, depth int) MoveList {
	return solveCore(table, depth, EmptyMove, GenerateMoveList(table), make(MoveList, 0))
}

func solveCore(src Table, depth int, lastMove Move, moves, solutionSoFar MoveList) MoveList {
	if depth == 0 {
		// log(" ---- ")
		// log(solutionSoFar.String(), src.String())
		log("----", src.String())
		return EmptyMoveList
	}

	for _, move := range moves {
		if move == lastMove || !move.Valid(src) {
			continue
		}

		table := src.Clone()
		move.Apply(table)
		for n := 1; n > 0; n = table.ResolveMatches() { }
		for n := 1; n > 0; n = table.ResolveGravity() { }

		solution := solveCore(table, depth-1, move, moves, append(solutionSoFar, move))
		if len(solution) > 0 {
			return solution
		}
	}

	return make(MoveList, 0)
}
