package main

import "strconv"

func Solve(table Table, depth int) MoveList {
	return solveCore(table, depth, EmptyMove, GenerateMoveList(table), make(MoveList, 0))
}

func solveCore(src Table, depth int, lastMove Move, moves, solutionSoFar MoveList) MoveList {
	if depth == 0 {
		return EmptyMoveList
	}

	table := src
	for _, move := range moves {
		if move == lastMove {
			continue
		}

		move.Apply(table)
		solution := solveCore(table, depth-1, move, moves, append(solutionSoFar, move))
		log("["+strconv.Itoa(depth)+"] "+move.String(), table)
		move.Apply(table)

		if len(solution) > 0 {
			return solution
		}
	}

	return make(MoveList, 0)
}
