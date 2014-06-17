package main

func Solve(table Table, depth int) MoveList {
	return solveCore(table, depth, EmptyMove, GenerateMoveList(table), make(MoveList, 0))
}

func solveCore(src Table, depth int, lastMove Move, moves, solutionSoFar MoveList) MoveList {
	if depth == 0 {
		// log("attempt: " + solutionSoFar.String(), src.String())

		if src.IsSolution() {
			return solutionSoFar
		}
		return EmptyMoveList
	}

	for _, move := range moves {
		if move == lastMove || !move.Valid(src) {
			continue
		}

		table := src.Clone()
		move.Apply(table)

		solution := solveCore(table, depth-1, move, moves, append(solutionSoFar, move))
		if len(solution) > 0 {
			return solution
		}
	}

	return make(MoveList, 0)
}
