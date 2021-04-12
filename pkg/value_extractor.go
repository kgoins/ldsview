package ldsview

import (
	hashset "github.com/kgoins/hashset/pkg"
)

func GetValues(source *LdifParser, attrName string) ([]string, error) {
	filterParts := []string{attrName}
	filter := BuildAttributeFilter(filterParts)
	source.SetAttributeFilter(filter)

	entities, done, cont := make(chan Entity), make(chan bool), make(chan bool)

	valSet := hashset.NewStrHashset()
	go func(ents chan Entity, done chan bool, list *hashset.StrHashset) {
		for entity := range ents {
			if entity.IsEmpty() {
				continue
			}

			attr, found := entity.GetAttribute(attrName)
			if !found {
				continue
			}

			valSet.Add(attr.Value.Values()...)
		}

		done <- true //signals to BuildEntities we're done processing
		cont <- true //signals to self we're done processing

	}(entities, done, &valSet)

	err := source.BuildEntities(entities, done)
	if err != nil {
		return nil, err
	}

	<-cont // wait for dnList to be populated
	return valSet.Values(), nil
}
