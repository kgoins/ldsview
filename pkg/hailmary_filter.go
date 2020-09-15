package ldsview

type HailmaryFilter struct {
	searchTerm string
}

var hailmaryAttributeNames = []string{
	"cn",
	"description",
	"samaccountname",
	"title",
	"name",
	"givenName",
	"l",
	"info",
}

func NewHailmaryFilter(searchTerm string) HailmaryFilter {
	return HailmaryFilter{searchTerm: searchTerm}
}

func (filter HailmaryFilter) Matches(entity Entity) bool {
	for _, attrName := range hailmaryAttributeNames {
		attr, found := entity.GetAttribute(attrName)
		if found {
			if attr.Value.ContainsSubstrIgnoreCase(filter.searchTerm) {
				return true
			}
		}
	}

	return false
}
