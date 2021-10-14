package ldsview

import "github.com/kgoins/ldapentity/entity"

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

func (filter HailmaryFilter) Matches(entity entity.Entity) bool {
	for _, attrName := range hailmaryAttributeNames {
		attr, found := entity.GetAttribute(attrName)
		if found {
			if attr.Value.ContainsSubstr(filter.searchTerm, true) {
				return true
			}
		}
	}

	return false
}
