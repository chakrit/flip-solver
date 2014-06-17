package main

func main() {
	table := ReadTableFile("puzzles/1-1")

	solution := Solve(table, 5)
	for _, move := range solution {
		move.Apply(table)
		log(move.String(), table.String())
	}
}
