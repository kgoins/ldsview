package ldsview

import (
	"github.com/kgoins/ldapentity/entity/ad"
	"github.com/kgoins/ldifparser/entitybuilder"
	"github.com/kgoins/ldsview/pkg/searcher"
)

// CountEntities returns the number of entities in the input file
func CountEntities(searcher searcher.LdapSearcher) (count int, err error) {
	done := make(chan bool)
	defer close(done)

	af := entitybuilder.NewAttributeFilter(ad.ATTR_DN)
	eStream := searcher.ReadAllEntities(done, af)

	for range eStream {
		count++
	}

	return
}
