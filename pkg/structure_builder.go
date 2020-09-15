package ldsview

import (
	"strings"

	"github.com/kgoins/ldsview/internal"
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

func GetStructure(source EntitySource) ([]string, error) {
	filterParts := []string{"dn", "distinguishedName"}
	filter := BuildAttributeFilter(filterParts)

	source.SetAttributeFilter(filter)

	entities, err := source.BuildEntities()
	if err != nil {
		return nil, err
	}

	dnList := []string{}
	for _, entity := range entities {
		dn, dnFound := entity.GetDN()
		if dnFound {
			dnList = append(dnList, dn.Value.GetSingleValue())
		}
	}

	return buildStructureFromDNs(dnList), nil
}

func buildStructureFromDNs(dnList []string) []string {
	structure := internal.NewHashSetStr()

	for _, dn := range dnList {
		ouPath := extractPathFromDN(dn)
		structure.Add(ouPath)
	}

	return structure.Values()
}
