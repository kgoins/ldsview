package ldsview

type IEntityFilter interface {
	Matches(entity Entity) bool
}

// MatchesFilter Returns false if an entity fails to match a provided filter.
// If no filter is specified, will always return true.
func MatchesFilter(entity Entity, filterConditions []IEntityFilter) bool {
	matches := true

	for _, filter := range filterConditions {
		if !filter.Matches(entity) {
			matches = false
			break
		}
	}

	return matches
}
