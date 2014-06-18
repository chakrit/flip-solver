package main

// Key is table.String() + " " + move.String()
type MoveCache map[string]Table

func NewMoveCache() MoveCache {
	return make(MoveCache, 1024*1024)
}

func (c MoveCache) Lookup(table Table, move Move) (Table, bool) {
	result, ok := c[cacheKey(table, move)]
	return result, ok
}

func (c MoveCache) Record(table Table, move Move, outcome Table) {
	c[cacheKey(table, move)] = outcome
}

func cacheKey(table Table, move Move) string {
	return table.String() + " " + move.String()
}

func (c MoveCache) Clear() {
	for key := range c {
		delete(c, key)
	}
}
