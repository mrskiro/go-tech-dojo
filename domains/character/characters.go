package character

type Characters []Character

func (cs Characters) ToMap() map[uint64]Character {
	results := make(map[uint64]Character, len(cs))
	for _, v := range cs {
		results[v.ID] = v
	}
	return results
}
