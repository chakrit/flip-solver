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

func (t Table) SwapY(x, y, destY int) {
	t[y][x], t[destY][x] = t[destY][x], t[y][x]
}

func (t Table) SwapX(y, x, destX int) {
	t[y][x], t[y][destX] = t[y][destX], t[y][x]
}

func (t Table) String() string {
	s := ""
	for _, row := range t {
		s += row.String() + "\n"
	}
	return s
}
