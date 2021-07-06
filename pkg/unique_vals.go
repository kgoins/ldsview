package ldsview

import (
	hashset "github.com/kgoins/hashset/pkg"
	"github.com/kgoins/ldifparser/entitybuilder"
	"github.com/kgoins/ldsview/pkg/searcher"
)

func GetUniqueValues(searcher searcher.LdapSearcher, attrName string) ([]string, error) {
	filter := entitybuilder.NewAttributeFilter(attrName)

	done := make(chan bool)
	defer close(done)

	results := searcher.ReadAllEntities(done, filter)
	valSet := hashset.NewStrHashset()

	for entity := range results {
		attr, found := entity.GetAttribute(attrName)
		if !found {
			continue
		}

		vals := attr.GetValues()
		valSet.Add(vals...)
	}

	return valSet.Values(), nil
}
