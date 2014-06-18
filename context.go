package main

type Context struct {
	table         Table
	depth         int
	lastMove      Move
	possibleMoves []Move
	solution      MoveList

	moveCache     MoveCache
	stateCache    StateCache
}

func NewContext(table Table, depth int) *Context {
	return &Context{
		table: table,
		depth: depth,
		lastMove: EmptyMove,
		possibleMoves: GenerateMoveList(table),
		solution: EmptyMoveList,

		moveCache: NewMoveCache(),
		stateCache: NewStateCache(),
	}
}

func (c *Context) AppendStep(table Table, move Move) *Context {
	child := &Context{}
	*child = *c
	child.table = table
	child.depth -= 1
	child.lastMove = move
	child.solution = append(c.solution, move)
	return child
}
