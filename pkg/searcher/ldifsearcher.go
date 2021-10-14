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

type EntityResult ldifparser.EntityResp

// The output channel is closed by the underlying reader when all entities
// have been read.
func (s LdifSearcher) ReadAllEntities(done <-chan bool, af entitybuilder.AttributeFilter) <-chan EntityResult {
	results := make(chan EntityResult)

	go func() {
		defer close(results)
		s.reader.SetAttributeFilter(af)

		for r := range s.reader.ReadEntitiesChanneled(done) {
			er := EntityResult(r)
			results <- er
		}
	}()

	return results
}

// SearchEntities will iterate through all entities in the read stream,
// returning those who match the input filters. If an error is encountered,
// the error will be packaged in the EntityResult, sent thorugh the channel,
// and further processing will stop.
func (s LdifSearcher) SearchEntities(
	done <-chan bool,
	af entitybuilder.AttributeFilter,
	ef filter.EntityFilter,
) <-chan EntityResult {
	resultStream := make(chan EntityResult)

	s.reader.SetAttributeFilter(af)
	entityStream := s.ReadAllEntities(done, af)

	go func() {
		defer close(resultStream)
		for entity := range entityStream {
			// This will halt further processing
			if entity.Error != nil {
				resultStream <- entity
				return
			}

			matches, err := s.matcher.Matches(entity.Entity, ef)

			// This will halt further processing
			if err != nil {
				entity.Error = err
				resultStream <- entity
				return
			}

			if matches {
				resultStream <- entity
				continue
			}
		}
	}()

	return resultStream
}
