package ldsview

import (
	"strings"

	hashset "github.com/kgoins/hashset/pkg"
	"github.com/kgoins/ldifparser/entitybuilder"
	"github.com/kgoins/ldsview/pkg/searcher"
)

func extractPathFromDN(dn string) string {
	dnParts := strings.SplitN(dn, ",", 2)
	if len(dnParts) < 2 {
		return ""
	}

	if strings.HasPrefix(dnParts[0], "CN=") {
		return dnParts[1]
	}

	return ""
}

func GetStructure(searcher searcher.LdapSearcher) ([]string, error) {
	filterParts := []string{"dn", "distinguishedName"}
	filter := entitybuilder.NewAttributeFilter(filterParts...)

	done := make(chan bool)
	defer close(done)

	entities := searcher.ReadAllEntities(done, filter)
	return buildStructureFromDNs(entities)
}

func buildStructureFromDNs(entities <-chan searcher.EntityResult) ([]string, error) {
	structure := hashset.NewStrHashset()

	for e := range entities {
		if e.Error != nil {
			return nil, e.Error
		}

		dn, found := e.Entity.GetDN()
		if !found {
			continue
		}

		ouPath := extractPathFromDN(dn)
		structure.Add(ouPath)
	}

	return structure.Values(), nil
}
