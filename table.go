package main

import (
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

func ReadTableFile(filename string) Table {
	file, e := os.Open(filename)
	noError(e)

	scanner := bufio.NewScanner(file)
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

	return table
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
	counts := make([][]int, len(t))
	for y, row := range t {
		direction[y] = make([]int, len(row))
		counts[y] = make([]int, len(row))
	}

	for y := range t {
		direction[y][0] = NOWHERE
		if t[y][0].Matchable() {
			counts[y][0] = 1
		} else {
			counts[y][0] = 0
		}
	}
	for x := range t[0] {
		direction[0][x] = NOWHERE
		if t[0][x].Matchable() {
			counts[0][x] = 1
		} else {
			counts[0][x] = 0
		}
	}

	// Count matchabilities from top-left.
	for y := 1; y < len(t); y++ {
		for x := 1; x < len(t[y]); x++ {
			cell := t[y][x]

			if !cell.Matchable() {
				counts[y][x] = 0
				direction[y][x] = NOWHERE
				continue
			}

			topCell, leftCell := t[y-1][x], t[y][x-1]
			if topCell == cell && leftCell == cell {
				if counts[y-1][x] > counts[y][x-1] {
					counts[y][x] = counts[y-1][x] + 1
					direction[y][x] = UP
				} else {
					counts[y][x] = counts[y][x-1] + 1
					direction[y][x] = LEFT
				}
			} else if topCell == cell {
				counts[y][x] = counts[y-1][x] + 1
				direction[y][x] = UP
			} else if leftCell == cell {
				counts[y][x] = counts[y][x-1] + 1
				direction[y][x] = LEFT
			} else {
				counts[y][x] = 1
				direction[y][x] = NOWHERE
			}
		}
	}

	// Resolve matches backwards from bottom-right.
	for y := len(t) - 1; y >= 0; y-- {
		for x := len(t[y]) - 1; x >= 0; x-- {
			count := counts[y][x]
			if count < 3 {
				continue
			}

			matches += 1
			if direction[y][x] == UP {
				for i := 0; i < count; i++ {
					t[y-i][x] = LAND
				}
			} else {
				for i := 0; i < count; i++ {
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
