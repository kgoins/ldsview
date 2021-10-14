package searcher

import (
	filter "github.com/kgoins/entityfilter/entityfilter"
	"github.com/kgoins/ldapentity/entity"
	"github.com/kgoins/ldifparser/entitybuilder"
)

type LdapSearcher interface {
	ReadEntity(keyAttrName, keyAttrVal string) (e entity.Entity, err error)

	// The output channel is closed by the underlying reader when all entities
	// have been read.
	ReadAllEntities(done <-chan bool, af entitybuilder.AttributeFilter) <-chan EntityResult

	// The output channel is closed by the underlying reader when all entities
	// have been read. Sending a value over the done channel will interrupt the search.
	SearchEntities(
		done <-chan bool,
		af entitybuilder.AttributeFilter,
		ef filter.EntityFilter,
	) <-chan EntityResult
}
