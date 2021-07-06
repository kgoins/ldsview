package searcher

import (
	filter "github.com/kgoins/entityfilter/entityfilter"
	"github.com/kgoins/ldapentity/entity"
	"github.com/kgoins/ldifparser/entitybuilder"
)

type LdapSearcher interface {
	ReadEntity(keyAttrName, keyAttrVal string) (e entity.Entity, err error)
	ReadAllEntities(done <-chan bool, af entitybuilder.AttributeFilter) <-chan entity.Entity

	SearchEntities(
		done <-chan bool,
		af entitybuilder.AttributeFilter,
		ef filter.EntityFilter,
	) <-chan entity.Entity
}
