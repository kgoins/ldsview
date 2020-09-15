package ldsview

import (
	"errors"
	"sort"
	"strings"
)

type FilterCondition string

const (
	FILTER_EQUALS       FilterCondition = ":="
	FILTER_NOT_EQUALS   FilterCondition = ":!="
	FILTER_CONTAINS     FilterCondition = ":~"
	FILTER_NOT_CONTAINS FilterCondition = ":!~"
)

var FilterConditions = []FilterCondition{
	FILTER_EQUALS,
	FILTER_NOT_EQUALS,
	FILTER_CONTAINS,
	FILTER_NOT_CONTAINS,
}

type EntityFilter struct {
	AttributeName string
	Value         string
	Condition     FilterCondition
	IsWildcard    bool
}

type filterByLength []FilterCondition

func (a filterByLength) Len() int      { return len(a) }
func (a filterByLength) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a filterByLength) Less(i, j int) bool {
	li, lj := len(a[i]), len(a[j])
	if li == lj {
		return a[i] > a[j]
	}
	return li > lj
}

func getFilterCondition(filterStr string, conditions []FilterCondition) (FilterCondition, error) {
	for _, condition := range conditions {
		if strings.Contains(filterStr, string(condition)) {
			return condition, nil
		}
	}

	return FILTER_EQUALS, errors.New("Unable to identify filter condition")
}

func buildEntityFilter(filterStr string) (EntityFilter, error) {
	condition, err := getFilterCondition(filterStr, FilterConditions)
	if err != nil {
		return EntityFilter{}, err
	}

	filterParts := strings.Split(filterStr, string(condition))
	if len(filterParts) != 2 {
		return EntityFilter{},
			errors.New("Invalid filter format: " + filterStr)
	}

	filterValue := filterParts[1]
	isWildcard := filterValue == "*"

	filter := EntityFilter{
		AttributeName: filterParts[0],
		Value:         filterValue,
		Condition:     condition,
		IsWildcard:    isWildcard,
	}

	return filter, nil
}

func BuildEntityFilter(filterStrings []string) ([]IEntityFilter, error) {
	filter := []IEntityFilter{}
	sort.Sort(filterByLength(FilterConditions))

	for _, filterStr := range filterStrings {
		if strings.TrimSpace(filterStr) == "" {
			continue
		}

		newFilter, err := buildEntityFilter(filterStr)
		if err != nil {
			return nil, err
		}

		filter = append(filter, newFilter)
	}

	return filter, nil
}

func (filter EntityFilter) Matches(entity Entity) bool {
	attr, found := entity.GetAttribute(filter.AttributeName)
	if !found {
		return false
	}

	if filter.IsWildcard {
		return true
	}

	matches := false

	switch filter.Condition {
	case FILTER_CONTAINS:
		matches = attr.Value.ContainsSubstr(filter.Value)
	case FILTER_NOT_CONTAINS:
		matches = !attr.Value.ContainsSubstr(filter.Value)
	case FILTER_EQUALS:
		matches = attr.Value.Contains(filter.Value)
	case FILTER_NOT_EQUALS:
		matches = !attr.Value.Contains(filter.Value)
	}

	return matches
}
