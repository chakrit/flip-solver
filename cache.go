package main

// Key is table.String() + " " + move.String()
type Cache map[string]Table

func NewCache() Cache {
	return make(Cache, 1024*1024)
}

func (c Cache) Lookup(table Table, move Move) (Table, bool) {
	result, ok := c[cacheKey(table, move)]
	return result, ok
}

func (c Cache) Record(table Table, move Move, outcome Table) {
	c[cacheKey(table, move)] = outcome
}

func cacheKey(table Table, move Move) string {
	return table.String() + " " + move.String()
}

func (c Cache) Clear() {
	for key := range c {
		delete(c, key)
	}
}
