package searcher

import (
	filter "github.com/kgoins/entityfilter/entityfilter"
	"github.com/kgoins/entityfilter/entitymatcher"
	"github.com/kgoins/ldapentity/entity"
)

type LdapEntityMatcher struct {
	conditions entitymatcher.ConditionMap
}

func NewLdapEntityMatcher() LdapEntityMatcher {
	return LdapEntityMatcher{
		conditions: entitymatcher.NewConditionMap(),
	}
}

func (matcher LdapEntityMatcher) matchesFilterEntry(entity entity.Entity, filterEntry filter.FilterEntry) (bool, error) {
	attr, found := entity.GetAttribute(filterEntry.AttributeName)
	if !found {
		return false, nil
	}

	if filterEntry.IsWildcard {
		return true, nil
	}

	for _, val := range attr.GetValues() {
		matched, err := matcher.conditions.MatchesCondition(
			val,
			filterEntry.Value,
			filterEntry.Condition,
		)
		if err != nil {
			return false, err
		}

		if matched {
			return true, nil
		}
	}

	return false, nil
}

func (matcher LdapEntityMatcher) Matches(entity entity.Entity, filter filter.EntityFilter) (bool, error) {
	for _, filterEntry := range filter.GetEntries() {
		matches, err := matcher.matchesFilterEntry(entity, filterEntry)
		if err != nil {
			return false, err
		}

		if matches {
			return true, nil
		}
	}

	return false, nil
}
