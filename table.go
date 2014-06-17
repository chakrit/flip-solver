package main

import (
	"bufio"
	"fmt"
	"os"
)

type Table []Row

func NewTable(rows int, columns int) *Table {
	table := Table(make([]Row, rows))
	for i := range table {
		table[i] = NewRow(columns)
	}
	return &table
}

func ReadTableFile(filename string) *Table {
	file, err := os.Open(filename)
	ensure(err)

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	ensure(file.Close())

	if len(lines) <= 0 {
		panic(fmt.Errorf("no data in file: %v", filename))
	}

	table := NewTable(len(lines), len(lines[0]))
	t := *table
	for row, line := range lines {
		for col, char := range line {
			t[row][col] = Cell(char)
		}
	}

	return table
}

func (table *Table) SwapY(x, y, destY int) {
	t := *table
	t[y][x], t[destY][x] = t[destY][x], t[y][x]
}

func (table *Table) SwapX(y, x, destX int) {
	t := *table
	t[y][x], t[y][destX] = t[y][destX], t[y][x]
}

func (table *Table) String() string {
	s := ""
	for _, row := range *table {
		s += row.String() + "\n"
	}
	return s
}
