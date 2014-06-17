package main

import "fmt"
import "flag"
import "time"

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		panic("no puzzle filename given.")
	}

	filename := flag.Arg(0)
	table, movesCount := ReadTableFile(filename)

	ticker := time.NewTicker(time.Minute)
	go func() {
		for _ = range ticker.C {
			fmt.Println(time.Now().Format(time.RFC3339))
		}
	}()

	solution := Solve(table, movesCount)
	ticker.Stop()

	if len(solution) == 0 {
		panic("no solution for: " + filename)
	}

	log("\ninitial table:", table.String())
	for _, move := range solution {
		move.Apply(table)
		log(move.String(), table.String())
	}
}
