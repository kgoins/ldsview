package ldsview

import (
	"strings"
)

type Entity struct {
	attributes map[string]EntityAttribute
}

func NewEntity() Entity {
	return Entity{
		attributes: make(map[string]EntityAttribute),
	}
}

func (e *Entity) AddAttribute(attr EntityAttribute) {
	attrName := strings.ToLower(attr.Name)

	existing, found := e.GetAttribute(attrName)
	if !found {
		e.attributes[attrName] = attr
		return
	}

	existing.Value.Add(attr.Value.Values()...)
}

func (e Entity) IsEmpty() bool {
	return len(e.attributes) == 0
}

func (e Entity) Groups() []string {
	groupAttr, found := e.GetAttribute("memberOf")
	if !found {
		return []string{}
	}

	return groupAttr.Value.Values()
}

func (e Entity) GetDN() (EntityAttribute, bool) {
	dn, found := e.GetAttribute("dn")
	if !found {
		dn, found = e.GetAttribute("distinguishedName")
	}

	return dn, found
}

func (e Entity) GetAllAttributeNames() []string {
	names := make([]string, e.Size())

	i := 0
	for key := range e.attributes {
		names[i] = key
		i++
	}

	return names
}

func (e Entity) GetAllAttributes() []EntityAttribute {
	attrs := make([]EntityAttribute, e.Size())

	i := 0
	for _, val := range e.attributes {
		attrs[i] = val
	}

	return attrs
}

func (e *Entity) SetAttribute(attr EntityAttribute) {
	attrName := strings.ToLower(attr.Name)
	e.attributes[attrName] = attr
}

func (e Entity) GetAttribute(name string) (EntityAttribute, bool) {
	val, found := e.attributes[strings.ToLower(name)]
	return val, found
}

func (e Entity) Size() int {
	return len(e.attributes)
}

func (e *Entity) decodeFromGeneralizedTime(attrName string) {
	timeAttr, found := e.GetAttribute(attrName)
	if !found {
		return
	}

	origTime := timeAttr.Value.Values()[0]
	decodedTime, _ := TimeFromADGeneralizedTime(origTime)

	timeAttr.SetValue(decodedTime.String())
	e.SetAttribute(timeAttr)
}

func (e *Entity) decodeFromADTimestamp(attrName string) {
	timeAttr, found := e.GetAttribute(attrName)
	if !found {
		return
	}

	origTime := timeAttr.Value.Values()[0]
	decodedTime := TimeFromADTimestamp(origTime)

	timeAttr.SetValue(decodedTime.String())
	e.SetAttribute(timeAttr)
}

func (e *Entity) DeocdeTimestamps() {
	e.decodeFromGeneralizedTime("whenCreated")
	e.decodeFromGeneralizedTime("whenChanged")

	e.decodeFromADTimestamp("pwdLastSet")
	e.decodeFromADTimestamp("lastLogon")
	e.decodeFromADTimestamp("lastLogonTimestamp")
}
