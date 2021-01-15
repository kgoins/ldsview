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

func GetStructure(source *LdifParser) ([]string, error) {
	filterParts := []string{"dn", "distinguishedName"}
	filter := BuildAttributeFilter(filterParts)
	source.SetAttributeFilter(filter)

	entities, done, cont := make(chan Entity), make(chan bool), make(chan bool)

	dnList := []string{}
	go func(ents chan Entity, done chan bool, list *[]string) {
		for entity := range ents {
			dn, dnFound := entity.GetDN()
			if dnFound {
				dnList = append(dnList, dn.Value.GetSingleValue())
			}
		}
		done <- true //signals to BuildEntities we're done processing
		cont <- true //signals to self we're done processing

	}(entities, done, &dnList)

	err := source.BuildEntities(entities, done)
	if err != nil {
		return nil, err
	}

	<-cont // wait for dnList to be populated
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
