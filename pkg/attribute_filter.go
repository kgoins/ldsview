package ldsview

import (
	"strings"

	"github.com/kgoins/ldsview/internal"
)

func BuildAttributeFilter(filterParts []string) internal.HashSetStr {
	set := internal.NewHashSetStr()

	for _, attr := range filterParts {
		set.Add(strings.ToLower(attr))
	}

	return set
}
