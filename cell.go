package main

type Cell byte

const (
	HOLLOW Cell = ' '
	LAND        = '.'
	CRATE       = '#'

	PINK  Cell = 'p'
	TEAL       = 't'
	BROWN      = 'b'
	BEIGE      = 'g'
	WHITE      = 'w'
)

func (c Cell) Matchable() bool {
	switch c {
	case HOLLOW, LAND, CRATE:
		return false
	default:
		return true
	}
}
