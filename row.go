package main

type Row []Cell

func NewRow(cells int) Row {
	row := Row(make([]Cell, cells))
	for i := range row {
		row[i] = HOLLOW
	}
	return row[:]
}

func (row Row) String() string {
	s := ""
	for _, char := range row {
		s += string(char) + " "
	}
	return s
}
