package ldsview

import "github.com/kgoins/ldsview/internal"

// EntitySource returns LDAP Entites constructed from an input data source.
// If filter values are not set, they will default to empty lists. Null values will cause a panic.
type EntitySource interface {
	BuildEntity(keyAttrName string, keyAttrVal string) (Entity, error)
	BuildEntities() ([]Entity, error)

	SetEntityFilter(filter []IEntityFilter)
	SetAttributeFilter(includeAttrs internal.HashSetStr)
}
