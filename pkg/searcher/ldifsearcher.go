package searcher

import (
	filter "github.com/kgoins/entityfilter/entityfilter"
	"github.com/kgoins/ldapentity/entity"

	"github.com/kgoins/ldifparser"
	"github.com/kgoins/ldifparser/entitybuilder"
)

type LdifSearcher struct {
	reader  ldifparser.LdifReader
	matcher LdapEntityMatcher
}

// Verify that LdifSearcher implements LdapSearcher
var _ LdapSearcher = LdifSearcher{}

func NewLdifSearcher(input ldifparser.ReadSeekerAt) LdifSearcher {
	return LdifSearcher{
		reader:  ldifparser.NewLdifReader(input),
		matcher: NewLdapEntityMatcher(),
	}
}

func (s LdifSearcher) ReadEntity(keyAttrName, keyAttrVal string) (e entity.Entity, err error) {
	return s.reader.ReadEntity(keyAttrName, keyAttrVal)
}

func (s LdifSearcher) ReadAllEntities(done <-chan bool, af entitybuilder.AttributeFilter) <-chan entity.Entity {
	s.reader.SetAttributeFilter(af)
	return s.reader.ReadEntitiesChanneled(done)
}

func (s LdifSearcher) SearchEntities(
	done <-chan bool,
	af entitybuilder.AttributeFilter,
	ef filter.EntityFilter,
) <-chan entity.Entity {
	resultStream := make(chan entity.Entity)

	s.reader.SetAttributeFilter(af)
	entityStream := s.reader.ReadEntitiesChanneled(done)

	go func() {
		for entity := range entityStream {
			matches, err := s.matcher.Matches(entity, ef)
			if err != nil {
				continue
			}

			if matches {
				resultStream <- entity
			}
		}
	}()

	return resultStream
}
