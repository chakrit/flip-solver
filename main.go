package main

import "fmt"
import "time"

func main() {
	table := ReadTableFile("input.txt")
	moves := GenerateMovesFromTable(table)

	report("initial", table)
	time.Sleep(1)

	for _, move := range moves {
		move.ApplyTo(table)
		report(move.String(), table)
		move.ApplyTo(table) // reverts it
	}
}

func report(label string, table *Table) {
	t := *table
	fmt.Printf("%v ---\n%v\n", label, t.String())
}

// shared utils
func ensure(e error) {
	if e != nil {
		panic(e)
	}
}
