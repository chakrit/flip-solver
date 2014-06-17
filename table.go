package main

import (
	"strconv"
	"bufio"
	"fmt"
	"os"
)

type Table []Row

func NewTable(rows int, columns int) Table {
	table := Table(make([]Row, rows))
	for i := range table {
		table[i] = NewRow(columns)
	}
	return table
}

func ReadTableFile(filename string) (Table, int) {
	file, e := os.Open(filename)
	noError(e)

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		panic(fmt.Errorf("first line must be optimal move count."))
	}

	moves, e := strconv.Atoi(scanner.Text())
	noError(e)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	noError(file.Close())

	if len(lines) <= 0 {
		panic(fmt.Errorf("no data in file: %v", filename))
	}

	table := NewTable(len(lines), len(lines[0]))
	for row, line := range lines {
		for col, char := range line {
			table[row][col] = Cell(char)
		}
	}

	return table, moves
}

func (t Table) Clone() Table {
	// TODO: Convert this block to use the built-in copy()
	clone := make(Table, len(t))
	for y, row := range t {
		clone[y] = make(Row, len(row))
		for x := range row {
			clone[y][x] = t[y][x]
		}
	}

	return clone
}

func (t Table) IsSolution() bool {
	for _, row := range t {
		for _, cell := range row {
			if cell.Matchable() {
				return false // There's still something matchable left.
			}
		}
	}

	return true
}

func (t Table) SwapY(x, y, destY int) {
	t[y][x], t[destY][x] = t[destY][x], t[y][x]
}

func (t Table) SwapX(y, x, destX int) {
	t[y][x], t[y][destX] = t[y][destX], t[y][x]
}

func (t Table) Resolve() int {
	return t.resolveMatches() + t.resolveGravity()
}

func (t Table) resolveMatches() int {
	// Count matches using DP counting table.
	matches := 0
	const (
		NOWHERE = iota
		UP
		LEFT
	)

	direction := make([][]int, len(t))
	countUps := make([][]int, len(t))
	countLefts := make([][]int, len(t))

	for y, row := range t {
		direction[y] = make([]int, len(row))
		countUps[y] = make([]int, len(row))
		countLefts[y] = make([]int, len(row))
	}

	same := func(y1, x1, y2, x2 int) bool { return t[y1][x1] == t[y2][x2] }
	matchable := func(y, x int) bool { return t[y][x].Matchable() }
	record := func(y, x, countUp, countLeft int) {
		countUps[y][x] = countUp
		countLefts[y][x] = countLeft
	}

	// Count top edge and left edge matchables.
	if matchable(0, 0) {
		record(0, 0, 1, 1)
	} else {
		record(0, 0, 0, 0)
	}

	for y := 1; y < len(t); y++ {
		if matchable(y, 0) {
			if same(y, 0, y-1, 0) {
				record(y, 0, countUps[y-1][0]+1, 1)
			} else {
				record(y, 0, 1, 1)
			}
		} else {
			record(y, 0, 0, 0)
		}
	}

	for x := 1; x < len(t[0]); x++ {
		if matchable(0, x) {
			if same(0, x, 0, x-1) {
				record(0, x, 1, countLefts[0][x-1]+1)
			} else {
				record(0, x, 1, 1)
			}
		} else {
			record(0, x, 0, 0)
		}
	}

	// Count matchabilities from top-left downwards and rightwards.
	for y := 1; y < len(t); y++ {
		for x := 1; x < len(t[y]); x++ {
			cell := t[y][x]

			if !cell.Matchable() {
				record(y, x, 0, 0)
			} else {
				sameUp, sameLeft := same(y, x, y-1, x), same(y, x, y, x-1)
				if sameUp && sameLeft {
					record(y, x, countUps[y-1][x]+1, countLefts[y][x-1]+1)
				} else if sameUp {
					record(y, x, countUps[y-1][x]+1, 1)
				} else if sameLeft  {
					record(y, x, 1, countLefts[y][x-1]+1)
				} else {
					record(y, x, 1, 1)
				}
			}
		}
	}

	// Resolve matches backwards from bottom-right.
	for y := len(t) - 1; y >= 0; y-- {
		for x := len(t[y]) - 1; x >= 0; x-- {
			ups := countUps[y][x]
			lefts := countLefts[y][x]

			if ups >= 3 {
				matches += 1
				for i := 0; i < ups; i++ {
					t[y-i][x] = LAND
				}
			}

			if lefts >= 3 {
				matches += 1
				for i := 0; i < lefts; i++ {
					t[y][x-i] = LAND
				}
			}
		}
	}

	return matches
}

func (t Table) resolveGravity() int {
	movements := 0
	for x := range t[0] {
		// Find the first gravity from the bottom. There is always only a single pocket of
		// land in any single run so we should be fine.
		landY, objectY, fallY := -1, -1, 0

		for y := len(t) - 1; y >= 0; y-- {
			switch t[y][x] {
			case LAND:
				landY = y
			case HOLLOW:
			default:
				continue
			}

			break
		}

		if landY == -1 {
			continue
		}

		for y := landY - 1; y >= 0; y-- {
			switch t[y][x] {
			case LAND:
				continue
			case HOLLOW:
			default:
				objectY = y
			}

			break
		}

		if objectY == -1 {
			continue
		}

		// Apply gravity effect.
		movements += 1
		fallY = landY - objectY
		for y := landY; y >= fallY; y-- {
			fallingCell := t[y-fallY][x]
			if fallingCell == HOLLOW {
				break
			}

			t.SwapY(x, y, y-fallY)
		}
	}

	return movements
}

func (t Table) String() string {
	s := ""
	for _, row := range t {
		s += row.String() + "\n"
	}
	return s
}
