package main

// TODO: Right now just using thise for impossibility cache, but we could extend this to
//   store more stateful information.
type StateCache map[string]bool

func NewStateCache() StateCache {
	return make(StateCache, 1024*1024)
}

func (c StateCache) Lookup(table Table) (bool, bool) {
	result, ok := c[c.cacheKey(table)]
	return result, ok
}

func (c StateCache) Record(table Table, possible bool) {
	c[c.cacheKey(table)] = possible
}

func (c StateCache) Clear() {
	for key := range c {
		delete(c, key)
	}
}

func (c StateCache) cacheKey(table Table) string {
	return table.String()
}
